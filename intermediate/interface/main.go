package main

import "fmt"

type Payment interface {
	Pay(amount int)
}

type UPI struct{}
type Card struct{}
type Cash struct{}

func (u UPI) Pay(amount int) {
	fmt.Println("Paid via UPI", amount)
}

func (c Card) Pay(amount int) {
	fmt.Println("Paid via Card", amount)
}

func (ch Cash) Pay(amount int) {
	fmt.Println("Paid via Cash", amount)
}

func processPayment(p Payment, amount int) {
	p.Pay(amount)
}

func main() {

	fmt.Println("Interface....")
	upi := UPI{}
	card := Card{}
	cash := Cash{}

	processPayment(upi, 1000)
	processPayment(card, 500)
	processPayment(cash, 100)

}
