package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		sortAndTotal(data)
	}
}
*/
//Удаление установки из теста
/*
func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sortAndTotal(data)
	}
}
*/
// Выполнение суббенчмаркингов

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{10, 100, 250}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array Size %v", size), func(subB *testing.B) {
			data := make([]int, size)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rand.Int()
				}
				subB.StartTimer()
				sortAndTotal(data)
			}
		})
	}
}
