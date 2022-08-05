package main

type client interface {
	setJSONConfig([]byte)
}