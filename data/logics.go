package data

type Logics struct {
	DocTree      []DocTree     `json:"docTree"`
	ParagraphKVs []ParagraphKV `json:"paragraphKVs"`
	TableKVs     []TableKV     `json:"tableKVs"`
}

type DocTree struct {
	Backlink Backlink `json:"backlink"`
	Level    int      `json:"level"`
	Link     Link     `json:"link"`
	UniqueID string   `json:"uniqueId"`
}

type Backlink struct {
	Superior []string `json:"上级"`
}
type Link struct {
	Subordinate []interface{} `json:"下级"`
	Include     []interface{} `json:"包含"`
}
type ParagraphKV struct {
	ExtInfo ExtInfo  `json:"extInfo"`
	Key     []string `json:"key"`
	Value   []string `json:"value"`
}
type ExtInfo struct {
	KeyConfidence   float64 `json:"keyConfidence"`
	KeyLayoutID     string  `json:"keyLayoutId"`
	ValueConfidence float64 `json:"valueConfidence"`
	ValueLayoutID   string  `json:"valueLayoutId"`
}
type TableKV struct {
	CellIDRelations []CellIDRelation `json:"cellIdRelations"`
	KvInfo          []KvInfo         `json:"kvInfo"`
	KvListInfo      [][]KvListInfo   `json:"kvListInfo"`
}

type CellIDRelation struct {
	ExtInfo  TableExtInfo `json:"extInfo"`
	Key      []string     `json:"key"`
	KeyPos   interface{}  `json:"keyPos"`
	Value    []string     `json:"value"`
	ValuePos interface{}  `json:"valuePos"`
}
type KvInfo struct {
	ExtInfo  TableExtInfo `json:"extInfo"`
	Key      []string     `json:"key"`
	KeyPos   interface{}  `json:"keyPos"`
	Value    []string     `json:"value"`
	ValuePos interface{}  `json:"valuePos"`
}
type KvListInfo struct {
	ExtInfo  TableExtInfo `json:"extInfo"`
	Key      []string     `json:"key"`
	KeyPos   interface{}  `json:"keyPos"`
	Value    []string     `json:"value"`
	ValuePos interface{}  `json:"valuePos"`
}

type TableExtInfo struct {
	TableID string `json:"table_id"`
}
