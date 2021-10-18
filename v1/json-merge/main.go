package main

import (
	"encoding/json"
	"fmt"

	"github.com/imdario/mergo"
)

type Foo struct {
	ID   int64
	Name string
}

type Bar struct {
	Type string
}

func main() {
	foo := Foo{
		ID:   1,
		Name: "Fred",
	}
	bar := Bar{
		Type: "human",
	}
	fooJSON, err1 := json.Marshal(foo)
	panicErr(err1)
	var fooMap map[string]interface{}
	err1 = json.Unmarshal(fooJSON, &fooMap)
	panicErr(err1)

	barJSON, err1 := json.Marshal(bar)
	panicErr(err1)
	var barMap map[string]interface{}
	err1 = json.Unmarshal(barJSON, &barMap)
	panicErr(err1)

	var destMap map[string]interface{}
	err1 = mergo.Merge(&destMap, fooMap)
	panicErr(err1)
	err1 = mergo.Merge(&destMap, barMap)
	panicErr(err1)
	fmt.Println(destMap)
	destJSON, err1 := json.MarshalIndent(destMap, "", "  ")
	panicErr(err1)
	fmt.Println(string(destJSON))
}

func panicErr(err error) {
	if err != nil {
		panic(fmt.Sprintln("argh ", err))
	}
}
