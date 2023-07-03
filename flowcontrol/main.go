package main

import "fmt"

func main() {
	//kayakPrice := 275.00

	//fmt.Println("Price:", kayakPrice) //Price: 275

	// if

	// if kayakPrice > 100 {
	// 	fmt.Println("Цена превышает 100") // ОТВЕТ: Цена превышает 100
	// }
	// будет работать:
	// if (kayakPrice > 100) {
	// 	fmt.Println("Цена превышает 100") // ОТВЕТ: Цена превышает 100
	// }

	//не работает :
	// if kayakPrice > 100
	// {
	// 	fmt.Println("Цена превышает 100") //syntax error: unexpected newline, expecting { after if clause
	// }
	// if (kayakPrice > 100                        // syntax error: unexpected newline, expecting )
	// 	&& kayakPrice < 500) {                 //syntax error: unexpected ), expecting {
	// 	fmt.Println("Цена превышает 100")
	// }

	// if kayakPrice > 500 {
	// 	fmt.Println("Цена превышает 500")
	// } else if kayakPrice < 300 {
	// 	fmt.Println("Цена меньше 300")
	// }

	// ОТВЕТ: Цена меньше 300

	// if kayakPrice > 500 {
	// 	fmt.Println("Цена превышает 500")
	// } else if kayakPrice < 100 {
	// 	fmt.Println("Цена меньше 100")
	// } else if kayakPrice > 200 && kayakPrice < 300 {
	// 	fmt.Println("Цена между 200 и 300")
	// }
	// ОТВЕТ: Цена между 200 и 30

	// if kayakPrice > 500 {
	// 	fmt.Println("Цена превышает 500")
	// } else if kayakPrice < 100 {
	// 	fmt.Println("Цена меньше 100")
	// } else {
	// 	fmt.Println("Цена, не совпадающая с предыдущими выражениями")
	// }

	// ОТВЕТ: Цена, не совпадающая с предыдущими выражениями

	//переменные в блоках локальны для своего предложения

	// if kayakPrice > 500 {
	// 	scopedVar := 500
	// 	fmt.Println("Цена превышает", scopedVar)
	// } else if kayakPrice < 100 {
	// 	scopedVar := "Цена меньше 100"
	// 	fmt.Println(scopedVar)
	// } else {
	// 	scopedVar := false
	// 	fmt.Println("соответствие", scopedVar)
	// }

	// ОТВЕТ: соответствие false

	// priceString := "275"

	// if kayakPrice, err := strconv.Atoi(priceString); err == nil {
	// 	fmt.Println("Price:", kayakPrice)
	// } else {
	// 	fmt.Println("Error:", err)
	// }
	// ОТВЕТ: Price: 275

	//for

	// counter := 0
	// for {
	// 	fmt.Println("Counter:", counter)
	// 	counter++
	// 	if counter > 3 {
	// 		break
	// 	}
	// }

	// ОТВЕТ: Counter: 0
	//        Counter: 1
	// 		  Counter: 2
	// 		  Counter: 3

	//условие можно задать после ключевого слова for
	// counter := 0

	// for counter <= 3 {
	// 	fmt.Println("Counter:", counter)
	// 	counter++
	// }

	// ОТВЕТ: Counter: 0
	//        Counter: 1
	// 		  Counter: 2
	// 		  Counter: 3

	// так же это может выглядеть

	// for counter := 0; counter <= 3; counter++ {
	// 	fmt.Println("Counter:", counter)
	// }

	// ОТВЕТ: Counter: 0
	//        Counter: 1
	// 		  Counter: 2
	// 		  Counter: 3

	// В GO нет цикла DO... WHILE, но при необходимсти всегда можно с имитировать:

	// for counter := 1; true; counter++ {
	// 	fmt.Println("Counter:", counter)
	// 	if counter > 0 {
	// 		break
	// 	}
	// }
	// ОТВЕТ: Counter: 1 один раз цикл отработает

	// for counter := 0; counter <= 3; counter++ {
	// 	if counter == 1 {
	// 		continue
	// 	}
	// 	fmt.Println("Counter:", counter)
	// }

	// ОТВЕТ: Counter: 0
	// 		  Counter: 2
	// 		  Counter: 3

	//for range

	// product := "kayak"

	// for index, character := range product {
	// 	fmt.Println("index:", index, "Character:", string(character))
	// }
	// ОТВЕТ:
	//  	index: 0 Character: k
	// 		index: 1 Character: a
	// 		index: 2 Character: y
	// 		index: 3 Character: a
	// 		index: 4 Character: k

	// for index := range product {
	// 	fmt.Println("index:", index)
	// }
	// ОТВЕТ:
	//  	index: 0
	// 		index: 1
	// 		index: 2
	// 		index: 3
	// 		index: 4

	// for _, character := range product {
	// 	fmt.Println("Character:", string(character))
	// }
	// ОТВЕТ:
	//  	Character: k
	// 		Character: a
	// 		Character: y
	// 		Character: a
	// 		Character: k

	// products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	// for index, element := range products {
	// 	fmt.Println("Index:", index, "Element", element)
	// }
	// ОТВЕТ:
	// 		Index: 0 Element Kayak
	// 		Index: 1 Element Lifejacket
	// 		Index: 2 Element Soccer Ball

	// switch

	// product := "Kayak"
	// for index, character := range product {
	// 	switch character {
	// 	case 'K':
	// 		fmt.Println("K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }
	// ОТВЕТ:
	// 		K at position 0
	// 		y at position 2

	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		fmt.Println("K or k at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }
	// ОТВЕТ:
	// 		K or k at position 0
	// 		y at position 2
	// 		K or k at position 4

	// пример использования break
	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		if character == 'k' {
	// 			fmt.Println("Lowercase k at position", index)
	// 			break
	// 		}
	// 		fmt.Println("Uppercase K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }
	// ОТВЕТ:
	// 		Uppercase K at position 0
	// 		y at position 2
	// 		Lowercase k at position 4

	// пример использования fallthrough
	// for index, character := range product {
	// 	switch character {
	// 	case 'K':
	// 		fmt.Println("Uppercase character")
	// 		fallthrough
	// 	case 'k':
	// 		fmt.Println("k at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	}
	// }
	// ОТВЕТ: (обрати внимание на "position")
	// 		Uppercase character
	// 		k at position 0
	// 		y at position 2
	// 		k at position 4

	//пример использования default
	// for index, character := range product {
	// 	switch character {
	// 	case 'K', 'k':
	// 		if character == 'k' {
	// 			fmt.Println("Lowercase k at position", index)
	// 			break
	// 		}
	// 		fmt.Println("Uppercase K at position", index)
	// 	case 'y':
	// 		fmt.Println("y at position", index)
	// 	default:
	// 		fmt.Println("Character", string(character), "at position", index)
	// 	}
	// }
	// ОТВЕТ:
	// 		Uppercase K at position 0
	// 		Character a at position 1
	// 		y at position 2
	// 		Character a at position 3
	// 		Lowercase k at position 4

	// for counter := 0; counter < 20; counter++ {
	// 	switch counter / 2 {
	// 	case 2, 3, 5, 7:
	// 		fmt.Println("Prime value:", counter/2)
	// 	default:
	// 		fmt.Println("Non-prime value:", counter/2)
	// 	}
	// }

	// от повторного выполнения операции можно отказаться:
	// for counter := 0; counter < 20; counter++ {
	// 	switch val := counter / 2; val {
	// 	case 2, 3, 5, 7:
	// 		fmt.Println("Prime value:", val)
	// 	default:
	// 		fmt.Println("Non-prime value:", val)
	// 	}
	// }

	// ОТВЕТ: (одинаковый в обоих случаях)
	// Non-prime value: 0
	// Non-prime value: 0
	// Non-prime value: 1
	// Non-prime value: 1
	// Prime value: 2
	// Prime value: 2
	// Prime value: 3
	// Prime value: 3
	// Non-prime value: 4
	// Non-prime value: 4
	// Prime value: 5
	// Prime value: 5
	// Non-prime value: 6
	// Non-prime value: 6
	// Prime value: 7
	// Prime value: 7
	// Non-prime value: 8
	// Non-prime value: 8
	// Non-prime value: 9
	// Non-prime value: 9

	// В отличии от JS и PHP в операторах case можно использовать выражения
	// for counter := 0; counter < 10; counter++ {
	// 	switch {
	// 	case counter == 0:
	// 		fmt.Println("Zero value")
	// 	case counter < 3:
	// 		fmt.Println(counter, "is < 3")
	// 	case counter >= 3 && counter < 7:
	// 		fmt.Println(counter, "is >= 3 && < 7")
	// 	default:
	// 		fmt.Println(counter, "is >= 7")
	// 	}
	// }
	// ОТВЕТ:
	// 		1 is < 3
	// 		2 is < 3
	// 		3 is >= 3 && < 7
	// 		4 is >= 3 && < 7
	// 		5 is >= 3 && < 7
	// 		6 is >= 3 && < 7
	// 		7 is >= 7
	// 		8 is >= 7
	//		9 is >= 7

	// goto

	counter := 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
	// ОТВЕТ:
	// 		Counter 0
	// 		Counter 1
	// 		Counter 2
	// 		Counter 3
	// 		Counter 4
}
