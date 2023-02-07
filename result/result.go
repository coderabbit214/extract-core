package result

import "github.com/coderabbit214/extract-core/utils"

type Result struct {
	ID        string               `json:"id"`
	FieldType string               `json:"fieldType"`
	Contexts  []Context            `json:"context"`
	Groups    []map[string]Context `json:"groups"`
}

const (
	FieldTypeAll     = "all"
	FieldTypeContext = "context"
	FieldTypeGroup   = "group"
)

type Context struct {
	Blocks []Block `json:"blocks"`
	Text   string  `json:"text"`
}

type Block struct {
	Pos        []utils.Point `json:"pos"`
	Text       string        `json:"text"`
	PageNumber int           `json:"pageNumber"`
	Height     int           `json:"height"`
	Width      int           `json:"width"`
	Angle      float64       `json:"angle"`
}
