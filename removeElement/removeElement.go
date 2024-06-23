package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
