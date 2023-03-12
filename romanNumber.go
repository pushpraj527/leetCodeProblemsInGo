/*
13. Roman to Integer: https://leetcode.com/problems/roman-to-integer/
*/

package main

import (
	"fmt"
)

func romanToInt(s string) int {
	rom := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int
	fmt.Println(len(s))
	for i, _ := range s {
		if (i+1 < len(s)) && (rom[string(s[i])] < rom[string(s[i+1])]) {
			sum = sum - rom[string(s[i])]
		} else {
			sum = sum + rom[string(s[i])]
		}
	}
	return sum
}

func main() {
	str1 := "XVX"
	result := romanToInt(str1)
	fmt.Println(result)
}
