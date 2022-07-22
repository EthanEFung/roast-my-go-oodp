package main

import "fmt"

type gmaps struct {}

func (m *gmaps) autocomplete(substr string) {
  fmt.Printf("Google searching for %v\n", substr)
}

