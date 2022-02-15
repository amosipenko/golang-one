package sliceSort

import (
	"math/rand"
	"testing"
	"time"
)

func initSliceInts(sliceSize int) []int {
	slice := make([]int, sliceSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

var sliceSize = 2000
var mySlice = initSliceInts(sliceSize)
var testingSlice = make([]int, sliceSize)

func BenchmarkSortByInserts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Т.к. функция сортирует переданный слайс по значению, то чтобы на втором шаге слайс снова
		// оказался неотсортированным, будем заполнять его на каждом шаге исходным значением слайса
		copy(testingSlice, mySlice)
		SortByInserts(testingSlice)
	}
}

func BenchmarkSortByBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Т.к. функция сортирует переданный слайс по значению, то чтобы на втором шаге слайс снова
		// оказался неотсортированным, будем заполнять его на каждом шаге исходным значением слайса
		copy(testingSlice, mySlice)
		SortByBubble(testingSlice)
	}
}

// Вызовы:
//go test -bench=SortByInserts ./... --benchmem
// в среднем 1,42 сек

//go test -bench=SortByBubble ./... --benchmem
// в среднем 1,75сек
