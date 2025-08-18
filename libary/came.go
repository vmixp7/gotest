package libary

import (
	"math/rand"
	"time"
)

type Participant struct {
	Name   string
	Weight int
}

// 平均機率隨機抽獎
// names := []string{"小明", "小美", "小王"}
func RandomDraw(participants []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(participants))
	return participants[index]
}

// 加權抽獎
//
//	participants := []libary.Participant{
//		{"Alice", 1},
//		{"Bob", 3},
//		{"Charlie", 6},
//	}
func WeightedDraw(participants []Participant) string {
	// rand.Seed(time.Now().UnixNano())
	// totalWeight := 0
	// for _, p := range participants {
	// 	totalWeight += p.Weight
	// }
	// r := rand.Intn(totalWeight)
	// for _, p := range participants {
	// 	fmt.Println("r-------------------------", r)
	// 	if r < p.Weight {
	// 		return p.Name
	// 	}
	// 	r -= p.Weight
	// 	fmt.Println("Weight------", p.Weight)
	// }
	// return ""

	rand.Seed(time.Now().UnixNano())
	weight := 0
	for _, p := range participants {
		weight += p.Weight
	}
	index := rand.Intn(weight)
	for _, p := range participants {
		if index < p.Weight {
			return p.Name
		}
		index -= p.Weight
	}
	return ""
}

// 不重複隨機抽出3位
// participants := []string{"A", "B", "C", "D", "E"}
// count :=3
func ShuffleAndPick(participants []string, count int) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(participants), func(i, j int) { //洗牌
		participants[i], participants[j] = participants[j], participants[i]
	})
	if count > len(participants) {
		count = len(participants)
	}
	return participants[:count]
}

// 機率抽中（例如 10%）
// probability := 0.1
func ChanceDraw(probability float64) bool {
	rand.Seed(time.Now().UnixNano())
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
