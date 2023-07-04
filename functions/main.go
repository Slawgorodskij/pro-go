package main

import "fmt"

// func printPrice() {
// 	kayakPrice := 275.00
// 	kayakTax := kayakPrice * 0.2
// 	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
// }

// func printPrice(product string, price float64, taxRate float64) { // можно записать:  func printPrice(product string, price, taxRate float64) у price будет тип такой же как у taxRate
// 	taxAmount := price * taxRate
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func printPrice(product string, price float64, _ float64) { // возможно использование пустого идентификатора
// 	taxAmount := price * 0.25
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func printPrice(string, float64, float64) { // могут опускать имена всех параметров
// 	fmt.Println("Нет параметров")
// }

// func printSuppliers(product string, suppliers []string) {
// 	for _, supplier := range suppliers {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

//Выше описаная функция с использованием переменных параметров (должен быть последний и может использоваться тольк один тип)
// func printSuppliers(product string, suppliers ...string) {
// 	for _, supplier := range suppliers {
// 		fmt.Println("Product:", product, "Supplier:", supplier)
// 	}
// }

// работа без аргументов для вариационного параметра
// func printSuppliers(product string, suppliers ...string) {
// 	if len(suppliers) == 0 {
// 		fmt.Println("Product:", product, "Supplier: (none)")
// 	} else {
// 		for _, supplier := range suppliers {
// 			fmt.Println("Product:", product, "Supplier:", supplier)
// 		}
// 	}
// }

//использование указателей
// func swapValue(first, second int) {
// 	fmt.Println("Before swap:", first, second) // Before swap: 10 20
// 	temp := first
// 	first = second
// 	second = temp
// 	fmt.Println("After swap:", first, second) // After swap: 20 10
// }

// func swapValue(first, second *int) {
// 	fmt.Println("Before swap:", *first, *second) // Before swap: 10 20
// 	temp := *first
// 	*first = *second
// 	*second = temp
// 	fmt.Println("After swap:", *first, *second) // After swap: 20 10
// }

// Определение и сипользование результатов функции
// func calcTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

//Возврат функцией нескольких результатов
// func swapValues(first, second int) (int, int) {
// 	return second, first
// }

//Использование нескольких результатов вместо нескольких значений
// func calcTax(price float64) float64 {
// 	if price > 100 {
// 		return price * 0.2
// 	}
// 	return -1
// }
// func calcTax(price float64) (float64, bool) {
// 	if price > 100 {
// 		return price * 0.2, true
// 	}
// 	return 0, false
// }

//использование именованых результатов
// func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
// 	total = minSpend
// 	for _, price := range products {
// 		if taxAmount, due := calcTax(price); due {
// 			total += taxAmount
// 			tax += taxAmount
// 		} else {
// 			total += price
// 		}
// 	}
// 	return
// }

//использование пустого идентификатора для сброса результатов
// func calcTotalPrice(products map[string]float64) (count int, total float64) {
// 	count = len(products)
// 	for _, price := range products {
// 		total += price
// 	}
// 	return
// }

//использование ключевого слова defer
func calcTotalPrice(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	//fmt.Println("Hello, Functions")

	// fmt.Println("Собираюсь вызвать функцию")
	// printPrice()
	// fmt.Println("функция завершена")

	// printPrice("Kayak", 275, 0.2)
	// printPrice("Lifejacket", 48.95, 0.2)
	// printPrice("Soccer Ball", 19.50, 0.15)

	// printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	// printSuppliers("Lifejacket", []string{"Sail Safe Co"})

	//вызов функции использующей переменные параметры
	//printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")

	// изменим вызов создав срез
	// names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	// printSuppliers("Kayak", names...)

	// printSuppliers("Lifejacket", "Sail Safe Co")
	// printSuppliers("Soccer Ball") //без проверки на nil не выполнится

	//использование указателей
	//val1, val2 := 10, 20
	// fmt.Println("Before calling function", val1, val2) // Before calling function 10 20
	// swapValue(val1, val2)
	// fmt.Println("After calling function", val1, val2) // After calling function 10 20

	// fmt.Println("Before calling function", val1, val2) // Before calling function 10 20
	// swapValue(&val1, &val2)
	// fmt.Println("After calling function", val1, val2) // After calling function 20 10

	// Определение и сипользование результатов функции
	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }

	// for product, price := range products {
	// 	priceWithTax := calcTax(price)
	// 	fmt.Println("Product:", product, "Price:", priceWithTax)
	// }
	//так лаконичнее
	// for product, price := range products {
	// 	fmt.Println("Product:", product, "Price:", calcTax(price))
	// }

	//Возврат функцией нескольких результатов
	// val1, val2 := 10, 20
	// fmt.Println("Before calling function", val1, val2) // Before calling function 10 20
	// val1, val2 = swapValues(val1, val2)
	// fmt.Println("After calling function", val1, val2) // After calling function 20 10

	//Использование нескольких результатов вместо нескольких значений
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	// for product, price := range products {
	// 	tax := calcTax(price)
	// 	if tax != -1 {
	// 		fmt.Println("product:", product, "Tax:", tax)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due (Налог не подлежит уплате)")
	// 	}
	// }

	//второй вариант
	// for product, price := range products {
	// 	taxAmount, taxDue := calcTax(price)
	// 	if taxDue {
	// 		fmt.Println("product:", product, "Tax:", taxAmount)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due (Налог не подлежит уплате)")
	// 	}
	// }

	// и более лаконичный
	// for product, price := range products {
	// 	if taxAmount, taxDue := calcTax(price); taxDue {
	// 		fmt.Println("product:", product, "Tax:", taxAmount)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due (Налог не подлежит уплате)")
	// 	}
	// }

	//использование именованых результатов

	// total1, tax1 := calcTotalPrice(products, 10)
	// fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	// total2, tax2 := calcTotalPrice(nil, 10)
	// fmt.Println("Total 2:", total2, "Tax 2:", tax2)

	//использование пустого идентификатора для сброса результатов
	_, total := calcTotalPrice(products)
	fmt.Println("Total:", total)
}
