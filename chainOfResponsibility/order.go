package main

type item struct {
	name     string
	quantity int
	price    float64
}

type payment struct {
	cardHolder   string
	cardNumber   string
	securityCode string
	expiration   string
	network      string
}

type order struct {
	address      string
	payment      payment
	basket       []item
	shippingCost float64
}
