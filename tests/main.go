package main

import (
	"log"
	"sort"
)

/*
func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
		total += val
		//total++
	}
	return
}
func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	log.Print("Sorted Data: ", sorted)
	log.Print("Total: ", total)
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
*/
/*
 ответ содержит ошибку
Sorted Data: [1 7 20 84 100]
Total: 217
*/

// Создание пользовательских регистраторов
func sortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), "sortAndTotal: ", log.Flags()|log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	logger.Printf("Sorted data: %v", sorted)
	for _, val := range sorted {
		total += val
	}
	logger.Printf("Total: %v", total)
	return
}

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	log.Print("Sorted Data: ", sorted)
	log.Print("Total: ", total)
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
