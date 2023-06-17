package tools

import (
	"encoding/json"
	"log"
	"os"
)

func SaveToJSON(result interface{}, outputFile string) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
