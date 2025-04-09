package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		Firstname:  "Eko",
		Middlename: "Kurniawan",
		Lastname:   "Khannedy",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Firstname":"Eko","Middlename":"Kurniawan","Lastname":"Khannedy","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n", customer)
	fmt.Println("\n", customer.Firstname)
	fmt.Println("\n", customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalanin tanah",
				Country:    "Indonesia",
				PostalCode: "15147",
			},
			{
				Street:     "Jalanin aspal",
				Country:    "Indonesia",
				PostalCode: "15148",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"Eko","Middlename":"","Lastname":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalanin tanah","Country":"Indonesia","PostalCode":"15147"},{"Street":"Jalanin aspal","Country":"Indonesia","PostalCode":"15148"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n", customer)
	fmt.Println("\n", customer.Firstname)
	fmt.Println("\n", customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalanin tanah","Country":"Indonesia","PostalCode":"15147"},{"Street":"Jalanin aspal","Country":"Indonesia","PostalCode":"15148"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n", addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Jalanin tanah",
			Country:    "Indonesia",
			PostalCode: "15147",
		},
		{
			Street:     "Jalanin aspal",
			Country:    "Indonesia",
			PostalCode: "15148",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
