package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	slc := []int{-100, -10, 0, 5, -150, 7, 10, 25}

	fmt.Println(selectionSort(slc))
}

func findMinIdx(slc []int) (minimalIdx int, err error) {
	if len(slc) == 0 {
		return 0, errors.New("empty slice")
	}

	minimal := math.MaxInt
	for i, v := range slc {
		if v < minimal {
			minimalIdx, minimal = i, v
		}
	}

	return minimalIdx, nil
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func selectionSort(slc []int) (sortSlc []int, err error) {
	for len(slc) > 0 {

		idx, err := findMinIdx(slc)
		if err != nil {
			return slc, err
		}

		sortSlc = append(sortSlc, slc[idx])
		slc = remove(slc, idx)
	}

	return sortSlc, nil
}
