package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname:  "Eko",
		Middlename: "Kurniawan",
		Lastname:   "Khannedy",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println("\n", string(bytes))
}
