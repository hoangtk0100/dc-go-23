package util

import (
	"encoding/json"
	"log"
	"os"
)

func OpenFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func CloseFile(file *os.File) error {
	err := file.Close()
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(file *os.File, data any) error {
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
