package main

import "fmt"

const (
	RedDressType  = "RedDress"
	BlueDressType = "BlueDress"
)

var dressFactoryInstance = &dressFactory{
	dressMap: make(map[string]dress),
}

/*
	dressFactory is an object that stores references to the different
	`dress` objects that our game supports.
*/
type dressFactory struct {
	dressMap map[string]dress
}

/*
	getDressByType will return a dress if not previously instantiated, but
	supported by our application.
*/
func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	t, exists := d.dressMap[dressType]
	if exists {
		return t, nil
	}
	if dressType == RedDressType {
		d.dressMap[dressType] = newRedDress()
		return d.dressMap[dressType], nil
	}
	if dressType == BlueDressType {
		d.dressMap[dressType] = newBlueDress()
		return d.dressMap[dressType], nil
	}
	return t, fmt.Errorf("could not find dress type %v", dressType)
}

func getDressFactoryInstance() *dressFactory {
	return dressFactoryInstance
}
