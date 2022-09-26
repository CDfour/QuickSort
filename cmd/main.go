package main

import (
	"fmt"
	"math/rand"
	"time"

	qs "quick_sort_go/pkg/quick_sort"
)

// Пример использования
func main() {
	slice := make([]int, 10)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100)
	}

	fmt.Println(slice)

	qs.QuickSort(slice)

	fmt.Println(slice)
}
