package data

type DocInfo struct {
	DocType string `json:"docType"`
	Pages   []Page `json:"pages"`
}
type Page struct {
	Angle           float64 `json:"angle"`
	ImageHeight     int     `json:"imageHeight"`
	ImageStorageKey string  `json:"imageStorageKey"`
	ImageType       string  `json:"imageType"`
	ImageURL        string  `json:"imageUrl"`
	ImageWidth      int     `json:"imageWidth"`
	PageIDAllDocs   int     `json:"pageIdAllDocs"`
	PageIDCurDoc    int     `json:"pageIdCurDoc"`
}
