package main

import "fmt"

func main() {
	card := payment{
		cardHolder:   "John Doe",
		cardNumber:   "5200 0000 0000 0000",
		securityCode: "123",
		expiration:   "2032-03-19T07:22Z",
		network:      "Master Card",
	}

	badPayment := payment{
		cardHolder:   "John Doe",
		cardNumber:   "4200 0000 0000 0000",
		securityCode: "123",
		expiration:   "2032-03-19T07:22Z",
		network:      "Master Card",
	}

	basket := []item{
		{"HealthAde Kombucha - Tropical Punch", 6, 4.99},
		{"Red Bull - Original", 6, 2.25},
	}

	goodOrder := &order{
		address: "1000 Vin Scully Ave, Los Angeles, CA 90012",
		payment: card,
		basket:  basket,
	}

	badOrder := &order{
		address: "1000 Vin Scully Ave, Los Angeles, CA 90012",
		payment: badPayment,
		basket:  basket,
	}

	/** CoC **/
	transactor := &transactionHandler{}

	shippingCostAdder := &shippingHandler{}
	shippingCostAdder.setNext(transactor)

	paymentValidator := &paymentValidationHandler{}
	paymentValidator.setNext(shippingCostAdder)

	fmt.Printf("...execute bad attempt...\n")
	paymentValidator.execute(badOrder)

	fmt.Printf("\n...execute good attempt...\n")
	paymentValidator.execute(goodOrder)
}
