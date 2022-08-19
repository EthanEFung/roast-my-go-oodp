package main

import (
	"fmt"
	"math"
)

/*
	Burger in this example would be an existing interface that decorators will adhere to
*/
type Burger interface {
	Ingredients() []string
	Price() float64
}

type BaseBurger struct {
	Burger
}

func (h BaseBurger) Ingredients() []string {
	return []string{}
}
func (h BaseBurger) Price() float64 {
	return 10.27
}

/*
	WithBun is a concrete decorator that takes in a base struct that adheres to the Burger interface
*/
type WithBun struct {
	Burger
	Type           string
	AdditionalCost float64
}

func (d WithBun) Ingredients() []string {
	return append(d.Burger.Ingredients(), d.Type+" Bun")
}

/*
	WithCheese is a concrete decorator that takes in a base struct that adheres to the Burger interface
*/
type WithCheese struct {
	Burger
	Type           string
	AdditionalCost float64
}

func (d WithCheese) Ingredients() []string {
	base := d.Burger.Ingredients()
	return append(base, d.Type+" Cheese")
}
func (d WithCheese) Price() float64 {
	return math.Floor((d.Burger.Price()+1.75)*100) / 100
}

/*
	WithPattie is a concrete decorator that takes in a base struct that adheres to the Burger interface
*/
type WithPattie struct {
	Burger
	Type           string
	AdditionalCost float64
}

func (d WithPattie) Ingredients() []string {
	return append(d.Burger.Ingredients(), d.Type+" Pattie")
}
func (d WithPattie) Price() float64 {
	return math.Floor((d.Burger.Price()+d.AdditionalCost)*100) / 100
}

func main() {
	base := &BaseBurger{}

	/*
		because all of these decorated objects adhere to the same interface as the base burger,
		this allows us the ability to use the same methods while modifying the processes that
		take place within those methods.
	*/
	hamburger := &WithPattie{&WithBun{base, "Brioche", 0}, "Beef", 0}
	cheeseBurger := &WithCheese{hamburger, "American", 0.75}
	veggieBurger := &WithPattie{&WithBun{base, "Brioche", 0}, "Impossible", 2.10}
	proteinStyle := &WithCheese{&WithPattie{&WithBun{base, "Lettuce", 0}, "Beef", 0}, "American", 0}

	fmt.Println("hamburger", hamburger.Ingredients(), hamburger.Price())
	fmt.Println("cheese burger", cheeseBurger.Ingredients(), cheeseBurger.Price())
	fmt.Println("veggie burger", veggieBurger.Ingredients(), veggieBurger.Price())
	fmt.Println("protein style", proteinStyle.Ingredients(), proteinStyle.Price())
}
