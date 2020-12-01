package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/4thel00z/dcm/pkg/libdcm/surveys"
	"io/ioutil"
	"log"
	"os"
)

var (
	output = flag.String("output", "", "path to output file, defaults to random path")
)

func main() {
	flag.Parse()
	var (
		out *os.File
		err error
	)

	if *output == "" {
		out, err = ioutil.TempFile("", "dcm")
		cry(err)
		defer func() {
			_ = out.Close()
		}()
		*output = out.Name() + ".json"
	} else {
		out, err = os.Create(*output)
		cry(err)
		defer func() {
			_ = out.Close()
		}()
	}

	domain, err := surveys.Domain()
	cry(err)

	marshalled, err := json.Marshal(domain)
	cry(err)

	_, err = out.WriteString(string(marshalled))
	cry(err)

	fmt.Print("Written file to: ", *output)
}

func cry(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
