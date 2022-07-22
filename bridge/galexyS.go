package main

import "fmt"

type galexyS struct {
  mapApp mapApp
}

func (d *galexyS) installMap(app mapApp) {
  d.mapApp = app
}

func (d *galexyS) searchAddress(addr string) {
  if d.mapApp == nil {
    fmt.Println("No map application installed on device")
    return
  }
  d.mapApp.autocomplete(addr)
}
