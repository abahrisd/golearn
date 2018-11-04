package main

import "fmt"

func main() {

	testArr := []int{1, 8, 2, 9, 3, 4, 5, 6}

	sorted := qsort(testArr)
	fmt.Println("sorted: ", sorted)

}

func qsort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	leftPart := []int{}
	rightPart := []int{}
	pIndex := len(arr) / 2
	pivot := arr[pIndex]

	for key, v := range arr {
		if pIndex == key {
			continue
		} else if v < pivot {
			leftPart = append(leftPart, v)
		} else {
			rightPart = append(rightPart, v)
		}
	}

	return append(append(qsort(leftPart), pivot), qsort(rightPart)...)
}
