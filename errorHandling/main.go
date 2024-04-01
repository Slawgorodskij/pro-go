package main

import "fmt"

//func main() {
//	categories := []string{"Watersports", "Chess"}
//
//	for _, cat := range categories {
//		total := Products.TotalPrice(cat)
//		fmt.Println(cat, "Total:", ToCurrency(total))
//	}
//	/*
//	Результат:
//	Watersports Total: $328.95
//	Chess Total: $1291.00
//	*/
//}

/* спровоцируем проблему*/

//func main() {
//	categories := []string{"Watersports", "Chess", "Running"}
//
//	for _, cat := range categories {
//		total := Products.TotalPrice(cat)
//		fmt.Println(cat, "Total:", ToCurrency(total))
//	}
//	/*
//		Результат:
//		Watersports Total: $328.95
//		Chess Total: $1291.00
//		Running Total: $0.00
//
//		Определить причину вывода Running Total: $0.00 может быть не очевидно, это как и отсутствие товара так и отсутствие самой категории
//	*/
//}

/*
		ГЕНЕРАЦИЯ ОШИБОК
	в файл operation.go внесены изменения
*/

//func main() {
//	categories := []string{"Watersports", "Chess", "Running"}
//
//	for _, cat := range categories {
//		total, err := Products.TotalPrice(cat)
//		if err == nil {
//			fmt.Println(cat, "Total:", ToCurrency(total))
//		} else {
//			fmt.Println(cat, "(no such category)")
//		}
//	}
//	/*
//		Результат:
//			Watersports Total: $328.95
//			Chess Total: $1291.00
//			Running (no such category)
//	*/
//}

/*
	 СООБЩЕНИЕ ОБ ОШИБКАХ ЧЕРЕЗ КАНАЛЫ
	в файл operation.go внесены изменения


	ИСПОЛЬЗОВАНИЕ УДОБНЫХ ФУНКЦИЙ ОБРАБОТОК ОШИБОК
		в файл operation.go внесены изменения

*/

//func main() {
//	categories := []string{"Watersports", "Chess", "Running"}
//
//	channel := make(chan ChannelMessage, 10)
//	go Products.TotalPriceAsync(categories, channel)
//
//	for message := range channel {
//		if message.CategoryError == nil {
//			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
//		} else {
//			fmt.Println(message.Category, "(no such category)")
//		}
//	}
//	/*
//		Результат:
//			Watersports Total: $328.95
//			Chess Total: $1291.00
//			Running (no such category)
//	*/
//}

/*
	  РАБОТА С НЕИСПРАВИМЫМИ ОШИБКАМИ
	в файл operation.go внесены изменения
*/

//func main() {
//	/*
//		recoverFunc := func() {
//			if arg := recover(); arg != nil {
//				if err, ok := arg.(error); ok {
//					fmt.Println("ERROR:", err.Error())
//				} else if str, ok := arg.(string); ok {
//					fmt.Println("Message:", str)
//				} else {
//					fmt.Println("Panic recovered")
//				}
//			}
//		}
//		defer recoverFunc()
//		более короткий вариант (используется анонимная функция)
//	*/
//	defer func() {
//		if arg := recover(); arg != nil {
//			if err, ok := arg.(error); ok {
//				fmt.Println("ERROR:", err.Error())
//				//panic(err)
//			} else if str, ok := arg.(string); ok {
//				fmt.Println("Message:", str)
//			} else {
//				fmt.Println("Panic recovered")
//			}
//		}
//	}()
//
//	categories := []string{"Watersports", "Chess", "Running"}
//
//	channel := make(chan ChannelMessage, 10)
//	go Products.TotalPriceAsync(categories, channel)
//
//	for message := range channel {
//		if message.CategoryError == nil {
//			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
//		} else {
//			panic(message.CategoryError)
//		}
//	}
//	/*
//				Результат:
//			     до добавления функции востановления
//					Watersports Total: $328.95
//					Chess Total: $1291.00
//					panic: Cannot find category: Running
//
//					goroutine 1 [running]:
//					main.main()
//					        /home/vladimir/project/pro-go/errorHandling/main.go:108 +0x2e9
//					exit status 2
//
//			    после добавления функции востановления
//					Watersports Total: $328.95
//					Chess Total: $1291.00
//					ERROR: Cannot find category: Running
//
//				В случае с паникой после востановления (стр. 118)
//					Watersports Total: $328.95
//					Chess Total: $1291.00
//					ERROR: Cannot find category: Running
//					panic: Cannot find category: Running [recovered]
//		        		panic: Cannot find category: Running
//
//					goroutine 1 [running]:
//					main.main.func1()
//							/home/vladimir/project/pro-go/errorHandling/main.go:118 +0x1c8
//					panic({0x489d60, 0xc00009e220})
//		        			/usr/lib/go-1.18/src/runtime/panic.go:838 +0x207
//					main.main()
//		        			/home/vladimir/project/pro-go/errorHandling/main.go:136 +0x313
//					exit status 2
//	*/
//}

/*
	востановление после паники в горутинах
*/

type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage{
				TerminalError: arg,
			}
			close(outChan)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {
	categories := []string{"Watersports", "Chess", "Running"}

	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)

	//for message := range channel {
	//	fmt.Println(message.Category, "Total:", message.Count)
	//}

	for message := range channel {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}

	/*
			Результат
				Watersports Total: 328
				Chess Total: 1291
				Cannot find category: Running
				fatal error: all goroutines are asleep - deadlock!

				goroutine 1 [chan receive]:
				main.main()
				        /home/vladimir/project/pro-go/errorHandling/main.go:210 +0x1a8
				exit status 2

			Ошибка связана с не возбновлением выполнения функции processCategories и следовательно не вызывается функция close
				код
			defer func() {
				if arg := recover(); arg != nil {
					fmt.Println(arg)
					close(outChan)
				}
			}()
		        исправляет эту ошибку (добавлена после)

				Watersports Total: 328
				Chess Total: 1291
				Cannot find category: Running
			Но это опять же таки не лучший вариант так произходит без указания функции main
				исправим добавлением строк 181, 188, 189, и измением цикла for message  в функции main
			Watersports Total: 328
			Chess Total: 1291
			Cannot find category: Running
			A terminal error occured

	*/

}
