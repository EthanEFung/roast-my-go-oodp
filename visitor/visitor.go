package main

type visitor interface {
	visitPikachu(*pikachu)
	visitCharmander(*charmander)
}