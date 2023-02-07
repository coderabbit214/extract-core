package data

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
