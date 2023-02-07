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
	err := extract.extract()
	if err != nil {
		return nil, err
	}
	return extract, nil
}

// extract 提取
func (m *Extract) extract() error {
	res := make(map[string]result.Result)
	rangeMap := make(map[string][]string)

	fields := m.Dictionary.Fields
	for _, field := range fields {
		//过滤
		filters := field.Filters
		if len(filters) > 0 {
			ranges, err := m.ParserData.GetUniqueIdsByField(filters, rangeMap)
			if err != nil {
				return err
			}
			rangeMap[field.Name] = ranges
		}
		//查找
		var r result.Result
		switch field.FieldType {
		case dictionary.FieldTypeContext:
			if field.Regular != "" {
				r = m.ParserData.ContextExtract(rangeMap[field.Name], field.Regular, field.Type)
			} else {
				r, _ = m.ParserData.GetContextResultsByUIds(rangeMap[field.Name], field.Type)
			}
		case dictionary.FieldTypeCells:
			r = m.ParserData.CellsExtract(rangeMap[field.Name], field.CellsRegulars)
		}
		res[field.Name] = r
	}
	m.Result = res

	if m.OnExtractEnd != nil {
		m.OnExtractEnd(m.Result)
	}
	return nil
}
