package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", "http://localhost:3000/3weeks-go/testGet", nil)
	if err != nil {
		panic(err)
	}
	body, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
