package utils

import (
	"testing"
)

func TestPrettyPrintStruct(t *testing.T) {
	type TestStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	testStruct := TestStruct{
		ID:   123,
		Name: "Test",
	}

	expectedOutput := "{\n\t\"id\": 123,\n\t\"name\": \"Test\"\n}"

	result := PrettyPrintStruct(testStruct)
	if result != expectedOutput {
		t.Errorf("Unexpected output. Expected:\n%s\nGot:\n%s\n", expectedOutput, result)
	}
}

func TestContains(t *testing.T) {
	slice := []string{"rocket", "propulsion", "spacex"}

	t.Run("Existing element", func(t *testing.T) {
		result := Contains(slice, "rocket")
		if !result {
			t.Errorf("Expected 'true', but got 'false' for existing element")
		}
	})

	t.Run("Non-existing element", func(t *testing.T) {
		result := Contains(slice, "grape")
		if result {
			t.Errorf("Expected 'false', but got 'true' for non-existing element")
		}
	})
}
