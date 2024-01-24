package main

import "fmt"

func main() {
	userNums := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sum(userNums))
	fmt.Println(sumRecursive(userNums))
}

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}

	return total
}

func sumRecursive(nums []int) (total int) {
	if len(nums) == 0 {
		return 0
	}

	first := nums[0]
	num := nums[1:]
	total = first + sumRecursive(num)

	return total
}
