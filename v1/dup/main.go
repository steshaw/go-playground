package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
		// ignoring input.Err()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning", err)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
