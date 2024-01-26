package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	foundItem := binarySearch(nums, 7)
	fItem := binarySearchRecursive(nums, 7)

	fmt.Println(foundItem, fItem)
}

func binarySearch(nums []int, item int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearchRecursive(nums []int, item int) int {
	if len(nums) == 0 {
		return -1
	}

	return binarySearchRecursiveInternal(nums, item, 0, len(nums)-1)
}

func binarySearchRecursiveInternal(nums []int, item, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	midValue := nums[mid]

	if midValue == item {
		return mid
	}

	if midValue < item {
		return binarySearchRecursiveInternal(nums, item, mid+1, high)
	} else {
		return binarySearchRecursiveInternal(nums, item, low, mid-1)
	}
}
