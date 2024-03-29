package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"unicode"

	"github.com/thoas/go-funk"
	"github.com/wesovilabs/koazee"
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
	globalFavsSet = map[string]struct{}{
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
	_, result := globalFavsSet[s]
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

func isFavouriteSwitchFallthrough(s string) bool {
	switch s {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fallthrough
	case "E":
		fallthrough
	case "M":
		fallthrough
	case "N":
		fallthrough
	case "X":
		fallthrough
	case "Y":
		fallthrough
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

var globalFavsStrings = []string{
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
	for _, fav := range globalFavsStrings {
		if s == fav {
			return true
		}
	}
	return false
}

func isFavouriteSearchStrings(s string) bool {
	index := sort.SearchStrings(globalFavsStrings, s)
	return index < len(globalFavsStrings) && globalFavsStrings[index] == s
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

var favsStream = koazee.StreamOf(globalFavsStrings)

func isFavouriteKoazeeContains(s string) bool {
	contained, err := favsStream.Contains(s)
	if err != nil {
		log.Fatalf("Argh Koazee.Contains failed...: %v", err)
	}
	return contained
}

func isFavouriteFunkContains(s string) bool {
	return funk.Contains(globalFavsStrings, s)
}

func main() {
	for i := 0; i <= unicode.MaxRune; i++ {
		s := "" + string(rune(i))
		results := []bool{
			isFavouriteInSet(s),
			isFavouriteInSetGlobal(s),
			isFavouriteSwitch(s),
			isFavouriteSwitchFallthrough(s),
			isFavouriteSwitchByte(s),
			isFavouriteContains(s),
			isFavouriteContainsGlobal(s),
			isFavouriteSearchStrings(s),
			isFavouriteContainsByte(s),
			isFavouriteStringContains(s),
			isFavouriteStringContainsRune(s),
			isFavouriteContainsRune(s),
			isFavouriteKoazeeContains(s),
			isFavouriteFunkContains(s),
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
