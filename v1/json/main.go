// From https://talks.golang.org/2017/state-of-go.slide#6
package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Person struct {
	Point
	Name     string `json:"name"`
	AgeYears int    `json:"age"`
	SSN      int    `json:"ssn"`
}

func main() {
	// And that for some reason, like JSON you also have:
	aux := Person{
		Point:    Point{2, 3},
		Name:     "Steven Shaw",
		AgeYears: 21,
		SSN:      1000000001,
	}

	marshaled, err := json.Marshal(aux)
	if err == nil {
		fmt.Println(string(marshaled))
	}
}
