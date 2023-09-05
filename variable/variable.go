package main

import "fmt"

func main() {
	//proper declaration for common language variable, datatype and declaration
	var a int = 10
	//we can declare this becuase it will take the datatypes and variable("on the go feature")
	b := 90
	//we can delare this also( datatypes is not required, it will take go itself)
	var c = 20
	fmt.Printf("%d, %d, %d", a, b, c) //we can use this both prints, formatter and printer
	fmt.Println(a, b, c)
}
