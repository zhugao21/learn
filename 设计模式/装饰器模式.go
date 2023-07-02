package designpatterns

// 装饰器模式主要解决继承关系过于负载的问题，通过组合代替继承，
// 主要作用给原始类增加新的功能

type IDraw interface {
	Draw() string
}

type Square struct {
}

// Draw 输出描述信息
func (s *Square) Draw() string {
	return "this is square"
}

type ColorSquare struct {
	square Square
	color  string
}

// Draw 输出在原有的基础上，加上颜色描述
func (c *ColorSquare) Draw() string {
	return c.square.Draw() + " color is " + c.color
}
