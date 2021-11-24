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

func (p *Point) MarshalJSON() ([]byte, error) {
	return []byte(`["hi"]`), nil
}

type Person struct {
	Point    Point
	Name     string `json:"name"`
	AgeYears int    `json:"age"`
	SSN      int    `json:"ssn"`
}

func (p *Person) MarshalJSON() ([]byte, error) {
	return []byte(`["hi"]`), nil
}

func main() {
	p := Person{
		Point:    Point{2, 3},
		Name:     "Steven Shaw",
		AgeYears: 21,
		SSN:      1000000001,
	}

	marshaled, err := json.Marshal(p)
	if err == nil {
		fmt.Println("p =", string(marshaled))
	} else {
		fmt.Println(err)
	}

	marshaled, err = json.Marshal(&p)
	if err == nil {
		fmt.Println("&p =", string(marshaled))
	} else {
		fmt.Println(err)
	}
}
