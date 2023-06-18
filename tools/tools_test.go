package tools

import (
	"testing"
	"time"
)

func TestTimeFormatter(t *testing.T) {
	// Define the test input
	timestamp := int64(1624010403000) // June 18, 2021 12:00:03 CEST

	// Define the time zone for comparison
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Fatal(err)
	}

	// Call the function being tested
	result := TimeFormatter(timestamp)

	// Parse the expected time in the time zone
	expectedTime, err := time.ParseInLocation("2006-01-02 15:04:05 MST", "2021-06-18 12:00:03 CEST", loc)
	if err != nil {
		t.Fatal(err)
	}

	expected := expectedTime.Format("2006-01-02 15:04:05 MST")

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

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
