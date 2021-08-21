// An example question from Gophers channel.
// We just use strings.Join.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"{port:21}", "{port:443}", "{port:31337}", "{port:9929}"}

	fmt.Println("[" + strings.Join(s, ", ") + "]")
}
