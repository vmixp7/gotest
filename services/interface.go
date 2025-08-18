package services

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，不过我们并不需要显式声明这一点。
func (t T) M() {
	fmt.Println(t.S)
}

func InterfaceTest1() {
	var i I = T{"hello"}
	i.M()
}

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

func InterfaceTest2() {
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

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func InterfaceTest3() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = struct {
		Name string
		Age  int
	}{
		Name: "Alice", Age: 30,
	}

	i = map[string]int{"one": 1, "two": 2}
	describe(i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 长度为 %v 字节\n", v, len(v))
	default:
		fmt.Printf("我不知道类型 %T!\n", v)
	}
}

func InterfaceTest4() {
	do(21)
	do("hello")
	do(true)
}

type Cat struct {
	name string
}

func InterfaceTest5() {
	var i interface{}
	fmt.Println(i == nil) // i 尚未被賦值 → interface 預設值是 nil（type 與 data 都是 nil）

	var p *Cat = nil
	i = p
	fmt.Println(i == nil) // i 不是 nil，因為 type 欄位已經是 *Cat，即使 value 是 nil
}
