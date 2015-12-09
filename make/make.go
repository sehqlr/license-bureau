package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"license-bureau/bureau"
)

type SoftwareLicense struct {
	Local string
	Name  string
	Web   string
}

type Component struct {
	License SoftwareLicense
	URL     string
}

func main() {
	filepath := os.Args[1]

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Loaded: %s\n", filepath)
	}

	config := bureau.ParseConfig(data)

	var colophon bytes.Buffer
	var licenseList []string

	for name, component := range config {
		licenseList = append(licenseList, component.License.Local)

		colophon.Write([]byte("---\n"))
		colophon.Write([]byte("Component: " + name + "\n"))
		colophon.Write([]byte("Repo URL: " + component.URL + "\n"))
		colophon.Write([]byte("License: " + component.License.Name + "\n"))
		colophon.Write([]byte("Web URL: " + component.License.Web + "\n"))
		colophon.Write([]byte("Local Copy: " + component.License.Local + "\n"))

		resp, err := http.Get(component.License.Web)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		colophon.Write([]byte("FULL TEXT OF LICENSE:\n"))
		colophon.Write(body)
		colophon.Write([]byte("\n...\n"))
	}

	fmt.Print(colophon.String())
}
