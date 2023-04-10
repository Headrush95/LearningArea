package Chapter_4

import (
	"math"
)

// SampleVariance расчитывает выборочную дисперсию для одномерного массива
func SampleVariance(arr []int) float64 {
	totalTmp := 0
	for _, num := range arr {
		totalTmp += num
	}

	length := float64(len(arr))
	var average = float64(totalTmp) / length

	res := 0.00
	for _, num := range arr {
		res += math.Pow(float64(num)-average, 2)
	}
	res /= length
	return res
}
