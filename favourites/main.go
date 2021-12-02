package main

import "fmt"

func isFavourite1(s string) bool {
	favs := map[string]struct{}{
		"A": {},
		"B": {},
		"C": {},
		"D": {},
		"E": {},
		"M": {},
		"N": {},
		"X": {},
		"Y": {},
		"Z": {},
	}
	_, result := favs[s]
	return result
}

var (
	globalFavs = map[string]struct{}{
		"A": {},
		"B": {},
		"C": {},
		"D": {},
		"E": {},
		"M": {},
		"N": {},
		"X": {},
		"Y": {},
		"Z": {},
	}
)

func isFavourite2(s string) bool {
	_, result := globalFavs[s]
	return result
}

func main() {
	for i := 0; i < 127; i++ {
		s := "" + string(rune(i))
		fmt.Printf("isFavourite(%q) = %v\n",
			s,
			isFavourite1(s),
		)
	}
}
