package designpatterns

import "fmt"

// 建造者模式：用于创建同一类的参数复杂的对象

type Person struct {
	Name string
	Age  int
}

type PersonBuilder struct {
	Name string
	Age  int
}

func (b *PersonBuilder) SetName(name string) {
	b.Name = name
}

func (b *PersonBuilder) SetAge(age int) {
	b.Age = age
}

func (b *PersonBuilder) Builder() (*Person, error) {
	if b.Name == "" {
		b.Name = "jack"
	}

	if b.Age == 0 {
		b.Age = 18
	}

	if b.Age < 0 || b.Age > 100 {
		return nil, fmt.Errorf("age(%d) invalid", b.Age)
	}
	return &Person{
		Name: b.Name,
		Age:  b.Age,
	}, nil
}

// 使用option方式实现

type PersonOption struct {
	Name string
	Age  int
}

type PersonOptionFunc func(*PersonOption)

func NewPerson(name string, opts ...PersonOptionFunc) *Person {
	personOption := &PersonOption{}

	for _, opt := range opts {
		opt(personOption)
	}
	return &Person{
		Name: name,
		Age:  personOption.Age,
	}
}
