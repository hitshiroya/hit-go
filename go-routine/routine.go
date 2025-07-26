package main

import (
	"fmt"
	"time"
)

func printNum(i int) {
	fmt.Println(i)
}

func main() {

	for i := 0; i <= 10; i++ {
		go printNum(i)
	}

	time.Sleep(time.Millisecond * 3)
}
