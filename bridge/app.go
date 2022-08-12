package main

/*
mapApp is the "implementation" interface of the bridge pattern
*/
type mapApp interface {
	autocomplete(substr string)
}
