package main

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
)

type Baz struct {
	First string
	Last string
}

type Foo struct {
	ID   int64
	Name string
	Baz Baz
	Hi string
}

type Bar struct {
	Type string
	Baz Baz
	Bye string
}

func main() {
	foo := Foo{
		ID:   1,
		Name: "Fred",
		Baz: Baz{First: "first"},
		Hi: "hi",
	}
	bar := Bar{
		Type: "human",
		Baz: Baz{Last: "last"},
	}
	fooJSON, err := json.Marshal(foo)
	panicErr(err)
	var fooMap map[string]interface{}
	err = json.Unmarshal(fooJSON, &fooMap)
	panicErr(err)

	barJSON, err := json.Marshal(bar)
	panicErr(err)
	var barMap map[string]interface{}
	err = json.Unmarshal(barJSON, &barMap)
	panicErr(err)

	fmt.Println("Merging bar")
	var destMap map[string]interface{}
	err = mergo.Merge(&destMap, fooMap)
	panicErr(err)
	err = mergo.Merge(&destMap, barMap)
	panicErr(err)
	fmt.Println(destMap)
	destJSON, err := json.MarshalIndent(destMap, "", "  ")
	panicErr(err)
	fmt.Println(string(destJSON))

	fmt.Println("Merging bar1")
	bar1 := Bar{
		Type: "dog",
		Baz: Baz{
			First: "Bob",
			Last: "Smith",
		},
		Bye: "bye",
	}
	bar1JSON, err := json.MarshalIndent(bar1, "", "  ")
	panicErr(err)
	fmt.Println(string(bar1JSON))
	var bar1Map map[string]interface{}
	err = json.Unmarshal(bar1JSON, &bar1Map)
	panicErr(err)
	fmt.Printf("destMap=%#v bar1Map=%#v\n", destMap, bar1Map)
	// Use bar1Map as destination so that it has priority.
	err = mergo.Merge(&bar1Map, destMap)
	panicErr(err)
	fmt.Println(bar1Map)
	destJSON, err = json.MarshalIndent(bar1Map, "", "  ")
	panicErr(err)
	fmt.Println(string(bar1JSON))
}

func panicErr(err error) {
	if err != nil {
		panic(fmt.Sprintln("argh ", err))
	}
}
