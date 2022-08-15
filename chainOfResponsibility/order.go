/*
  the structs in this file do not pertain to the pattern but
  are used to demonstrate the pattern in the example
*/
package main

/*
  item represents an end users product to purchase
*/
type item struct {
	name     string
	quantity int
	price    float64
}

/*
  payment represents the end users method of paying
*/
type payment struct {
	cardHolder   string
	cardNumber   string
	securityCode string
	expiration   string
	network      string
}

/*
  order represents the vital information of the end user needed to
  complete a transaction
*/
type order struct {
	address      string
	payment      payment
	basket       []item
	shippingCost float64
}
