package main

import (
	"fmt"
	"math"
)

func main() {
	userNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(maximum(userNums))
	fmt.Println(maximumRecursive(userNums))
}

func maximum(nums []int) (maxNum int) {
	maxNum = math.MinInt

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func maximumRecursive(nums []int) (maxNum int) {
	if len(nums) == 0 {
		maxNum = math.MinInt
		return maxNum
	}

	num := nums[0]
	maxNum = maximumRecursive(nums[1:])
	if num >= maxNum {
		maxNum = num
		return maxNum
	}

	return maxNum
}
