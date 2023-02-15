package dictionary

type CellsRegular struct {
	Kind    string `yaml:"kind" json:"kind"`
	Regular string `yaml:"regular" json:"regular"`
}

const (
	KindRow = "row"
	KindCol = "col"
)
