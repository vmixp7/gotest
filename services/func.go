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
func counter() func() int {
	// `count` 是外部函數的局部變數
	count := 0

	// 這裡返回的是一個匿名函數（內層函數）
	// 這個內層函數「捕獲」了外部函數的 `count` 變數
	return func() int {
		count++ // 每次呼叫這個內層函數，`count` 的值都會增加
		return count
	}
}

func MakeCounter() {
	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3
}

// CounterFunc 定義計數器返回的函式型別
type CounterFunc func(delta int, reset bool) int

// counter 返回一個閉包，可以累計、加 N 或重置
func counter2() CounterFunc {
	count := 0
	return func(delta int, reset bool) int {
		if reset {
			count = 0
		} else {
			count += delta
		}
		return count
	}
}

func MakeCounter2() {
	c := counter2()

	fmt.Println(c(1, false)) // +1 → 1
	fmt.Println(c(1, false)) // +1 → 2
	fmt.Println(c(5, false)) // +5 → 7
	fmt.Println(c(0, true))  // 重置 → 0
	fmt.Println(c(2, false)) // +2 → 2
}
