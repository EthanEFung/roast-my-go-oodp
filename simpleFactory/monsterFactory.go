package main
type monsterFactory struct {

}
func (f monsterFactory) create(name string) card {
	return &monsterCard{name}
}
