package services

import (
	"fmt"
	"math"
)

/* 交换函数这样写更加简洁，也是 go 语言的特性，可以用下，c++ 和 c# 是不能这么干的 */
func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func PointerTest() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

func PointerTest2() {
	person := &Person{Name: "Alice", Age: 30}
	fmt.Println("Before:", person.Name) // Alice
	person.UpdateName("Bob")
	fmt.Println("After:", person.Name) // Bob
}

// 建立 person type
type person2 struct {
	firstName string
	lastName  string
}

// 建立 person 的 function receiver
func (p person2) updateName(newFirstName string) {
	// 沒使用指針,不會更新 firstName
	p.firstName = newFirstName
}

// 建立 person 的 function receiver
func (p *person2) updateNameForPoint(newFirstName string) {
	// 使用指针接收者更新 firstName
	p.firstName = newFirstName
}

func (p person2) print() {
	fmt.Printf("Current person is: %+v\n", p)
}

func PointerTest3() {
	jim := person2{
		firstName: "Jim",
		lastName:  "Party",
	}

	fmt.Printf("Before person is: %+v\n", jim)

	// jim.updateName("Aaron") // 沒有使用指针接收者，firstName 不會被更新
	jim.updateNameForPoint("Aaron") // 使用指针接收者更新 firstName

	fmt.Printf("After person is: %+v\n", jim)
}

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerTest4() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
