package main

import "fmt"

type transactionHandler struct {
	next handler
}

func (h transactionHandler) setNext(next handler) {
	h.next = next
}
func (h transactionHandler) execute(order *order) {
	var total float64
	fmt.Printf("\n::: totals :::\n")

	for _, item := range order.basket {
		subtotal := (float64(item.quantity) * item.price)
		fmt.Printf("%v: $%v * %v = $%v\n", item.name, item.price, item.quantity, subtotal)
		total += subtotal
	}
	fmt.Printf("shipping cost $%v\n", order.shippingCost)
	total += order.shippingCost

	fmt.Printf("\ncharging: $%v\n", total)
	if h.next == nil {
		return 
	}
	h.next.execute(order)
}
