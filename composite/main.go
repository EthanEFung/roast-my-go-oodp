package main

func main() {

	file1 := file{"File A", "The Quick Brown Fox Jumps Over The Lazy Dog"}
	file2 := file{"File B", "Peter Piper Picked A Pack Of Pickled Peppers"}
	file3 := file{"File C", "What Now Brown Cow"}
	file4 := file{"File D", "Never Gonna Give You Up, Never Gonna Let You Down"}
	file5 := file{"File E", "Take Me Out To The Ball Game"}
	file6 := file{"File F", "Snoop Doggy Dawg"}
	file7 := file{"File G", "Cowards Never Prosper"}

	folder1 := folder{"Folder 1", []composite{file1}}
	folder2 := folder{"Folder 2", []composite{file2, file3}}
	folder3 := folder{"Folder 3", []composite{file4, file5, file6}}
	folder4 := folder{"Folder 4", []composite{folder1, folder2, file7}}
	root := folder{"Root", []composite{folder3, folder4}}
	root.search("Cow")
	root.search("Dog")
}
