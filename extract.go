package extract

import (
	"github.com/coderabbit214/extract-core/data"
	"github.com/coderabbit214/extract-core/dictionary"
	"github.com/coderabbit214/extract-core/result"
	"github.com/coderabbit214/extract-core/utils"
)

type Extract struct {
	ID           string
	Dictionary   dictionary.Dictionary
	ParserData   data.ParserData
	Result       map[string]result.Result
	OnExtractEnd func(map[string]result.Result)
}

func NewExtract(dictionary dictionary.Dictionary, parserData data.ParserData, onExtractEnd func(map[string]result.Result)) (*Extract, error) {
	extract := &Extract{
		ID:           utils.GenerateUUID(),
		Dictionary:   dictionary,
		ParserData:   parserData,
		OnExtractEnd: onExtractEnd,
	}
	extract.extract()
	return extract, nil
}

// extract 提取
func (m *Extract) extract() {
	res := make(map[string]result.Result)
	fields := m.Dictionary.Fields
	rangeMap := make(map[string][]utils.Range)
	for _, field := range fields {
		//过滤
		filters := field.Filters
		var uIds []string
		if len(filters) > 0 {
			uIds = make([]string, 0)
			ranges, _ := m.ParserData.GetUniqueIdsByField(filters, rangeMap)
			rangeMap[field.Name] = ranges
			for _, ran := range ranges {
				uIds = append(uIds, ran.UniqueIds...)
			}
		}
		//查找
		var r result.Result
		switch field.FieldType {
		case dictionary.FieldTypeContext:
			if field.Regular != "" {
				r = m.ParserData.ContextExtract(uIds, field.Regular, field.Type)
			} else {
				r, _ = m.ParserData.GetContextResultsByUIds(uIds, field.Type)
			}
		case dictionary.FieldTypeCells:
			r = m.ParserData.CellsExtract(uIds, field.CellsRegulars)
		}
		res[field.Name] = r
	}
	m.Result = res

	if m.OnExtractEnd != nil {
		m.OnExtractEnd(m.Result)
	}
}
