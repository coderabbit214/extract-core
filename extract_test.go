package extract

import (
	"encoding/json"
	"fmt"
	"github.com/coderabbit214/extract-core/data"
	"github.com/coderabbit214/extract-core/dictionary"
	"os"
	"testing"
)

func TestNewExtract(t *testing.T) {
	content, err := os.ReadFile("example/extract-core.json")
	if err != nil {
		fmt.Println("json read err:", err.Error())
		return
	}
	value, err := data.NewData(string(content))
	if err != nil {
		fmt.Println("NewData err:", err.Error())
		return
	}
	yamlFile, err := os.ReadFile("example/extract-core.yaml")
	if err != nil {
		fmt.Println("yaml read err:", err.Error())
		return
	}
	dic, err := dictionary.NewDictionaryByYaml(string(yamlFile))
	if err != nil {
		fmt.Println("NewDictionary err:", err.Error())
		return
	}
	e, err := NewExtract(dic, value, nil)
	if err != nil {
		fmt.Println("Extract err:", err.Error())
		return
	}
	marshal, err := json.Marshal(e.Result)
	if err != nil {
		fmt.Println("Marshal err:", err.Error())
		return
	}
	fmt.Println(string(marshal))
}
