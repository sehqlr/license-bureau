package bureau

import (
	"log"

	"gopkg.in/yaml.v2"
)

func ParseConfig(data []byte) map[string]Component {
	config := map[string]Component{}

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
