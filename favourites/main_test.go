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
	for n := 0; n < b.N; n++ {
		isFav(s)
	}
}

func BenchmarkFavA1(b *testing.B)              { benchmarkFav(b, isFavourite1, "A") }
func BenchmarkFavA2(b *testing.B)              { benchmarkFav(b, isFavourite2, "A") }
func BenchmarkFavASwitch(b *testing.B)         { benchmarkFav(b, isFavouriteSwitch, "A") }
func BenchmarkFavASwitchByte(b *testing.B)     { benchmarkFav(b, isFavouriteSwitchByte, "A") }
func BenchmarkFavAContains(b *testing.B)       { benchmarkFav(b, isFavouriteContains, "A") }
func BenchmarkFavAContainsGlobal(b *testing.B) { benchmarkFav(b, isFavouriteContainsGlobal, "A") }
func BenchmarkFavAContainsString(b *testing.B) { benchmarkFav(b, isFavouriteContainsString, "A") }
func BenchmarkFavAContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "A") }

func BenchmarkFavE1(b *testing.B)              { benchmarkFav(b, isFavourite1, "E") }
func BenchmarkFavE2(b *testing.B)              { benchmarkFav(b, isFavourite2, "E") }
func BenchmarkFavESwitch(b *testing.B)         { benchmarkFav(b, isFavouriteSwitch, "E") }
func BenchmarkFavESwitchByte(b *testing.B)     { benchmarkFav(b, isFavouriteSwitchByte, "E") }
func BenchmarkFavEContains(b *testing.B)       { benchmarkFav(b, isFavouriteContains, "E") }
func BenchmarkFavEContainsGlobal(b *testing.B) { benchmarkFav(b, isFavouriteContainsGlobal, "E") }
func BenchmarkFavEContainsString(b *testing.B) { benchmarkFav(b, isFavouriteContainsString, "E") }
func BenchmarkFavEContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "E") }

func BenchmarkFavY1(b *testing.B)              { benchmarkFav(b, isFavourite1, "Y") }
func BenchmarkFavY2(b *testing.B)              { benchmarkFav(b, isFavourite2, "Y") }
func BenchmarkFavYSwitch(b *testing.B)         { benchmarkFav(b, isFavouriteSwitch, "Y") }
func BenchmarkFavYSwitchByte(b *testing.B)     { benchmarkFav(b, isFavouriteSwitchByte, "Y") }
func BenchmarkFavYContains(b *testing.B)       { benchmarkFav(b, isFavouriteContains, "Y") }
func BenchmarkFavYContainsGlobal(b *testing.B) { benchmarkFav(b, isFavouriteContainsGlobal, "Y") }
func BenchmarkFavYContainsString(b *testing.B) { benchmarkFav(b, isFavouriteContainsString, "Y") }
func BenchmarkFavYContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "Y") }

func BenchmarkFava1(b *testing.B)              { benchmarkFav(b, isFavourite1, "a") }
func BenchmarkFava2(b *testing.B)              { benchmarkFav(b, isFavourite2, "a") }
func BenchmarkFavaSwitch(b *testing.B)         { benchmarkFav(b, isFavouriteSwitch, "a") }
func BenchmarkFavaSwitchByte(b *testing.B)     { benchmarkFav(b, isFavouriteSwitchByte, "a") }
func BenchmarkFavaContains(b *testing.B)       { benchmarkFav(b, isFavouriteContains, "a") }
func BenchmarkFavaContainsGlobal(b *testing.B) { benchmarkFav(b, isFavouriteContainsGlobal, "a") }
func BenchmarkFavaContainsString(b *testing.B) { benchmarkFav(b, isFavouriteContainsString, "a") }
func BenchmarkFavaContainsRune(b *testing.B)   { benchmarkFav(b, isFavouriteContainsRune, "a") }
