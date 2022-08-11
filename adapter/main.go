package main

func main() {
	app := &app{}
	app.exportJSON([]byte("json"))

	yamlService := &yamlService{}
	
	a := &adapter{service: yamlService}
	a.exportJSON([]byte{})
}