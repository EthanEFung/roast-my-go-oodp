package main

func main() {
	app := &app{}
	app.exportJSON([]byte("json"))

	yamlService := &yamlService{}
	
	a := &adapter{service: yamlService}
<<<<<<< HEAD
=======
	a.convert()
>>>>>>> b9317a09d281e91a5c996d0f38f35734ed52e827
	a.exportJSON([]byte{})
}