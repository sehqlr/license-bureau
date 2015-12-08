package main

import (
	//	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type SoftwareLicense struct {
	Local string
	Name  string
	Web   string
}

type Component struct {
	License SoftwareLicense
	Name    string
	URL     string
}

func main() {
	filepath := os.Args[1]

	data, io_err := ioutil.ReadFile(filepath)
	if io_err != nil {
		log.Fatal(io_err)
	} else {
		fmt.Printf("Loaded: %s\n", filepath)
	}

	config := map[string]Component{}

	yaml_err := yaml.Unmarshal(data, &config)
	if yaml_err != nil {
		log.Fatal(yaml_err)
	}

	//var colophon, licenseText bytes.Buffer

	for _, component := range config {
		msg := "Component: %s\nURL: %s\nLicense:\n\tName: %s\n\tWeb URL: %s\n\tLocal Copy: %s\n"
		fmt.Printf(msg, component.Name, component.URL, component.License.Name, component.License.Web, component.License.Local)
	}

}
