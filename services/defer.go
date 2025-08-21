package services

import (
	"fmt"
)

func foo() (n int) {
	defer func() {
		n += 5
		fmt.Println("a:", n) // 13 + 5 = 18
	}()

	defer func() {
		n += 3               // 10 + 3 = 13
		fmt.Println("b:", n) // defer 執行順序是 後進先出
	}()
	n = 10
	return n // 最後返回值是18
}

func DeferTest1() {
	fmt.Println("c:", foo())
}
