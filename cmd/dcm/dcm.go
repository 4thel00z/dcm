package main

import (
	"encoding/json"
	"fmt"
	"github.com/4thel00z/dcm/pkg/libdcm/surveys"
	"log"
)

func main() {
	domain, err := surveys.Domain()
	if err != nil {
		log.Fatal(err.Error())
	}

	marshalled, err := json.Marshal(domain)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(string(marshalled))
}
