package main

import "fmt"

// declare const variable which later can not be modify

const val = 333

var num int32 = 4444

// this syntax does not work outside main
// credit :=  444

func main() {

	// variable declaration
	name := "hit shiroya"

	// normal print using fmt
	fmt.Println("Hit Shiroya")
	fmt.Println(name)

	fmt.Println(num)

	// print const var
	fmt.Println(val)
}
