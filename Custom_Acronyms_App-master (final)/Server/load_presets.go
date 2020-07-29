package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Acronym represents an acronym and its definition
type Acronym struct {
	Acronym string `json:"acronym"`
	Def     string `json:"def"`
}

func loadPreset(path string) []Acronym {
	data, err := ioutil.ReadFile("./" + path)
	if err != nil {
		fmt.Println(err)
	}

	var acronym []Acronym

	err = json.Unmarshal(data, &acronym)
	if err != nil {
		fmt.Println(err)
	}

	return acronym
}
