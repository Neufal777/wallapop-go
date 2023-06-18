package tools

import (
	"encoding/json"
	"time"
)

func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func TimeFormatter(timestamp int64) string {
	seconds := timestamp / 1000             // Convert milliseconds to seconds
	nanoseconds := (timestamp % 1000) * 1e6 // Convert milliseconds to nanoseconds

	// Create a time object using the Unix function
	t := time.Unix(seconds, nanoseconds)

	// Format the time object to the desired date format
	date := t.Format("2006-01-02 15:04:05 MST")

	return date
}

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
