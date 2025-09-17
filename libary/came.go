package libary

import (
	"fmt"
	"math/rand"
	"time"
)

// 平均機率隨機抽獎
// names := []string{"小明", "小美", "小王"}
func RandomDraw() string {
	participants := []string{"小明", "小美", "小王"}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(participants))
	fmt.Println("中獎者是:", participants[index])
	return participants[index]
}

// 加權抽獎
type participant struct {
	Name   string
	Weight int
}

func WeightedDraw() string {
	participants := []participant{
		{"Alice", 1},
		{"Bob", 3},
		{"Charlie", 6},
	}

	rand.Seed(time.Now().UnixNano())
	weight := 0
	// 計算總權重
	for _, p := range participants {
		weight += p.Weight
	}
	index := rand.Intn(weight)
	fmt.Println("index1--", index)
	// 根據權重選擇中獎者
	for _, p := range participants {
		if index < p.Weight {
			fmt.Println("winner--", p.Name)
			return p.Name
		}
		// 減去當前參與者的權重，繼續檢查下一個
		index -= p.Weight
		fmt.Println("index2--", index)
	}

	return ""
}

// 不重複隨機抽出3位
// participants := []string{"A", "B", "C", "D", "E"}
// count :=3
func ShuffleAndPick() []string {
	participants := []string{"A", "B", "C", "D", "E"}
	count := 3
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(participants), func(i, j int) { //洗牌
		// 交換位置
		participants[i], participants[j] = participants[j], participants[i]
	})
	// 確保不超過參與者數量
	if count > len(participants) {
		count = len(participants)
	}
	fmt.Println("中獎者是:", participants[:count])
	return participants[:count]
}

// 機率抽中（例如 10%）
// probability := 0.1
func ChanceDraw() bool {
	probability := 0.1
	rand.Seed(time.Now().UnixNano())
	fmt.Println("是否中獎:", rand.Float64() < probability)
	return rand.Float64() < probability
}

// 保底抽獎（N 次必中）
// 8(次數), 0.1, 10(保底次數)
func PityDraw(currentAttempts int, baseChance float64, pityLimit int) bool {
	if currentAttempts >= pityLimit {
		return true
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() < baseChance
}
