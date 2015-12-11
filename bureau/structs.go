package bureau

import (
	"io/ioutil"
	"log"

	"net/http"
)

type SoftwareLicense struct {
	Name, Local, Web string
}

type Component struct {
	License   SoftwareLicense
	Name, URL string
}

func (license SoftwareLicense) GetFullText() []byte {
	resp, err := http.Get(license.Web)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
