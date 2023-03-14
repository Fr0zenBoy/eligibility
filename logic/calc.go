package logic

func GetPercentege(amount, limit float64) float64 {
	return ((float64(amount) / float64(limit)) * 100)
}
