package main

type yamlService struct {}

func (service yamlService) exportYAML() string {
	return "yaml"
}