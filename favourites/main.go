package main

import (
	"fmt"
	"strings"
)

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

func isFavouriteContainsString(s string) bool {
	return strings.Contains("ABCDEMNXYZ", s)
}

func main() {
	for i := 0; i < 127; i++ {
		s := "" + string(rune(i))
		results := []bool{
			isFavourite1(s),
			isFavourite2(s),
			isFavouriteSwitch(s),
			isFavouriteSwitchByte(s),
			isFavouriteContains(s),
			isFavouriteContainsGlobal(s),
			isFavouriteContainsByte(s),
			isFavouriteContainsString(s),
			isFavouriteContainsRune(s),
		}
		fmt.Printf("isFavourite(%q) = %v\n", s, results)
		for _, r := range results[1:] {
			if results[0] != r {
				panic("functions are not equivalent")
			}
		}
	}
}
