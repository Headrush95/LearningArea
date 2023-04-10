package Chapter_4

// FindMedian находит медиану одномерного массива
func FindMedian(arr []int) float64 {
	length := len(arr)
	if len(arr)%2 != 0 {
		return float64(arr[length/2])
	}

	return (float64(arr[length/2]) + float64(arr[length/2-1])) / 2
}
