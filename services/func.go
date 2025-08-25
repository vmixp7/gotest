package services

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function當值
func FuncTest1() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// 閉包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func FuncTest2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		// 返回当前值并更新 a, b
		result := a
		a, b = b, a+b
		return result
	}
}
func FuncTest3() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 泛型函數
// T 是一個型別參數，any 代表可以是任何型別
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Generic() {
	ints := []int{1, 2, 3}
	strs := []string{"Go", "泛型", "讚"}

	PrintSlice(ints) // 自動推斷 T = int
	PrintSlice(strs) // 自動推斷 T = string
}

// 是外部函數，它會返回一個閉包
func MakeCounter() func() int {
	// `count` 是外部函數的局部變數
	count := 0

	// 這裡返回的是一個匿名函數（內層函數）
	// 這個內層函數「捕獲」了外部函數的 `count` 變數
	return func() int {
		count++ // 每次呼叫這個內層函數，`count` 的值都會增加
		return count
	}
}
