//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080",
		bytes.NewBuffer([]byte("hello world")))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}
