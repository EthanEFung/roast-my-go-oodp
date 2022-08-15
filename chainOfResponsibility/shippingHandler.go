package main

import "fmt"

/*
  shippingHandler is a concrete handler of the chain of responsibility pattern
*/
type shippingHandler struct {
	next handler
}

func (h *shippingHandler) setNext(next handler) {
	h.next = next
}
func (h *shippingHandler) execute(order *order) {
	// here we would do something like calculate the distance
	// and fees of our carrier
	order.shippingCost = 5.99
	fmt.Println("shipping cost added to order")
	h.next.execute(order)
}
