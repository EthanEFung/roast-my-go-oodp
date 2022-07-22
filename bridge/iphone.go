package main

import "fmt"

type iphone struct {
  mapApp mapApp
}

func(d *iphone) searchAddress(addr string) {
  if d.mapApp == nil {
    fmt.Println("iphone does not have a map app installed")
    return
  }
  d.mapApp.autocomplete(addr)
}

func(d *iphone) installMap(app mapApp) {
  d.mapApp = app
}



