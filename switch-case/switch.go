package main

import "fmt"

func main() {
	var c byte
	fmt.Println("Do you want to buy? (y/n)")
	fmt.Scanf("%c", &c)

	switch c {
	case 'y':
		fmt.Println("Thank you")
		//fallthrough  --> in C we are using break statement for every switch case but in go break is not required
		// if we using fallthrough the statement will not break, it will execude for both statement(which means Thank you , no problem both will print)
	case 'n':
		fmt.Println("No problem")
	default:
		fmt.Println("Invalid")

	}
}
