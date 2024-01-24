package main

import "fmt"

func main() {
	userNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(count(userNums))
	fmt.Println(countRecursive(userNums))
}

func count(nums []int) (total int) {
	total = len(nums)

	return total
}

func countRecursive(nums []int) (total int) {
	if len(nums) == 0 {
		return 0
	}

	total = 1 + countRecursive(nums[1:])
	return total
}
