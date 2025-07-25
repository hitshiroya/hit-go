package main

import (
	"fmt"
	"time"
)

type order struct {
	id     int
	amount float32
	status string
	time   time.Time
}

func (o order) checkPayment(id int) bool {
	return true
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	order1 := order{
		id:     1,
		amount: 45.96,
		status: "Pending",
		time:   time.Now(),
	}

	paymentStatus := order1.checkPayment(order1.id)

	if paymentStatus == true {
		order1.changeStatus("Shipped")
	}

	fmt.Println(order1)
}
