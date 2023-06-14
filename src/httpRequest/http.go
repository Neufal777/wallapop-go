package httpRequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAPIResponse(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make the HTTP request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %s", err)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %s", err)
	}

	return nil
}
