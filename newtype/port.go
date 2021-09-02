package main

import (
	"errors"
	"fmt"
	"math"
)

type Port uint16

func parsePort(s string) (port Port, err error) {
	n, err := fmt.Sscan(s, &port)
	if err != nil {
		return 0, err
	} else if n == 1 {
		return port, nil
	} else {
		return 0, errors.New(fmt.Sprintf("Expected a single port, got %d", n))
	}
}

const (
	maxPort = math.MaxUint16
	minPort = 0
)

func main() {
	s1 := fmt.Sprintf("%d", minPort)
	fmt.Printf("s1 = %q\n", s1)
	if p1, err := parsePort(fmt.Sprintf("%v", minPort)); err == nil {
		fmt.Println("p1", p1 == minPort)
	}

	if p2, err := parsePort("100"); err == nil {
		fmt.Println("p2", p2 == 100)
	}

	if p3, err := parsePort("8080"); err == nil {
		fmt.Println("p3", p3 == 8080)
	}

	if p4, err := parsePort("65535"); err == nil {
		fmt.Println("p4", p4 == 65535)
	}

	if p5, err := parsePort("65536"); err == nil {
		fmt.Printf("p5 = %v\n", p5)
		var i uint16 = math.MaxUint16
		fmt.Printf("i = %v\n", i)
		fmt.Println("p5", p5 == 65535)
	} else {
		println("p5 failed")
	}

	if p6, err := parsePort("-99"); err == nil {
		fmt.Println("p6", p6 == 0)
	} else {
		println("p6 failed")
	}
}
