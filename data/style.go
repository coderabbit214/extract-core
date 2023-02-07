package data

type Style struct {
	Bold       bool    `json:"bold"`
	CharScale  float64 `json:"charScale"`
	Color      string  `json:"color"`
	DeleteLine bool    `json:"deleteLine"`
	FontName   string  `json:"fontName"`
	FontSize   int     `json:"fontSize"`
	Italic     bool    `json:"italic"`
	StyleID    int     `json:"styleId"`
	Underline  bool    `json:"underline"`
}
