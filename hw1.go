package main

import (
	"fmt"
)

func main() {
	i:=1
	for i=1; i<=100; i++ {
		fmt.Println("Convert number ",i," to roman is ",convToRoman(i))
	}
	
	
}

func convToRoman(x int) string {
	roman := ""

	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	
	i := 0
	for x > 0 {
		k := x / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			x -= values[i]
		}
		i++
	}
	return roman
}