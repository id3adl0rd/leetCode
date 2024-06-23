package main

import "fmt"

func romanToInt(s string) int {
	sum := 0

	rome := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, i2 := range s {
		if _, ok := rome[i2]; ok {
			sum += rome[i2]

			if i == 0 {
				continue
			}

			if rome[rune(s[i-1])] < rome[rune(s[i])] {
				sum -= 2 * rome[rune(s[i-1])]
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
