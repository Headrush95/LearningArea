// Package SliceSort содержит различные способы сортировки слайсов с типом элементов integer
package SliceSort

import (
	"errors"
)

var emptySliceErr = errors.New("slice is empty")

// InsertionSort реализует сортировку вставкой для слайсов с данными типа integer
func InsertionSort(arr []int) (err error) {
	// проверяем на нулевой массив
	if len(arr) == 0 {
		err = emptySliceErr
		return
	}

	for notSorted := 1; notSorted < len(arr); notSorted++ {
		// запоминаем следующее неотсортированное число
		current := arr[notSorted]

		lessSorted := 0 // индекс отсортированного числа, которое меньше current
		// ищем место в отсортированной части массива
		for arr[lessSorted] < current {
			lessSorted++
		}

		// смещаем все значения в отсортированной части, чтобы вставить новое
		for k := notSorted; k > lessSorted; k-- {
			arr[k] = arr[k-1]
		}
		arr[lessSorted] = current
	}
	return
}

// SelectionSort реализует сортировку выбором для слайсов с данными типа integer
func SelectionSort(arr []int) (err error) {
	// проверяем на нулевой массив
	if len(arr) == 0 {
		err = emptySliceErr
		return
	}

	var min int
	// идем по неотсортированной части массива
	for sorted := 0; sorted < len(arr)-1; sorted++ {
		min = sorted
		// ищем минимальное значение в неотсортированной части
		for notSorted := sorted + 1; notSorted < len(arr); notSorted++ {
			if arr[min] > arr[notSorted] {
				min = notSorted
			}
		}
		// меняем местами минимальное значение в неотсортированной части и текущее
		arr[sorted], arr[min] = arr[min], arr[sorted]
	}

	return
}

// BubbleSort реализует классическую пузырьковую сортировку для слайсов с данными типа integer
func BubbleSort(arr []int) (err error) {
	// проверяем на нулевой массив
	if len(arr) == 0 {
		err = emptySliceErr
		return
	}

	notSorted := true

	for notSorted {
		notSorted = false

		for idx := 0; idx < len(arr)-1; idx++ {
			if arr[idx] > arr[idx+1] {
				arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
				notSorted = true
			}
		}
	}

	return
}

// BubbleSortMod модифицированная версия классической пузырьковой сортировки: проход по массиву осуществляется
// с двух направлений и запоминаются верхняя и нижняя границы уже отсортированных частей массива
func BubbleSortMod(arr []int) (err error) {
	// проверяем на нулевой массив
	if len(arr) == 0 {
		err = emptySliceErr
		return
	}

	notSorted := true

	// нижняя граница, отделяет отсортированную часть снизу
	lowerBound := 0
	// верхняя граница, отделяет отсортированную часть сверху
	upperBound := len(arr) - 1
	for notSorted {
		notSorted = false
		tempUpperBound := upperBound
		tempLowerBound := lowerBound

		// проходимся "снизу вверх"
		for idx := lowerBound; idx < upperBound; idx++ {
			if arr[idx] > arr[idx+1] {
				arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
				tempUpperBound = idx // обновляем потенциальную верхнюю границу
				notSorted = true
			}
		}

		// обновляем верхнюю границу
		upperBound = tempUpperBound

		// если массив уже отсортирован, то проход в другом направлении не требуется
		if !notSorted {
			break
		}

		// проходимся "сверху вниз"
		for idx := upperBound - 1; idx >= lowerBound; idx-- {
			if arr[idx] > arr[idx+1] {
				arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
				tempLowerBound = idx + 1 // обновляем потенциальную нижнюю границу
				notSorted = true
			}
		}
		// обновляем нижнюю границу
		lowerBound = tempLowerBound
	}

	return
}
