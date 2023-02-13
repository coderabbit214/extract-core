package data

import "github.com/coderabbit214/extract-core/result"

type Cell struct {
	Alignment string `json:"alignment"`
	CellID    int    `json:"cellId"`
	//全局id
	CellUniqueID string `json:"cellUniqueId"`
	//内嵌版面信息
	Layouts []Layout      `json:"layouts"`
	PageNum []int         `json:"pageNum"`
	Pos     []interface{} `json:"pos"`
	Type    string        `json:"type"`
	// 结束单元格横向是第几个
	Xec int `json:"xec"`
	// 起始单元格横向是第几个
	Xsc int `json:"xsc"`
	// 结束单元格纵向是第几个
	Yec int `json:"yec"`
	// 起始单元格纵向是第几个
	Ysc int `json:"ysc"`
}

func (c *Cell) GetText() string {
	text := ""
	for _, l := range c.Layouts {
		text += l.Text
	}
	return text
}

func (c *Cell) GetBlocks(parserData *ParserData) []result.Block {
	blocks := make([]result.Block, 0)
	for _, l := range c.Layouts {
		height, width, angle := parserData.GetPageInfoByPageNum(l.PageNum[0])
		for _, b := range l.Blocks {
			blocks = append(blocks, result.Block{
				Pos:        b.Pos,
				Text:       b.Text,
				PageNumber: l.PageNum[0],
				Height:     height,
				Width:      width,
				Angle:      angle,
			})
		}
	}

	return blocks
}
