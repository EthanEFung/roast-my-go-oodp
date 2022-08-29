package main

/*
	monsterFactory is a concrete cardFactory that produces the concrete product `monsterCard`
*/
type monsterFactory struct {
}
func (f monsterFactory) create(name string) card {
	return &monsterCard{name}
}
