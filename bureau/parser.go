package bureau

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func ParseConfig(data []byte) map[string]bureau.Component {
	config := map[string]bureau.Component{}

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
