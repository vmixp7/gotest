package libary

import (
	"fmt"
	"regexp"
	"sort"
	"sync"
	"time"
)

// 移除index得值
func RemoveIndex() []int {
	s := []int{10, 20, 30, 40, 50}
	index := 2
	return append(s[:index], s[index+1:]...)
}

// 反轉一個字串
func FirstReverse() string {
	str := "hello world!"
	runes := []rune(str) // 將字串轉換為rune切片以處理Unicode字符
	fmt.Println("Original string:", runes)
	// 雙指針法
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 判斷括號是否正確配對
func BracketMatcher() int {
	str := "(coder)(byte))"
	stack := []rune{}
	for _, ch := range str {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				return 0
			}
			// 彈出堆疊頂部元素
			stack = stack[:len(stack)-1]
		}
		fmt.Println("Current stack:", stack)
	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}

type Interval struct {
	Start, End int
}

// 計算某個人「課表」的空閒時間 (假設 0~24 小時)
func getFreeTimes(courses []Interval) []Interval {
	if len(courses) == 0 {
		return []Interval{{0, 24}}
	}

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].Start < courses[j].Start
	})

	free := []Interval{}
	start := 0
	for _, c := range courses {
		if start < c.Start {
			free = append(free, Interval{start, c.Start})
		}
		if start < c.End {
			start = c.End
		}
	}
	if start < 24 {
		free = append(free, Interval{start, 24})
	}
	fmt.Println("free-----", free)
	return free
}

// 求兩組時間區間的交集
func intersect(a, b []Interval) []Interval {
	res := []Interval{}
	i, j := 0, 0
	// 使用雙指標法
	for i < len(a) && j < len(b) {
		start := max(a[i].Start, b[j].Start)
		end := min(a[i].End, b[j].End)
		// 有交集
		if start < end {
			res = append(res, Interval{start, end})
		}
		// 移動結束時間較早的指標
		if a[i].End < b[j].End {
			i++
		} else {
			j++
		}
	}
	fmt.Println("res-----", res)
	return res
}

// 求多個人的交集
func intersectAll(allFree [][]Interval) []Interval {
	if len(allFree) == 0 {
		return nil
	}
	result := allFree[0]
	for i := 1; i < len(allFree); i++ {
		result = intersect(result, allFree[i])
		if len(result) == 0 {
			break
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Exam1() {
	// 假設一個班級 (可彈性增加/減少人數)
	classCourses := [][]Interval{
		{{2, 4}, {12, 15}}, // 小明
		{{4, 5}},           // 大雄
		{{7, 8}},           // 小王
	}

	// 每個人計算空閒
	allFree := [][]Interval{}
	for _, courses := range classCourses {
		free := getFreeTimes(courses)
		allFree = append(allFree, free)
	}

	// 算出大家同時有空的時段
	common := intersectAll(allFree)

	// 輸出結果
	fmt.Print("所有人同時有空的時間: ")
	for _, c := range common {
		fmt.Printf("[%d~%d] ", c.Start, c.End)
	}
	fmt.Println()
}

func intersectIntervals(a, b [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		start := max(a[i][0], b[j][0])
		end := min(a[i][1], b[j][1])
		if start < end {
			res = append(res, []int{start, end})
		}
		if a[i][1] < b[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

// 取兩個陣列交集區間
func Exam2() {
	a := [][]int{{0, 2}, {5, 7}, {8, 12}}
	b := [][]int{{1, 3}, {6, 10}}

	intersections := intersectIntervals(a, b)
	fmt.Println("交集區間:", intersections)
}

// 給一個數列，檢查是否存在任意子集合加總等於最大值
func ArrayAdditionI() bool {
	arr := []int{4, 6, 23, 10, 1, 3}
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
func CodelandUsernameValidation() string {
	str := "u__hello_world"
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
// numbers1 := []int{2, 7, 11, 15}
// target1 := 9
func TwoSum() []int {
	numbers := []int{2, 7, 11, 15}
	target := 9
	// 初始化左右指針
	left := 0
	right := len(numbers) - 1

	// 當 left 指針小於 right 指針時，繼續搜尋
	for left < right {
		currentSum := numbers[left] + numbers[right]

		// 如果當前和等於目標值，返回索引（題目要求從 1 開始）
		if currentSum == target {
			fmt.Println("找到目標值，索引為:", []int{left + 1, right + 1})
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

type GameServer struct {
	actions chan string
	wg      *sync.WaitGroup
}

func NewGameServer(wg *sync.WaitGroup) *GameServer {
	return &GameServer{
		actions: make(chan string, 100), // 使用緩衝通道
		wg:      wg,
	}
}

func (s *GameServer) HandleActions() {
	defer s.wg.Done() // 確保在處理完成後減少WaitGroup計數
	for action := range s.actions {
		fmt.Println("Handling action:", action)
		// 在這裡處理遊戲邏輯
		time.Sleep(100 * time.Millisecond) // 模擬處理時間
	}
}

func Winners() {

	var wg sync.WaitGroup
	wg.Add(1) // 增加WaitGroup計數
	server := NewGameServer(&wg)
	go server.HandleActions()

	server.actions <- "Player1 joined the game"
	server.actions <- "Player2 joined the game"
	server.actions <- "Player3 moved to position (10, 20)"
	close(server.actions) // 關閉通道以結束處理

	wg.Wait() // 等待所有處理完成

	fmt.Println("winners")
}
