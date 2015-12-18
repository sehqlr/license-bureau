package main

import (
	"fmt"
	"io/ioutil"

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

	var licenseList []string

	colophon := bureau.GenerateMarkdown(config, "LICENSES.md")

	colophon.Write([]byte("## Licenses\n"))
	for _, component := range config {
		licenseList = append(licenseList, component.License.Local)

		colophon.Write([]byte("####" + component.License.Name + "\n"))
		colophon.Write([]byte("FULL TEXT FROM " + component.License.Web + ":\n"))
		colophon.Write(component.License.GetFullText())
		colophon.Write([]byte("\n"))
	}

	colophon.Flush()
}
