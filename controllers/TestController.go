package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 反轉一個字串
func FirstReverse(str string) string {
	runes := []rune(str) // 將字串轉換為rune切片以處理Unicode字符
	fmt.Println("Original string:", runes)
	// 雙指針法
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 判斷括號是否正確配對
func BracketMatcher(str string) int {
	stack := []rune{}
	for _, ch := range str {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
		fmt.Println("Current stack:", stack)
	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}

// 給一個數列，檢查是否存在任意子集合加總等於最大值
func ArrayAdditionI(arr []int) bool {
	target := 0
	maxIdx := 0
	for i, v := range arr {
		if v > target {
			target = v
			maxIdx = i
		}
	}
	// Remove max from array
	arr = append(arr[:maxIdx], arr[maxIdx+1:]...)
	n := len(arr)
	for i := 1; i < (1 << n); i++ { // 1<<n 為 2^n，表示所有可能的子集合
		sum := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 { // 檢查第j位是否為1
				sum += arr[j]
				fmt.Println("j:", arr[j])
			}
		}
		fmt.Println("Current combination sum:", sum)
		if sum == target {
			return true
		}
	}
	return false
}

// Codeland用戶名驗證
func CodelandUsernameValidation(str string) string {
	// 長度限制
	if len(str) < 4 || len(str) > 25 {
		return "false"
	}

	// 必須以字母開頭
	startsWithLetter := regexp.MustCompile(`^[a-zA-Z]`)
	if !startsWithLetter.MatchString(str) {
		return "false"
	}

	// 只能包含字母、數字、底線
	validChars := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validChars.MatchString(str) {
		return "false"
	}

	// 不能以底線結尾
	if str[len(str)-1] == '_' {
		return "false"
	}

	return "true"
}

// 在排序陣列中尋找和為目標值的兩個數 (相向指針)
func twoSum(numbers []int, target int) []int {
	// 初始化左右指針
	left := 0
	right := len(numbers) - 1

	// 當 left 指針小於 right 指針時，繼續搜尋
	for left < right {
		currentSum := numbers[left] + numbers[right]

		// 如果當前和等於目標值，返回索引（題目要求從 1 開始）
		if currentSum == target {
			return []int{left + 1, right + 1}
		} else if currentSum < target {
			// 如果和太小，左指針向右移動
			left++
		} else {
			// 如果和太大，右指針向左移動
			right--
		}
	}

	// 如果沒有找到（根據題目假設，這不應該發生，但作為函數的完整性，通常會這麼處理）
	return []int{}
}

func GetTest(c *gin.Context) {
	// ch := make(chan int)

	// // 啟動一個Goroutine，將數字從1到10傳遞到通道中

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// // 在主Goroutine中讀取通道中的數字，並將它們加總

	// sum := 0
	// for num := range ch {
	// 	sum += num
	// }
	// fmt.Println("The sum is:", sum)

	numbers1 := []int{2, 7, 11, 15}
	target1 := 9

	data := twoSum(numbers1, target1)

	c.JSON(http.StatusOK, data)
}
