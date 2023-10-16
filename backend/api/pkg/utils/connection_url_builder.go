package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	switch n {
	case "fiber":
		// URL for fiber connection
		url = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	default:
		return "", fmt.Errorf("connection name %v is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
