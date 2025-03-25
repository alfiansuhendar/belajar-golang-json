package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
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
