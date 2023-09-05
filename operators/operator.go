package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("%s", FizzBuzz(num))
}

// N is not calling on main but we calling this (n int) on main method. so it will replace our num into n and give the result
func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
