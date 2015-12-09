package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sehqlr/license-bureau/bureau"
)

func main() {
	filepath := "license.yml"

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("We didn't find %s in this directory.", filepath)
	} else {
		fmt.Printf("Loaded: %s\n", filepath)
	}

	config := bureau.ParseConfig(data)

	f, err := os.Create("BUREAU.md")
	if err != nil {
		log.Fatal(err)
	}

	colophon := bufio.NewWriter(f)
	var licenseList []string

	for name, component := range config {
		licenseList = append(licenseList, component.License.Local)

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
		colophon.Write([]byte("\n"))
	}

	colophon.Flush()
}
