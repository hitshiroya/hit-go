package main

import "fmt"

// declare const variable which later can not be modify

// const val = 333

// var num int32 = 4444

// this syntax does not work outside main
// credit :=  444

// func add(a int, b int) int {
// 	return a + b
// }

// func mul(a, b, c int) int {
// 	return a * b * c
// }
// func information() (string, string, string) {
// 	return "Hit", "Shiroya", "Than"
// }

// func avg(nums ...int) int {
// 	total := 0

// 	for _, num := range nums {
// 		total = total + num
// 	}

// 	return total / len(nums)
// }

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

	// var names [4]string
	// names[0] = "hit"
	// names[1] = "shiroya"

	// fmt.Println(names[1])

	// balance := [5]int{400, 300, 200, 20, 2}

	// var totalBalance int

	// for i := range len(balance) {
	// 	totalBalance += balance[i]
	// }

	// fmt.Println(totalBalance)

	// fmt.Println(balance)

	// slices dynamic size change on run time not fix

	// using simple

	// var arrSlice []int

	// arrSlice = append(arrSlice, 2)
	// arrSlice = append(arrSlice, 4)
	// fmt.Print(cap(arrSlice))

	// var value = make([]int, 2, 10)
	// value = append(value, 4)
	// value = append(value, 2)
	// value = append(value, 34)

	// fmt.Println(value)

	// fmt.Println(cap(value))

	// m := make(map[string]int)

	// m["hit"] = 333
	// m["neel"] = 222

	// fmt.Println(m["hit"])

	// two syntax to create map

	// make syntax

	// m := make(map[string]int)

	// m["hit"] = 1
	// m["x"] = 2

	// v, ok := m["p"]

	// if ok {
	// 	fmt.Print("hello")
	// 	fmt.Print(v)
	// } else {
	// 	fmt.Print("not done")
	// }

	// m := map[string]int{
	// 	"hit": 1,
	// 	"x":   333,
	// 	"y":   444,
	// }

	// v, ok := m["t"]

	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("Wrong request")
	// }

	// result := add(3, 4)

	// resultMul := mul(4, 5, 6)
	// fmt.Println(result)
	// fmt.Println(resultMul)

	// firstName, lastName, _ := information()

	// fmt.Println(firstName, lastName)

	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// r1 := avg(nums...)

	// fmt.Println(r1)

}
