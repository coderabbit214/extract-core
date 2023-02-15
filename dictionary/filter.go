package dictionary

// Filter 过滤字段
type Filter struct {
	Kind    string `yaml:"kind" json:"kind"`
	Name    string `yaml:"name" json:"name"`
	Type    string `yaml:"type" json:"type"`
	Regular string `yaml:"regular" json:"regular"`
}

const (
	KindParent = "parent"
	KindUnder  = "under"
	KindAbove  = "above"
)
