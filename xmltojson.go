package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
	kz "github.com/qntfy/kazaam"
)

const (
	benchmarkJSON = `{"topKeyA": {"arrayKey": [{"foo": 0}, {"foo": 1}, {"foo": 1}, {"foo": 2}], "notArrayKey": "Sun Jul 23 08:15:27 +0000 2017", "deepArrayKey": [{"key0":["val0", "val1"]}]}, "topKeyB":{"nextKeyB": "valueB"}}`
)

func main() {
	lg := log.New(os.Stdout, "xml2json ", log.LstdFlags)
	// fmt.Println("Starting the main application!")

	file, err := os.Open("src.xml")
	if err != nil {
		lg.Fatal("failed to read the source file")
		os.Exit(1)
	}

	json, err := xj.Convert(file)
	if err != nil {
		lg.Fatal("failed to convert xml file!")
		os.Exit(1)
	}
	// fmt.Println(json.String())

	specb, err := ioutil.ReadFile("spec.json")
	spec := string(specb)

	xform, _ := kz.NewKazaam(spec)
	kout, _ := xform.TransformJSONStringToString(json.String())

	fmt.Println(kout)
}
