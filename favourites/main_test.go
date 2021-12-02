package main

import (
	"testing"
)

var result bool

func benchmarkFav(
	b *testing.B,
	isFav func(string) bool,
	s string,
) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = isFav(s)
	}
	result = r
}

func BenchmarkFavAInSet(b *testing.B)             { benchmarkFav(b, isFavouriteInSet, "A") }
func BenchmarkFavAInSetGlobal(b *testing.B)       { benchmarkFav(b, isFavouriteInSetGlobal, "A") }
func BenchmarkFavASwitch(b *testing.B)            { benchmarkFav(b, isFavouriteSwitch, "A") }
func BenchmarkFavASwitchFallthrough(b *testing.B) { benchmarkFav(b, isFavouriteSwitchFallthrough, "A") }
func BenchmarkFavASwitchByte(b *testing.B)        { benchmarkFav(b, isFavouriteSwitchByte, "A") }
func BenchmarkFavAContains(b *testing.B)          { benchmarkFav(b, isFavouriteContains, "A") }
func BenchmarkFavAContainsGlobal(b *testing.B)    { benchmarkFav(b, isFavouriteContainsGlobal, "A") }
func BenchmarkFavASearchStrings(b *testing.B)     { benchmarkFav(b, isFavouriteSearchStrings, "A") }
func BenchmarkFavAStringContains(b *testing.B)    { benchmarkFav(b, isFavouriteStringContains, "A") }
func BenchmarkFavAStringContainsRune(b *testing.B) {
	benchmarkFav(b, isFavouriteStringContainsRune, "A")
}
func BenchmarkFavAContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "A") }
func BenchmarkFavAKoazeeContains(b *testing.B) { benchmarkFav(b, isFavouriteKoazeeContains, "A") }
func BenchmarkFavAFunkContains(b *testing.B)   { benchmarkFav(b, isFavouriteFunkContains, "A") }

func BenchmarkFavEInSet(b *testing.B)             { benchmarkFav(b, isFavouriteInSet, "E") }
func BenchmarkFavEInSetGlobal(b *testing.B)       { benchmarkFav(b, isFavouriteInSetGlobal, "E") }
func BenchmarkFavESwitch(b *testing.B)            { benchmarkFav(b, isFavouriteSwitch, "E") }
func BenchmarkFavESwitchFallthrough(b *testing.B) { benchmarkFav(b, isFavouriteSwitchFallthrough, "E") }
func BenchmarkFavESwitchByte(b *testing.B)        { benchmarkFav(b, isFavouriteSwitchByte, "E") }
func BenchmarkFavEContains(b *testing.B)          { benchmarkFav(b, isFavouriteContains, "E") }
func BenchmarkFavEContainsGlobal(b *testing.B)    { benchmarkFav(b, isFavouriteContainsGlobal, "E") }
func BenchmarkFavESearchStrings(b *testing.B)     { benchmarkFav(b, isFavouriteSearchStrings, "E") }
func BenchmarkFavEStringContains(b *testing.B)    { benchmarkFav(b, isFavouriteStringContains, "E") }
func BenchmarkFavEStringContainsRune(b *testing.B) {
	benchmarkFav(b, isFavouriteStringContainsRune, "E")
}
func BenchmarkFavEContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "E") }
func BenchmarkFavEKoazeeContains(b *testing.B) { benchmarkFav(b, isFavouriteKoazeeContains, "E") }
func BenchmarkFavEFunkContains(b *testing.B)   { benchmarkFav(b, isFavouriteFunkContains, "E") }

func BenchmarkFavYInSet(b *testing.B)             { benchmarkFav(b, isFavouriteInSet, "Y") }
func BenchmarkFavYInSetGlobal(b *testing.B)       { benchmarkFav(b, isFavouriteInSetGlobal, "Y") }
func BenchmarkFavYSwitch(b *testing.B)            { benchmarkFav(b, isFavouriteSwitch, "Y") }
func BenchmarkFavYSwitchFallthrough(b *testing.B) { benchmarkFav(b, isFavouriteSwitchFallthrough, "Y") }
func BenchmarkFavYSwitchByte(b *testing.B)        { benchmarkFav(b, isFavouriteSwitchByte, "Y") }
func BenchmarkFavYContains(b *testing.B)          { benchmarkFav(b, isFavouriteContains, "Y") }
func BenchmarkFavYContainsGlobal(b *testing.B)    { benchmarkFav(b, isFavouriteContainsGlobal, "Y") }
func BenchmarkFavYSearchStrings(b *testing.B)     { benchmarkFav(b, isFavouriteSearchStrings, "Y") }
func BenchmarkFavYStringContains(b *testing.B)    { benchmarkFav(b, isFavouriteStringContains, "Y") }
func BenchmarkFavYStringContainsRune(b *testing.B) {
	benchmarkFav(b, isFavouriteStringContainsRune, "Y")
}
func BenchmarkFavYContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "Y") }
func BenchmarkFavYKoazeeContains(b *testing.B) { benchmarkFav(b, isFavouriteKoazeeContains, "Y") }
func BenchmarkFavYFunkContains(b *testing.B)   { benchmarkFav(b, isFavouriteFunkContains, "Y") }

func BenchmarkFavaInSet(b *testing.B)             { benchmarkFav(b, isFavouriteInSet, "a") }
func BenchmarkFavaInSetGlobal(b *testing.B)       { benchmarkFav(b, isFavouriteInSetGlobal, "a") }
func BenchmarkFavaSwitch(b *testing.B)            { benchmarkFav(b, isFavouriteSwitch, "a") }
func BenchmarkFavaSwitchFallthrough(b *testing.B) { benchmarkFav(b, isFavouriteSwitchFallthrough, "a") }
func BenchmarkFavaSwitchByte(b *testing.B)        { benchmarkFav(b, isFavouriteSwitchByte, "a") }
func BenchmarkFavaContains(b *testing.B)          { benchmarkFav(b, isFavouriteContains, "a") }
func BenchmarkFavaContainsGlobal(b *testing.B)    { benchmarkFav(b, isFavouriteContainsGlobal, "a") }
func BenchmarkFavaSearchStrings(b *testing.B)     { benchmarkFav(b, isFavouriteSearchStrings, "a") }
func BenchmarkFavaStringContains(b *testing.B)    { benchmarkFav(b, isFavouriteStringContains, "a") }
func BenchmarkFavaStringContainsRune(b *testing.B) {
	benchmarkFav(b, isFavouriteStringContainsRune, "a")
}
func BenchmarkFavaContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "a") }
func BenchmarkFavaKoazeeContains(b *testing.B) { benchmarkFav(b, isFavouriteKoazeeContains, "a") }
func BenchmarkFavaFunkContains(b *testing.B)   { benchmarkFav(b, isFavouriteFunkContains, "a") }
