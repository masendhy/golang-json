package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName" :"masendhy","Address":"Blora","Age":42}`
	jsonBytes := []byte(jsonString)

	customer :=
		&Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)
}
