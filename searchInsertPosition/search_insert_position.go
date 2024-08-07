package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if target == num {
			return i
		} else if num > target {
			return i
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
