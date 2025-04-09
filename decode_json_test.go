package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"Firstname":"Eko","Middlename":"Kurniawan","Lastname":"Khannedy"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n", customer)
	fmt.Println("\n", customer.Firstname)
	fmt.Println("\n", customer.Middlename)
	fmt.Println("\n", customer.Lastname)
}
