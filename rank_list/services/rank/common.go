package rank

import "math"

const MaxPointPart = 9999999999999 //13位的9，用于时间戳转小数的计算
func calculateScore(score float64, timestamp int64) float64 {
	return  score + float64(MaxPointPart - timestamp) * 0.0000000000001
}

func reductionScore(baseScore float64) float64 {
	return math.Trunc(baseScore)
}
