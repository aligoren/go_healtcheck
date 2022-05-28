package providers

import (
	"fmt"
	"net/http"
	"os"
)

func CheckGoogle() ProviderResult {

	url := os.Getenv("GOOGLE_WEBSITE")

	result := ProviderResult{
		Name:    "GOOGLE_WEBSITE",
		Status:  true,
		Message: fmt.Sprintf("%s is accessible", url),
	}

	resp, err := http.Get(url)

	if err != nil {

		result.Status = false
		result.Error = fmt.Sprintf("Error while connecting site: %v", err)

		return result
	}

	if resp.StatusCode != http.StatusOK {

		result.Status = false
		result.Error = fmt.Sprintf("%s is inaccessible. Status code: %d", url, resp.StatusCode)

		return result
	}

	return result
}
