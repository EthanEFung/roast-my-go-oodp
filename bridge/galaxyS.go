package main

import "fmt"

/*
galaxyS is a concrete "abstraction" of the `mobileDevice` interface
*/
type galaxyS struct {
	mapApp mapApp
}

func (d *galaxyS) installMap(app mapApp) {
	d.mapApp = app
}

func (d *galaxyS) searchAddress(addr string) {
	if d.mapApp == nil {
		fmt.Println("No map application installed on device")
		return
	}
	d.mapApp.autocomplete(addr)
}
