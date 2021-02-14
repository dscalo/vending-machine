package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Snack struct {
	Name string 	`json:"name"`
	Price float64 	`json:"price"`
	Qty int			`json:"qty"`
	Desc string		`json:"desc"`
}

type Snacks struct {
	Snacks []Snack `json:"snacks"`
}

func (s *Snacks) longestName() int {
	length := 0

	for _, snack := range s.Snacks {
		nameLength := len(snack.Name)
		if  nameLength > length {
			length = nameLength
		}
	}

	return length
}


func GetSnacks(path string) (*Snacks, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var snacks Snacks

 	err = json.Unmarshal(byteVal, &snacks)

	if err != nil {
		return nil, err
	}

	return &snacks, nil
}
