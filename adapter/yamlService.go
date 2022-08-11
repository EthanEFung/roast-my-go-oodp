package main

/*
yamlService is a concrete service
*/
type yamlService struct {}

func (service yamlService) exportYAML() string {
	return "yaml"
}