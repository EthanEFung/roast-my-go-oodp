package main

import "fmt"

type appleMaps struct {}

func (app *appleMaps) autocomplete(substr string) {
  fmt.Printf("Apple Maps searching for %v\n", substr)
}

