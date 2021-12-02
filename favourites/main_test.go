package main

import (
	"math/rand"
	"testing"
)

func randMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}

func randS() string {
	return string(rune(randMinMax(32, 126)))
}

func BenchmarkIsFavourite1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = isFavourite1(randS())
	}
}

func BenchmarkIsFavourite2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = isFavourite2(randS())
	}
}
