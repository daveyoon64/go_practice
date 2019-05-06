package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	htype := resp.Header.Get("Content-Type")
	if htype == "" {
		return "", fmt.Errorf("Response error (%s)", url)
	}
	return htype, nil
}

func main() {
	htype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(htype)
	}
}
