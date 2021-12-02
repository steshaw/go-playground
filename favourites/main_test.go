package main

import (
	"math/rand"
	"testing"
)

var result bool

func randMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}

func randS() string {
	//return string(rune(randMinMax(32, 126)))
	return "A"
}

func BenchmarkIsFavourite1(b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = isFavourite1(randS())
	}
	result = r
}

func BenchmarkIsFavourite2(b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = isFavourite2(randS())
	}
	result = r
}
