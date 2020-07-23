package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Contacts struct {
	Name  string
	Phone string
}

var PhoneBook []Contacts

func main() {
	PhoneBook = append(PhoneBook, Contacts{"Ivanov", "+7-999-111-11-11"})
	PhoneBook = append(PhoneBook, Contacts{"Petrov", "+7-999-222-22-22"})
	PhoneBook = append(PhoneBook, Contacts{"Sidorov", "+7-999-333-33-33"})

	fmt.Println(PhoneBook)

	x, _ := json.Marshal(PhoneBook)
	ioutil.WriteFile("PhoneBook", x, 0644)
}
