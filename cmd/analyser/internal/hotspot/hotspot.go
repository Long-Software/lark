package hotspot

func Calculate(complexity, lines int) float64 {
	complexityWeight := float64(complexity) * 2.0
	sizeWeight := float64(lines) * 0.1

	return complexityWeight + sizeWeight
}

type Entry struct {
	Path  string  `json:"path"`
	Score float64 `json:"score"`
}
