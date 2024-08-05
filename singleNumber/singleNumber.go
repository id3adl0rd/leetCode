package singleNumber

import "fmt"

func singleNumber(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
		fmt.Println(result)
	}

	return result
}
