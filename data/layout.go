package data

import (
	"github.com/coderabbit214/extract-core/dictionary"
	"github.com/coderabbit214/extract-core/result"
	"github.com/coderabbit214/extract-core/utils"
	"math"
	"regexp"
)

type Layout struct {
	// 版面信息唯一id
	UniqueID string `json:"uniqueId"`
	// 版面阅读顺序
	Index int `json:"index"`
	// 版面类型
	Type string `json:"type"`
	// 版面子类型
	SubType string `json:"subType"`
	// 文本
	Text string `json:"text"`
	// 间距枚举 center left right
	Alignment string `json:"alignment"`
	// 坐标
	Pos []utils.Point `json:"pos"`
	// 版面所在页数（可能多页）
	PageNum []int `json:"pageNum"`
	// 行平均高度（段落版面特有）
	LineHeight int `json:"lineHeight,omitempty"`
	// 文字首行缩进 （段落版面特有）
	FirstLinesChars int `json:"firstLinesChars,omitempty"`
	// 子数据
	Blocks []Block `json:"blocks"`
	// 单元格信息（类型是表格才有）
	Cells []Cell `json:"cells,omitempty"`
	// 表格总行数（类型是表格才有）
	NumCol int `json:"numCol,omitempty"`
	// 表格总列数（类型是表格才有）
	NumRow int `json:"numRow,omitempty"`
	// ParserData
	ParserData *ParserData
}

const (
	TypeTable = "table"
)

// TableCellsRegular 表格内部筛选
func (l *Layout) TableCellsRegular(regulars []dictionary.CellsRegular) []int {
	var xCellIds []int
	var yCellIds []int
	//单元格根据CellsRegular过滤
	for _, regular := range regulars {
		reg := regexp.MustCompile(regular.Regular)
		switch regular.Kind {
		case dictionary.KindCol:
			yCellIds = make([]int, 0)
			for _, cell := range l.Cells {
				if cell.Ysc != 0 {
					continue
				}
				text := cell.GetText()
				if reg.MatchString(text) {
					yCellIds = l.GetColCellIds(cell.Xsc)
				}
			}
		case dictionary.KindRow:
			xCellIds = make([]int, 0)
			for _, cell := range l.Cells {
				if cell.Xsc != 0 {
					continue
				}
				text := cell.GetText()
				if reg.MatchString(text) {
					xCellIds = l.GetRowCellIds(cell.Ysc)
				}
			}
		}
	}
	var cellIds []int
	if yCellIds == nil || xCellIds == nil {
		cellIds = utils.UnionInts(xCellIds, yCellIds)
	} else {
		cellIds = utils.IntersectionInts(xCellIds, yCellIds)
	}
	return cellIds
}

// GetCellsByIds 根据id获取Cells
func (l *Layout) GetCellsByIds(cellIds []int) []Cell {
	cells := make([]Cell, 0)
	for _, cell := range l.Cells {
		if utils.InInt(cell.CellID, cellIds) {
			cells = append(cells, cell)
		}
	}
	return cells
}

// GetRowCellIds 获取同行Cells
func (l *Layout) GetRowCellIds(y int) []int {
	ids := make([]int, 0)
	for _, cell := range l.Cells {
		if cell.Ysc >= y && cell.Yec <= y {
			ids = append(ids, cell.CellID)
		}
	}
	return ids
}

// GetColCellIds 获取同列Cells
func (l *Layout) GetColCellIds(x int) []int {
	ids := make([]int, 0)
	for _, cell := range l.Cells {
		if cell.Xsc >= x && cell.Xec <= x {
			ids = append(ids, cell.CellID)
		}
	}
	return ids
}

// ToResultContext 上下文转换为Context
func (l *Layout) ToResultContext() *result.Context {
	dataBlocks := l.Blocks
	blocks := make([]result.Block, 0)
	for _, dataBlock := range dataBlocks {
		b := result.Block{
			Pos:        dataBlock.Pos,
			Text:       dataBlock.Text,
			PageNumber: l.PageNum[0],
		}
		blocks = append(blocks, b)
	}
	context := &result.Context{
		Blocks: blocks,
		Text:   l.Text,
	}
	return context
}

// TableToResultGroup Table转换为Group
func (l *Layout) TableToResultGroup() []map[string]result.Context {
	groups := make([]map[string]result.Context, l.NumRow)

	heads := make([]Cell, 0)
	for _, cell := range l.Cells {
		if cell.Ysc == 0 {
			heads = append(heads, cell)
		}
	}

	for _, cell := range l.Cells {
		if cell.Ysc == 0 {
			continue
		}
		m := groups[cell.Ysc]
		if m == nil {
			m = make(map[string]result.Context)
		}
		for _, head := range heads {
			if cell.Xsc == head.Xsc && cell.Xec == head.Xec {
				m[head.GetText()] = result.Context{
					Blocks: cell.GetBlocks(l.ParserData),
					Text:   cell.GetText(),
				}
			}
		}
		groups[cell.Ysc] = m
	}

	for i := 0; i < len(groups); i++ {
		if groups[i] == nil {
			groups = append(groups[:i], groups[i+1:]...)
		}
	}

	return groups
}

// ToResultContextByRegularAndGroupNumber 正则过滤上下文Layout转换为Context
func (l *Layout) ToResultContextByRegularAndGroupNumber(regular string, groupNumber int) (*[]result.Context, bool) {
	reg := regexp.MustCompile(regular)
	contexts := make([]result.Context, 0)
	submatch := reg.FindAllStringSubmatch(l.Text, -1)
	indexs := reg.FindAllStringSubmatchIndex(l.Text, -1)
	for i, sub := range submatch {
		text := sub[groupNumber]
		blocks := l.getBlocks(indexs[i][groupNumber*2], indexs[i][groupNumber*2+1])
		context := &result.Context{
			Blocks: *blocks,
			Text:   text,
		}
		contexts = append(contexts, *context)
	}
	return &contexts, true
}

func (l *Layout) getBlocks(begin int, end int) *[]result.Block {
	blocks := make([]result.Block, 0)
	dataBlocks := l.Blocks
	temp := 0
	for _, dataBlock := range dataBlocks {
		text := dataBlock.Text
		le := len(text)
		e := temp + le
		if end < temp {
			continue
		}
		if e < begin {
			break
		}
		minEnd := int(math.Min(float64(end), float64(e)))
		maxBegin := int(math.Max(float64(begin), float64(temp)))
		str := text[maxBegin-temp : minEnd-temp]
		pos := utils.CalculationResultPos(dataBlock.Pos, le, maxBegin-temp, minEnd-temp)
		height, width, angle := l.ParserData.GetPageInfoByPageNum(l.PageNum[0])
		b := &result.Block{
			Pos:        pos,
			Text:       str,
			PageNumber: l.PageNum[0],
			Height:     height,
			Width:      width,
			Angle:      angle,
		}
		blocks = append(blocks, *b)
		temp += le
	}
	return &blocks
}

// ToResultGroupByRegular 分组正则过滤上下文Layout转换为Context
func (l *Layout) ToResultGroupByRegular(regular string) (*[]map[string]result.Context, bool) {
	reg := regexp.MustCompile(regular)
	groupss := make([]map[string]result.Context, 0)
	if !reg.MatchString(l.Text) {
		return &groupss, false
	}
	groupNames := reg.SubexpNames()
	submatchs := reg.FindAllStringSubmatch(l.Text, -1)
	indexs := reg.FindAllStringSubmatchIndex(l.Text, -1)
	for j, submatch := range submatchs {
		groups := make(map[string]result.Context)
		for i := 1; i < len(submatch); i++ {
			blocks := l.getBlocks(indexs[j][i*2], indexs[j][i*2+1])
			group := &result.Context{
				Text:   submatchs[j][i],
				Blocks: *blocks,
			}
			groups[groupNames[i]] = *group
		}
		groupss = append(groupss, groups)
	}
	return &groupss, true
}
