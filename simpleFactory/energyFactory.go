package main

/*
	energyFactory is a concrete cardFactory that produces the concrete product `energyCard`
*/
type energyFactory struct {}
func (f energyFactory) create(name string) card {
	return &energyCard{name}
}