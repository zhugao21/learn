package designpatterns

// 简单工厂：创建相同接口的不同类对象

type SimpleParser interface {
	Parser([]byte)
}

type jsonSimpleParser struct{}

func (j *jsonSimpleParser) Parser([]byte) {

}

type yamlSimpleParser struct{}

func (y *yamlSimpleParser) Parser([]byte) {

}

func NewParser(t string) SimpleParser {
	switch t {
	case "json":
		return &jsonSimpleParser{}
	case "yaml":
		return &yamlSimpleParser{}
	}
	panic("type error.")
}

// 工厂方法：
