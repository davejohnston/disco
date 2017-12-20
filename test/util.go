package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func LoadFixtureFromFile(filename string, fixture interface{}) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &fixture)
}
