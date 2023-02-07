package dictionary

type CellsRegular struct {
	Kind    string `yaml:"kind"`
	Regular string `yaml:"regular"`
}

const (
	KindRow = "row"
	KindCol = "col"
)
