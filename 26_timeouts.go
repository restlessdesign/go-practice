package main

import (
	"fmt"
	"time"
)

func makeRequest(request_channel chan string, response string, duration int) {
	time.Sleep(time.Second * time.Duration(duration))
	request_channel <- response
}

func handleResponse(request chan string) {
	select {
	case response := <-request:
		fmt.Println(response)
	case <-time.After(time.Second * 2):
		fmt.Println("Request timed out")
	}
}

func main() {
	request_1 := make(chan string)
	go makeRequest(request_1, "loldata", 1)
	handleResponse(request_1)

	request_2 := make(chan string)
	go makeRequest(request_2, "loldata", 3)
	handleResponse(request_2)
}
