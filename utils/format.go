package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func CheckIfCanSplit(input string, separator string) ([]string, bool) {
	if strings.Contains(input, separator) {
		values := strings.Split(input, separator)
		return values, true
	}

	return nil, false
}

// PrettyPrintStruct Prints structures with indents
func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// AnyTypeToString Converts any type to <string> type
func AnyTypeToString(input interface{}) string {
	return fmt.Sprintf("%v", input)
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
