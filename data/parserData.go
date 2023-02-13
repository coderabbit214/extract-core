package data

import (
	"encoding/json"
	"github.com/coderabbit214/extract-core/dictionary"
	"github.com/coderabbit214/extract-core/result"
	"github.com/coderabbit214/extract-core/utils"
	"regexp"
)

type ParserData struct {
	DocInfo DocInfo  `json:"docInfo"`
	Layouts []Layout `json:"layouts"`
	Logics  Logics   `json:"logics"`
	Styles  []Style  `json:"styles"`
	Version string   `json:"version"`
}

// NewData init ParserData
func NewData(str string) (ParserData, error) {
	data := &ParserData{}
	err := json.Unmarshal([]byte(str), data)
	if err != nil {
		return *data, err
	}
	for i := 0; i < len(data.Layouts); i++ {
		data.Layouts[i].ParserData = data
	}
	return *data, nil
}

// GetUniqueIdsByField 根据filter筛查结果范围
func (a *ParserData) GetUniqueIdsByField(filters []dictionary.Filter, rangeMap map[string][]string) ([]string, error) {
	var ranges []string
	//数据过滤
	for i, filter := range filters {
		if filter.Name != "" && rangeMap != nil {
			ranges = rangeMap[filter.Name]
			continue
		}
		if i == 0 {
			ranges = a.filter(filter, nil)
			continue
		}
		ranges = a.filter(filter, ranges)
	}
	return ranges, nil
}

// 过滤
func (a *ParserData) filter(filter dictionary.Filter, ranges []string) []string {
	extract := a.getUniqueIDsTextExtract(ranges, filter.Regular, filter.Type)
	var uIds []string
	switch filter.Kind {
	case dictionary.KindParent:
		for _, v := range extract {
			uIds = a.findUniqueIdByParentUniqueIdFromLogics(v)
		}
	case dictionary.KindUnder:
		for _, v := range extract {
			uIds = a.findUniqueIdByUnderUniqueIdFromLogics(v)
		}
	case dictionary.KindAbove:
		for _, v := range extract {
			uIds = a.findUniqueIdByAboveUniqueIdFromLogics(v)
		}
	}
	if ranges == nil {
		return uIds
	}
	ranges = utils.IntersectionStrings(ranges, uIds)
	return ranges
}

// getUniqueIDsTextExtract 普通筛选 非表格内部
func (a *ParserData) getUniqueIDsTextExtract(ids []string, regular string, dataType string) []string {
	res := make([]string, 0)

	reg := regexp.MustCompile(regular)
	for _, layout := range a.Layouts {
		if layout.Type == dataType && reg.MatchString(layout.Text) {
			if ids == nil || utils.InString(layout.UniqueID, ids) {
				res = append(res, layout.UniqueID)
			}
		}
	}
	return res
}

func (a *ParserData) ContextExtract(ids []string, regular string, dataType string) result.Result {
	reg := regexp.MustCompile(regular)
	groupNames := reg.SubexpNames()
	if len(groupNames) == 2 && groupNames[1] == utils.Val {
		return a.ContentExtract(ids, regular, dataType, 1)
	} else if len(groupNames) >= 2 {
		return a.GroupExtract(ids, regular, dataType)
	} else {
		return a.ContentExtract(ids, regular, dataType, 0)
	}
}

// ContentExtract 普通(上下文段落)提取
func (a *ParserData) ContentExtract(ids []string, regular string, dataType string, groupNumber int) result.Result {
	res := result.Result{
		ID:        utils.GenerateUUID(),
		FieldType: result.FieldTypeContext,
		Contexts:  nil,
		Groups:    nil,
	}
	contexts := make([]result.Context, 0)

	for _, layout := range a.Layouts {
		if layout.Type == dataType && (ids == nil || utils.InString(layout.UniqueID, ids)) {
			context, ok := layout.ToResultContextByRegularAndGroupNumber(regular, groupNumber)
			if ok {
				contexts = append(contexts, *context...)
			}
		}
	}
	res.Contexts = contexts
	return res
}

// GroupExtract 多结果提取
func (a *ParserData) GroupExtract(ids []string, regular string, dataType string) result.Result {
	res := result.Result{
		ID:        utils.GenerateUUID(),
		FieldType: result.FieldTypeGroup,
		Contexts:  nil,
		Groups:    nil,
	}
	groupss := make([]map[string]result.Context, 0)

	for _, layout := range a.Layouts {
		if layout.Type == dataType && (ids == nil || utils.InString(layout.UniqueID, ids)) {
			groups, ok := layout.ToResultGroupByRegular(regular)
			if ok {
				groupss = append(groupss, *groups...)
			}
		}
	}

	res.Groups = groupss
	return res
}

// CellsExtract 单元格提取
func (a *ParserData) CellsExtract(ids []string, regulars []dictionary.CellsRegular) result.Result {
	res := result.Result{
		ID:        utils.GenerateUUID(),
		FieldType: result.FieldTypeContext,
		Contexts:  nil,
		Groups:    nil,
	}
	contexts := make([]result.Context, 0)
	for _, layout := range a.Layouts {
		if layout.Type == TypeTable && (ids == nil || utils.InString(layout.UniqueID, ids)) {
			cellIds := layout.TableCellsRegular(regulars)
			//过滤结果处理
			if len(cellIds) > 0 {
				localCells := layout.GetCellsByIds(cellIds)
				for _, cell := range localCells {
					context := &result.Context{
						Text:   cell.GetText(),
						Blocks: cell.GetBlocks(a),
					}
					contexts = append(contexts, *context)
				}
			}
		}
	}
	res.Contexts = contexts
	return res
}

// 在上下文中根据父元素查找所有子元素（不包括自己）
func (a *ParserData) findUniqueIdByParentUniqueIdFromLogics(uniqueId string) []string {
	l := a.Logics
	treeUniqueIds := make([]string, 0)
	treeUniqueIds = append(treeUniqueIds, uniqueId)
	for _, d := range l.DocTree {
		for _, uId := range d.Backlink.Superior {
			if utils.InString(uId, treeUniqueIds) {
				treeUniqueIds = append(treeUniqueIds, d.UniqueID)
			}
		}
	}
	return utils.DeleteItem(uniqueId, treeUniqueIds)
}

// 在上下文中根据某一元素查找下文元素（包括自己）
func (a *ParserData) findUniqueIdByUnderUniqueIdFromLogics(uniqueId string) []string {
	res := make([]string, 0)
	flag := false
	for _, v := range a.Layouts {
		if flag {
			res = append(res, v.UniqueID)
		}
		if uniqueId == v.UniqueID {
			flag = true
		}
	}
	return res
}

// 在上下文中根据某一元素查找上文元素（包括自己）
func (a *ParserData) findUniqueIdByAboveUniqueIdFromLogics(uniqueId string) []string {
	res := make([]string, 0)
	for i, v := range a.Layouts {
		if uniqueId == v.UniqueID {
			for j := 0; j < i; j++ {
				res = append(res, a.Layouts[j].UniqueID)
			}
			return res
		}
	}
	return res
}

// GetContextResultsByUIds 根据UIds获取包转结果
func (a *ParserData) GetContextResultsByUIds(ids []string, t string) (result.Result, error) {
	contexts := make([]result.Context, 0)
	groups := make([]map[string]result.Context, 0)
	for _, layout := range a.Layouts {
		if (ids == nil || utils.InString(layout.UniqueID, ids)) && layout.Type == t {
			if layout.Type == TypeTable {
				g := layout.TableToResultGroup()
				groups = append(groups, g...)
			} else {
				c := layout.ToResultContext()
				contexts = append(contexts, *c)
			}
		}
	}
	return result.Result{
		ID:        utils.GenerateUUID(),
		FieldType: result.FieldTypeAll,
		Contexts:  contexts,
		Groups:    groups,
	}, nil
}

// GetPageInfoByPageNum 获取页面信息
func (a *ParserData) GetPageInfoByPageNum(number int) (height int, width int, angle float64) {
	pages := a.DocInfo.Pages
	if len(pages) <= number {
		return -1, -1, 0
	}
	page := pages[number]
	return page.ImageHeight, page.ImageWidth, page.Angle
}
