package libary

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
