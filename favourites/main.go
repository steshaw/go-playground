package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func isFavouriteInSet(s string) bool {
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

func isFavouriteInSetGlobal(s string) bool {
	_, result := globalFavs[s]
	return result
}

func isFavouriteSwitch(s string) bool {
	switch s {
	case "A":
		return true
	case "B":
		return true
	case "C":
		return true
	case "D":
		return true
	case "E":
		return true
	case "M":
		return true
	case "N":
		return true
	case "X":
		return true
	case "Y":
		return true
	case "Z":
		return true
	default:
		return false
	}
}

func isFavouriteSwitchByte(s string) bool {
	if len(s) != 1 {
		return false
	}
	switch s[0] {
	case 'A':
		return true
	case 'B':
		return true
	case 'C':
		return true
	case 'D':
		return true
	case 'E':
		return true
	case 'M':
		return true
	case 'N':
		return true
	case 'X':
		return true
	case 'Y':
		return true
	case 'Z':
		return true
	default:
		return false
	}
}

func isFavouriteContains(s string) bool {
	favs := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"M",
		"N",
		"X",
		"Y",
		"Z",
	}
	for _, fav := range favs {
		if s == fav {
			return true
		}
	}
	return false
}

var globalFavsSlice = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"M",
	"N",
	"X",
	"Y",
	"Z",
}

func isFavouriteContainsGlobal(s string) bool {
	for _, fav := range globalFavsSlice {
		if s == fav {
			return true
		}
	}
	return false
}

var globalFavsByte = []byte{
	'A',
	'B',
	'C',
	'D',
	'E',
	'M',
	'N',
	'X',
	'Y',
	'Z',
}

func isFavouriteContainsByte(s string) bool {
	if len(s) != 1 {
		return false
	}
	b := s[0]
	for _, fav := range globalFavsByte {
		if b == fav {
			return true
		}
	}
	return false
}

var globalFavsRune = []rune{
	'A',
	'B',
	'C',
	'D',
	'E',
	'M',
	'N',
	'X',
	'Y',
	'Z',
}

func isFavouriteContainsRune(s string) bool {
	runes := []rune(s)
	if len(runes) != 1 {
		return false
	}
	r := runes[0]
	for _, fav := range globalFavsRune {
		if r == fav {
			return true
		}
	}
	return false
}

func isFavouriteStringContains(s string) bool {
	return strings.Contains("ABCDEMNXYZ", s)
}

func isFavouriteStringContainsRune(s string) bool {
	runes := []rune(s)
	if len(runes) != 1 {
		return false
	}
	r := runes[0]
	return strings.ContainsRune("ABCDEMNXYZ", r)
}

func main() {
	for i := 0; i <= unicode.MaxRune; i++ {
		s := "" + string(rune(i))
		results := []bool{
			isFavouriteInSet(s),
			isFavouriteInSetGlobal(s),
			isFavouriteSwitch(s),
			isFavouriteSwitchByte(s),
			isFavouriteContains(s),
			isFavouriteContainsGlobal(s),
			isFavouriteContainsByte(s),
			isFavouriteStringContains(s),
			isFavouriteStringContainsRune(s),
			isFavouriteContainsRune(s),
		}
		msg := fmt.Sprintf("isFavourite(%q) = %v", s, results)
		fmt.Println(msg)
		for _, r := range results[1:] {
			if results[0] != r {
				log.Fatalf("Functions are not equivalent: %v", msg)
			}
		}
	}
}
