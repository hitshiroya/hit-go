package main

import "fmt"

// declare const variable which later can not be modify

const val = 333

var num int32 = 4444

// this syntax does not work outside main
// credit :=  444

func main() {

	// variable declaration
	//name := "hit shiroya"

	// normal print using fmt
	// fmt.Println("Hit Shiroya")
	// fmt.Println(name)

	// fmt.Println(num)

	// // print const var
	// fmt.Println(val)

	// looping only for loop is there

	// while loop using for loop

	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// if age := 22; age >= 18 {
	// 	fmt.Println("Elible to vote")
	// } else if age <= 18 {
	// 	fmt.Println("Not eligible")
	// }

	var names [4]string
	names[0] = "hit"
	names[1] = "shiroya"

	fmt.Println(names[1])

	balance := [5]int{400, 300, 200, 20, 2}

	var totalBalance int

	for i := range len(balance) {
		totalBalance += balance[i]
	}

	fmt.Println(totalBalance)

	fmt.Println(balance)

	matrixMal := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i,j:=range 3,3{
		
	}
	fmt.Println(matrixMal)
}
