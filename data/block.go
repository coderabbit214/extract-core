package data

import (
	"github.com/coderabbit214/extract-core/utils"
)

type Block struct {
	Pos     []utils.Point `json:"pos"`
	StyleID int           `json:"styleId"`
	Text    string        `json:"text"`
}
