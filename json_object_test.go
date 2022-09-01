package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname  string
	MiddleName string
	LastName   string
	Address    string
	Age        int
	Hobbies    []string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		Firstname: "masendhy",
		Address:   "Blora",
		Age:       42,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
