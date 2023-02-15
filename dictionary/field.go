package dictionary

// Field 过滤字段
type Field struct {
	Name          string         `yaml:"name" json:"name"`
	FieldType     string         `yaml:"field-type" json:"fieldType"`
	Type          string         `yaml:"type" json:"type"`
	Regular       string         `yaml:"regular" json:"regular"`
	CellsRegulars []CellsRegular `yaml:"cells-regulars" json:"cellsRegulars"`
	Filters       []Filter       `yaml:"filters" json:"filters"`
}

const (
	FieldTypeContext = "context"
	FieldTypeCells   = "cells"
)
