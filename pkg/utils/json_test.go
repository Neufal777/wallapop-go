package utils

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSaveToJSON(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_output*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Define the test data
	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	// Call the function being tested
	SaveToJSON(data, tempFile.Name())

	// Read the contents of the saved JSON file
	savedData, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the saved JSON data
	var parsedData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err = json.Unmarshal(savedData, &parsedData)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the data was saved correctly
	expectedName := "John Doe"
	expectedAge := 30
	if parsedData.Name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, parsedData.Name)
	}
	if parsedData.Age != expectedAge {
		t.Errorf("Expected age to be %d, but got %d", expectedAge, parsedData.Age)
	}
}
