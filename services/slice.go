package services

import (
	"fmt"
)

func arr_set(arr []int) {
	arr[0] = 123 // 因為 slice 的 header（長度、容量、指標）被複製，但底層陣列是共用的，所以修改會影響原始的底層陣列
}

func arr_append(arr []int) {
	arr = append(arr, 777) // 尝试添加一个新元素
	// 這個新的 slice 只在 arr_append 內部有效，因為參數傳遞是 值拷貝，外部的 arr 不會改變
	// 但如果你需要修改原始数组，你应该返回新的切片或使用指针。
	fmt.Println(arr)
}

func Slice1() {
	arr := []int{1, 2, 3}
	arr_set(arr)
	arr_append(arr)
	fmt.Println(arr)
}

func Slice2() {
	originalSlice := []int{1, 2, 3, 4, 5}
	// 'newSlice' 的 header(長度、容量、指標)被複製，但底層陣列與 'originalSlice' 相同
	newSlice := originalSlice

	fmt.Printf("Original Slice: %v\n", originalSlice)
	fmt.Printf("New Slice:      %v\n", newSlice)

	// 透過 'newSlice' 修改元素
	newSlice[0] = 99

	fmt.Println("\n--- After modification ---")
	// 'originalSlice' 也被改變了，因為它們共用底層陣列
	fmt.Printf("Original Slice: %v\n", originalSlice)
	fmt.Printf("New Slice:      %v\n", newSlice)
}

func Slice3() {
	originalSlice := []int{1, 2, 3, 4, 5}
	// 建立一個新的 slice，並複製其底層資料
	newSlice := make([]int, len(originalSlice))
	copy(newSlice, originalSlice)

	fmt.Printf("Original Slice: %v\n", originalSlice)
	fmt.Printf("New Slice:      %v\n", newSlice)

	// 透過 'newSlice' 修改元素
	newSlice[0] = 99

	fmt.Println("\n--- After modification ---")
	// 這次 'originalSlice' 不會受到影響
	fmt.Printf("Original Slice: %v\n", originalSlice)
	fmt.Printf("New Slice:      %v\n", newSlice)
}
