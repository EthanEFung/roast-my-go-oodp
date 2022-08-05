package main

func main() {
	app := &app{}
	app.setJSONConfig([]byte("json"))

	service := &yamlService{}
	
	a := &adapter{service, "yaml"}
	a.convert()
	a.setJSONConfig([]byte{})
}