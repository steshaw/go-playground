package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	nowGMT := now.In(time.FixedZone("GMT", 0))
	nowUTC := now.UTC()

	fmt.Printf("now    = %v\n", now)
	fmt.Printf("nowGMT = %v\n", nowGMT)
	fmt.Printf("nowUTC = %v\n", nowUTC)
	println()

	fmt.Printf("now    = %+v\n", now)
	fmt.Printf("nowGMT = %+v\n", nowGMT)
	fmt.Printf("nowUTC = %+v\n", nowUTC)
}
