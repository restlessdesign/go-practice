package main

import (
	"encoding/json"
	"fmt"
)

type UserResponse struct {
	Name      string
	Age       int
	LovesDogs bool
}

type UserResponseFormatted struct {
	// IMPORTANT
	// There must not be a space between the colon and double-quoted key name!
	Name      string `json:"name"`
	Age       int    `json:"age"`
	LovesDogs bool   `json:"loves_dogs"`
}

func main() {
	rawResponse := &UserResponse{
		Name:      "Kevin",
		Age:       31,
		LovesDogs: true,
	}

	rawJSON, err := json.Marshal(rawResponse)
	if err == nil {
		fmt.Println(string(rawJSON))
	}

	formattedResponse := &UserResponseFormatted{
		Name:      "Kevin",
		Age:       31,
		LovesDogs: true,
	}

	formattedJSON, err := json.Marshal(formattedResponse)
	if err == nil {
		fmt.Println(string(formattedJSON))
	}

	jsonStr := `{"name": "Melissa", "age": 25, "loves_dogs": true}`
	res := &UserResponseFormatted{}
	fmt.Println(jsonStr)
	json.Unmarshal([]byte(jsonStr), res)
	fmt.Printf("%+v", res)
}
