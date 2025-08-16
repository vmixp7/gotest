package services

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func InterfaceTest() {
	c := Circle{Radius: 5}
	var s Shape = c
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

// 定义接口
type Speaker interface {
	Speak()
}

// 父结构体
type Animal struct {
	Name string
}

// 实现接口方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal
	Breed string
}

func InterfaceSpeaker() {
	var speaker Speaker

	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	speaker = &dog
	speaker.Speak() // 通过接口调用方法
}
