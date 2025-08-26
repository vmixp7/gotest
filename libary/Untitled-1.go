
type Interval struct {
	Name string
	Draw int
}

func ffffff() {
	participants := []Interval{
		{"A", 1},
		{"B", 8},
		{"C", 5},
	}
	rand.Seed(time.Now().UnixNano())
	weightSum := 0
	for _, p := range participants {
		weightSum += p.Draw
	}
	index := rand.Intn(len(weightSum))
	for i, p := range participants {
		if index < p.Draw {
			fmt.Println("中獎者是:", participants[index])
			break
		}
		index -= p.Draw
	}

}