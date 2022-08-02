package main

import "fmt"
func main() {
	fileA, fileB, fileC := &file{"file A"}, &file{"file B"}, &file{"file C"}
	folderA := &folder{"folder A", []printablePrototype{fileA, fileB}}
	folderB := &folder{"folder B", []printablePrototype{folderA, fileC}}
	folderC := &folder{"folder C", []printablePrototype{}}
	folderD := &folder{"folder D", []printablePrototype{folderB, folderC}}

	fmt.Println(folderD.print(0))
	
	folderDClone := folderD.clone()

	fmt.Println(folderDClone.print(0))
}