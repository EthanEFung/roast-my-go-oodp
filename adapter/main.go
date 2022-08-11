package main

func main() {
	app := &app{}
	app.exportJSON([]byte("json"))

	yamlService := &yamlService{}
	
	a := &adapter{service: yamlService}
	a.convert()
	a.exportJSON([]byte{})
}