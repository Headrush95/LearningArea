package RadixSort

import (
	"fmt"
	"strings"
)

func main2() {
	data := []string{
		"ABCD",
		"ABCA",
		"BBCD",
		"ACCD",
		"ABBD",
		"CADD",
		"DADA",
		"CADA",
	}

	//RadixSort(data, 3)
	//RadixSort(data, 2)
	//RadixSort(data, 1)
	RadixSort(data, 0)
	fmt.Println(strings.Join(data, "\n"))

}

func RadixSort(data []string, position int) {
	size := 4
	temp := make([]string, len(data), len(data))
	counters := make([]int, size+1, size+1)
	for i := 0; i < len(data); i++ {
		counters[data[i][position]-65+1]++
	}
	fmt.Println(counters)

	for i := 0; i < len(counters)-1; i++ {
		counters[i+1] += counters[i]
	}
	fmt.Println(counters)

	for i := 0; i < len(data); i++ {
		temp[counters[data[i][position]-65]] = data[i]
		counters[data[i][position]-65]++
	}

	for i := 0; i < len(data); i++ {
		data[i] = temp[i]
	}
}
