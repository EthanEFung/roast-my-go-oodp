package main

/*
client is an interface which only has a method to set a config when passed json
*/
type client interface {
	exportJSON([]byte)
}