package main

import (
	"errors"
	"fmt"
)

func main() {
	slc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	foundItem, _ := binarySearch(slc, 7)

	fmt.Println(foundItem)
}

func binarySearch(slc []int, item int) (int, error) {
	if len(slc) == 0 {
		return 0, errors.New("empty slice")
	}

	low := 0
	high := len(slc) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := slc[mid]

		if guess == item {
			return mid, nil
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("unknown error")
}
