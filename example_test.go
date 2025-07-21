package grd_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/imadselka/grd"
)

// Example demonstrates basic usage of grd package
func ExampleTry() {
	result := grd.Try(func() (int, error) {
		return 42, nil
	}).Then(func(val int) (int, error) {
		return val * 2, nil
	}).Catch(func(err error) int {
		return -1
	})

	fmt.Println(result)
	// Output: 84
}

// Example demonstrates error handling
func ExampleTry_withError() {
	result := grd.Try(func() (int, error) {
		return 0, errors.New("something went wrong")
	}).Then(func(val int) (int, error) {
		// This won't be called due to error
		return val * 2, nil
	}).Catch(func(err error) int {
		return -1
	})

	fmt.Println(result)
	// Output: -1
}

// Example demonstrates string processing pipeline
func ExampleTry_stringProcessing() {
	input := "  hello world  "

	result := grd.Try(func() (string, error) {
		if input == "" {
			return "", errors.New("empty input")
		}
		return input, nil
	}).Then(func(s string) (string, error) {
		return strings.TrimSpace(s), nil
	}).Then(func(s string) (string, error) {
		return strings.ToUpper(s), nil
	}).Finally(func() {
		log.Println("String processing completed")
	}).Catch(func(err error) string {
		return "ERROR"
	})

	fmt.Println(result)
	// Output: HELLO WORLD
}

// Example demonstrates JSON processing
func ExampleTry_jsonProcessing() {
	jsonData := `{"name": "john", "age": 30}`

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// First, parse JSON to Person
	person := grd.Try(func() (Person, error) {
		if jsonData == "" {
			return Person{}, errors.New("empty json")
		}
		var p Person
		err := json.Unmarshal([]byte(jsonData), &p)
		return p, err
	}).Then(func(person Person) (Person, error) {
		// Transform name to title case
		person.Name = strings.Title(person.Name)
		return person, nil
	}).Catch(func(err error) Person {
		return Person{Name: "Unknown", Age: -1}
	})

	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	// Output: John is 30 years old
}
