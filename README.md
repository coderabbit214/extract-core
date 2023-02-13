# extract-core

![go](https://img.shields.io/badge/Go-v1.19-blue)

## 关于 extract-core

`extract-core`是基于[阿里云文档智能](https://www.aliyun.com/product/ai/docmind)的规则提取引擎。

## 使用

### 安装

```bash
 go get github.com/coderabbit214/extract-core
```

### 用例

> [用例文件地址](https://github.com/coderabbit214/extract-core/tree/main/example)

#### HelloWorld

> `NewExtract(dictionary dictionary.Dictionary, parserData data.ParserData, onExtractEnd func(map[string]result.Result))`参数说明:
>
> - dictionary.Dictionary:抽取字典，更多说明查看下文。
> - data.ParserData:阿里云文档智能返回的解析结果。
> - func(map[string]result.Result):抽取结束后，对抽取结果进行处理。

```go
func main() {
	content, err := os.ReadFile("extract-core.json")
	if err != nil {
		fmt.Println("json read err:", err.Error())
		return
	}
	value, err := data.NewData(string(content))
	if err != nil {
		fmt.Println("NewData err:", err.Error())
		return
	}
	yamlFile, err := os.ReadFile("hello.yaml")
	if err != nil {
		fmt.Println("yaml read err:", err.Error())
		return
	}
	dic, err := dictionary.NewDictionaryByYaml(string(yamlFile))
	if err != nil {
		fmt.Println("NewDictionary err:", err.Error())
		return
	}
	e, err := extract.NewExtract(dic, value, nil)
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
```

#### 字典

##### 示例

```yaml
name: extract-core
fields:
  - name: text
    type: text
    field-type: context
    regular: 表决.*(?P<val>通过|不通过|否决|弃权)
    filters:
      - kind: parent
        type: title
        regular: .*第三部分
      - kind: above
        type: text
        regular: 第三段
  - name: group
    type: text
    field-type: context
    regular: (?P<name>.*)[:：](?P<email>[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?)
    filters:
      - kind: parent
        type: title
        regular: .*第一部分
  - name: table
    type: table
    field-type: context
    filters:
      - kind: parent
        type: title
        regular: .*第二部分
  - name: cells
    field-type: cells
    cells-regulars:
      - kind: row
        regular: CC
      - kind: col
        regular: 身高
```

##### 说明

主要分为两部分：

第一部分：`name`，字典唯一标识。

第二部分：`fields`，提取字段以及对应规则。

- `name`：字段名称，必填
- `field-type`：提取方式，目前支持以下两种
  - `context`：段落中提取
  - `cells`：表格中提取单元格
- `type`：提取数据类型，详情查看[阿里云文档智能](https://www.aliyun.com/product/ai/docmind)说明文档
- `regular`：提取数据正则，在一段话中提取一个结果使用`(?P<val>)`，提取多个使用多个名称，两种返回数据结构不同
- `filters`：数据过滤，多种过滤方式以及对应的`type`和`regular`
  - `parent`：在哪个标题下
  - `under`：上文
  - `above`：下文
- `cells-regulars`：`field-type`为`cells`时使用，单元格内过滤，支持`regular`
  - `row`：行过滤
  - `col`：列过滤

## License

Easegress is under the Apache 2.0 license. See the [LICENSE](https://github.com/coderabbit214/easegress/blob/main/LICENSE) file for details.

