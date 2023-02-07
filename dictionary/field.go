package dictionary

// Field 过滤字段
type Field struct {
	Name          string         `yaml:"name"`
	FieldType     string         `yaml:"field-type"`
	Type          string         `yaml:"type"`
	Regular       string         `yaml:"regular"`
	CellsRegulars []CellsRegular `yaml:"cells-regulars"`
	Filters       []Filter       `yaml:"filters"`
}

const (
	FieldTypeContext = "context"
	FieldTypeCells   = "cells"
)
