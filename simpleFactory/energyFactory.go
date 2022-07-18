package main

type energyFactory struct {}
func (f energyFactory) create(name string) card {
	return &energyCard{name}
}