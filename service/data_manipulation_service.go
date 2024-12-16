package service

import (
	"cadanapay/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetPersonsInstance() (model.Persons, error) {
	// read a list of person from /resources/persons.json
	file, err := os.Open("resources/persons.json")
	if err != nil {
		return model.Persons{}, fmt.Errorf("failed to open file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}(file)

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		return model.Persons{}, fmt.Errorf("failed to read file: %v", err)
	}

	// Unmarshal JSON into []Person
	var persons []model.Person
	if err := json.Unmarshal(data, &persons); err != nil {
		return model.Persons{}, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return model.Persons{
		Data: persons,
	}, nil
}
