package main

import (
	"fmt"
	"sort"
)

//"math/rand"

func main() {
	// fmt.Println("Value:", rand.Int())

	// fmt.Println("Hello, Go")
	// fmt.Println(20 + 20)
	// fmt.Println(30 + 20)

	// const price float32 = 275.00
	// const tax float32 = 27.50

	// fmt.Println(price + tax)

	//const quantity int = 2
	//fmt.Println("Total: ", quantity*(price+tax)) //не выполнится ./main.go:19:25: invalid operation: quantity * (price + tax) (mismatched types int and float32)

	// для выполнения подобного необходимо использовать не типизированную константу
	//const quantity = 2
	//fmt.Println("Total: ", quantity*(price+tax)) // так выполнится - Total:  605

	// константы можно объявлять :

	// const price, tax float32 = 275.00, 27.50
	// const quantity, inStock = 2, true

	// fmt.Println("Total: ", quantity*(price+tax))
	// fmt.Println("In stock: ", inStock)

	// испльзование литерального значения
	//fmt.Println("Total: ", 2*quantity*(price+tax)) // функция отработает - Total:  1210

	// var price float32 = 275.00
	// var tax float32 = 25.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	// var price float32 = 275.00
	// var price2 = price
	// fmt.Println(price) //275
	// fmt.Println(price2) //275

	//использование разных типов
	//var price = 275.00
	//var tax float32 = 275.00

	//fmt.Println(price + tax) //не выполнится из-за разных типов ./main.go:52:14: invalid operation: price + tax (mismatched types float64 and float32)

	//пропуск присвоения значения переменной
	//var price float32  //при отсутствии начального значения опускать тип НЕЛЬЗЯ!!!
	//fmt.Println(price) // 0
	//price = 275.00
	//fmt.Println(price) // 275

	//определение нескольих переменных
	//var price, tax float32 = 275.00, 27.50 // начальное значение определит тип переменной
	//fmt.Println(price + tax)               // ответ 302.5

	//использование краткого синтаксиса (можно использовать только внутри функций)
	//price := 275.00
	//fmt.Println(price) // выведется - 275

	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)

	//Переопределение переменных
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)

	// //var price2, tax = 200.00, 25.00 //недопустимо
	// //fmt.Println("Total 2:", price2 + tax) // tax redeclared in this block

	// price2, tax := 200.00, 25.00 //допустимо
	// fmt.Println("Total 2:", price2 + tax) // код отработает = 225

	// использование пустого идентификатора
	// price, tax, inStock, discount := 275.00, 27.50, true, true
	// var salesPerson = "Alice"
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)
	// ошибки: "discount declared but not used" и "salesPerson declared but not used" вызваны наличием
	// неиспользуемых переменных

	// если нет возможности удалить неиспользуемую переменную,
	// то можно использовать "пустой идентификатор"

	// price, tax, inStock, _ := 275.00, 27.50, true, true
	// var _ = "Alice"
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock) // код отработает как положено.

	//Понимание указателей

	// first := 100
	// second := first
	// first++
	// fmt.Println("First:", first)   // First: 101
	// fmt.Println("Second:", second) //Second: 100

	// Переменные не завясят друг от друга и ссылаются на разные ячейки памяти
	// Указатель - это переменная, значением которой является адрес памяти.

	// first := 100
	// var second *int = &first
	// first++
	// fmt.Println("First:", first)   // First: 101
	// fmt.Println("Second:", second) //Second: 0xc000014110 - выводится адрес памяти

	// *int - указывает что переменная имеет значение - "адрес памяти"
	// с помощью символа * можно прочитать значение по адресу памяти

	// first := 100
	// var second = &first
	// first++
	// fmt.Println("First:", first)    // First: 101
	// fmt.Println("Second:", *second) //Second: 101

	// *second++
	// fmt.Println("First:", first)    // First: 102
	// fmt.Println("Second:", *second) //Second: 102

	// var myNewPointer *int
	// myNewPointer = second
	// *myNewPointer++

	// fmt.Println("First:", first)    // First: 103
	// fmt.Println("Second:", *second) //Second: 103

	// first := 100
	// var second *int

	// fmt.Println("Second:", second) //Second: <nil>
	// //fmt.Println("Second:", *second) //panic: runtime error: invalid memory address or nil pointer dereference

	// second = &first

	// fmt.Println("Second:", second)  //Second: 0xc000014110
	// fmt.Println("Second:", *second) //Second: 100

	// first := 100
	// second :=&first
	// third :=&second

	// fmt.Println("First:", first) //First: 100
	// fmt.Println("Second:", *second)  //Second: 100
	// fmt.Println("Third:", **third) //Third: 100

	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	secondPosition := &names[1]
	fmt.Println(secondName)      // Charlie
	fmt.Println(*secondPosition) // Charlie
	sort.Strings(names[:])
	fmt.Println(secondName)      // Charlie сортировка не влияет на значение переменной "secondName"
	fmt.Println(*secondPosition) // Bob сортировка влияет на значение переменной "secondPosition" тк ссылается на ячейку памяти
}
