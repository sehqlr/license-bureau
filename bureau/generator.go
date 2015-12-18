package bureau

import (
	"bufio"
	"log"
	"os"
	"text/Template"
)

// Template for 'Component' section of BUREAU.md
const component_template = `
## Components
#### {{.Name}}
Project's page is [here]({{.URL}})
It is licensed under the terms of {{.License.Name}}
The project's copy of the license is in the [{{.License.Local}}]({{.URL}}/{{.License.Local}}) file.
The full text of the license can be found [below](#{{.License.Name}})
`

// Template for 'License' section of BUREAU.md
const license_template = `
## Licenses
#### {{.Name}}
This is the full text of the {{.Name}}. It is used for these projects:
{{range projects}}
* [{{.component}}](#{{.component}})
{{end}}
**FULL TEXT FOLLOWS**
{{fetch .URL}}
`

func GenerateMarkdown(config map[string]Component, outfile string) *bufio.Writer {
	out, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	markdown := bufio.NewWriter(out)
	c_tmpl := template.Must(template.New("component_template").Parse(component_template))
	//l_tmpl := template.Must(template.New("license_template").Parse(license_template))

	for name, component := range config {
		component.Name = name
		err = c_tmpl.Execute(markdown, component)
		if err != nil {
			log.Fatal(err)
		}
	}

	return markdown
}
