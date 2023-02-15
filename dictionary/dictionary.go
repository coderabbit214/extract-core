package dictionary

import "gopkg.in/yaml.v3"

// Dictionary 提取字典
type Dictionary struct {
	Name   string  `yaml:"name" json:"name"`
	Fields []Field `yaml:"fields" json:"fields"`
}

func NewDictionaryByYaml(str string) (Dictionary, error) {
	d := Dictionary{}
	err := yaml.Unmarshal([]byte(str), &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func NewDictionaryByJson(str string) (Dictionary, error) {
	d := Dictionary{}
	err := yaml.Unmarshal([]byte(str), &d)
	if err != nil {
		return d, err
	}
	return d, nil
}
