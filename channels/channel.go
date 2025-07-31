package main

import (
	"fmt"
)

// func manageChannel(numC chan int) {

// 	for i := range numC {
// 		fmt.Println("Processing numbers", i)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(resChannel chan int, num1, num2 int) {
// 	numResult := num1 + num2

// 	resChannel <- numResult
// }

func mailSender(email <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for i := range email {
		fmt.Println("Mail has been sent", i)
	}
}

func main() {

	emailChan := make(chan string, 20)
	done := make(chan bool)

	go mailSender(emailChan, done)

	for i := 0; i < 20; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Mail sending done")

	//it's important to close buffer channel
	close(emailChan)

	<-done

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println(res)

	// numChannel := make(chan int)

	// nums := [5]int{1, 23, 4, 5, 6}

	// go manageChannel(numChannel)

	// for i := 0; i < len(nums); i++ {
	// 	numChannel <- nums[i]
	// }
}
