package main

import "fmt"

func main() {

	PrintHello()

	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}
func PrintHello() {
	fmt.Println("Hello, GO")
}
func PrintNumber(number int) {
	fmt.Println(number)
}
// go fmt - форматирование текста