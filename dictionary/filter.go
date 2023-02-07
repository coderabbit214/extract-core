package dictionary

// Filter 过滤字段
type Filter struct {
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Regular string `yaml:"regular"`
}

const (
	KindParent = "parent"
	KindUnder  = "under"
	KindAbove  = "above"
)
