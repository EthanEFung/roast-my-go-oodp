package main

import "fmt"
func main() {
	fileA, fileB, fileC := &file{"file A"}, &file{"file B"}, &file{"file C"}
	folderA := &folder{"folder A", []printablePrototype{fileA, fileB}}
	folderB := &folder{"folder B", []printablePrototype{folderA, fileC}}

	fmt.Println(folderB.print(4))
	
	folderC := folderB.clone()

	fmt.Println(folderC.print(4))
}