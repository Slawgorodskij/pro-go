package main

import "fmt"

func main() {
	//fmt.Println("Hello, Collections")

	// var names [3]string
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"

	//names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	//fmt.Println(names) //[Kayak Lifejacket Paddle]

	// names := [3]string{"Kayak", "Lifejacket"}
	// fmt.Println(names) //[Kayak Lifejacket ]

	// Масивы Go одномерны, но их можно комбинировать для получения многомерных

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// var otherArray [4]string = names
	// fmt.Println(otherArray)// cannot use names (variable of type [3]string) as type [4]string in variable declaration

	// names := [...]string{"Kayak", "Lifejacket", "Paddle"}
	// var otherArray [4]string = names
	// fmt.Println(otherArray) // cannot use names (variable of type [3]string) as type [4]string in variable declaration

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// otherArray := names
	// names[0] = "Canoe"
	// fmt.Println("names:", names)           //names: [Canoe Lifejacket Paddle]
	// fmt.Println("otherArray:", otherArray) //otherArray: [Kayak Lifejacket Paddle]

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// otherArray := &names
	// names[0] = "Canoe"
	// fmt.Println("names:", names)            //names: [Canoe Lifejacket Paddle]
	// fmt.Println("otherArray:", *otherArray) //otherArray: [Canoe Lifejacket Paddle]

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// same := names == moreNames
	// fmt.Println("comparison (сравнение)", same) //массивы равны - comparison (сравнение) true

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}

	// for index, value := range names {
	// 	fmt.Println("Index", index, "Value", value)
	// }
	// ОТВЕТ:
	// 		Index 0 Value Kayak
	// 		Index 1 Value Lifejacket
	// 		Index 2 Value Paddle

	// names := [3]string{"Kayak", "Lifejacket", "Paddle"}

	// for _, value := range names {
	// 	fmt.Println("Value", value)
	// }
	// ОТВЕТ:
	// 		Value Kayak
	// 		Value Lifejacket
	// 		Value Paddle

	//СРЕЗЫ

	// names := make([]string, 3) // make - функция; []string - тип; 3 - длина
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"

	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names) // [Kayak Lifejacket Paddle]

	//ДОБАВЛЕНИЕ ЭЛЕМЕНТОВ В СРЕЗ

	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// names = append(names, "Hat", "Gloves") // append - создает новый массив

	// fmt.Println(names) // [Kayak Lifejacket Paddle Hat Gloves]

	// names := []string{"Kayak", "Lifejacket", "Paddle"}

	// appendedNames := append(names, "Hat", "Gloves")
	// names[0] = "Canoe"

	// fmt.Println(names)         // [Canoe Lifejacket Paddle]
	// fmt.Println(appendedNames) // [Kayak Lifejacket Paddle Hat Gloves]

	//ДОПОЛНИТЕЛЬНАЯ ЕМКОСТЬ ДДЛЯ СРЕЗОВ

	// names := make([]string, 3, 6) // make - функция; []string - тип; 3 - длина; 6 - емкость
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println("len:", len(names)) // 3
	// fmt.Println("cap:", cap(names)) // 6

	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// appendedNames := append(names, "Hat", "Gloves") // names - имеет нулевые длину и емкость
	// names[0] = "Canoe"
	// fmt.Println(names)         // [Canoe Lifejacket Paddle]
	// fmt.Println(appendedNames) // [Canoe Lifejacket Paddle Hat Gloves]

	// ДОБАВЛЕНИЕ ОДНОГО СРЕЗА К ДРУГОМУ

	// names := make([]string, 3, 6)
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// moreNames := []string{"Hat Gloves"}
	// appendedNames := append(names, moreNames...)

	// fmt.Println(names)         // [Kayak Lifejacket Paddle]
	// fmt.Println(moreNames)     // [Hat Gloves]
	// fmt.Println(appendedNames) // [Kayak Lifejacket Paddle Hat Gloves]

	//СОЗДАНИЕ СРЕЗОВ ИЗ МАССИВОВ

	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// someNames := products[1:3] // присвоит элементы с первого по третий (нумерация по индексам)
	// allNames := products[:]

	// fmt.Println("someNames", someNames)                                   // someNames [Lifejacket Paddle]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // omeNames len: 2 cap: 3
	// fmt.Println("allNames", allNames)                                     // allNames [Kayak Lifejacket Paddle Hat]
	// fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))    // allNames len: 4 cap: 4

	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// someNames := products[1:3] // присвоит элементы с первого по третий (нумерация по индексам)
	// allNames := products[:]

	// someNames = append(someNames, "Gloves")

	// fmt.Println("someNames", someNames)                                   // someNames [Lifejacket Paddle Gloves]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // omeNames len: 3 cap: 3
	// fmt.Println("allNames", allNames)                                     // allNames [Kayak Lifejacket Paddle Gloves]
	// fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))    // allNames len: 4 cap: 4

	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// someNames := products[1:3]
	// allNames := products[:]

	// someNames = append(someNames, "Gloves")
	// someNames = append(someNames, "Boots")

	// fmt.Println("someNames", someNames)                                   // someNames [Lifejacket Paddle Gloves Boots]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // omeNames len: 4 cap: 6
	// fmt.Println("allNames", allNames)                                     // allNames [Kayak Lifejacket Paddle Gloves]
	// fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))    // allNames len: 4 cap: 4

	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// someNames := products[1:3:3] // 1 - низ; 3 - верх; 3 - максимум
	// allNames := products[:]

	// fmt.Println("someNames", someNames)                                   // someNames [Lifejacket Paddle]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // omeNames len: 2 cap: 2

	// someNames = append(someNames, "Gloves") // емкость ограничина была до двух вследствие чего создается новый массив

	// fmt.Println("someNames", someNames)                                   // someNames [Lifejacket Paddle Gloves]
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // omeNames len: 3 cap: 4
	// fmt.Println("allNames", allNames)                                     // allNames [Kayak Lifejacket Paddle Hat]
	// fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))    // allNames len: 4 cap: 4

	//СОЗДАНИЕ СРЕЗОВ ИЗ ДРУГИХ СРЕЗОВ
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := allNames[1:3]
	// allNames = append(allNames, "Gloves")
	// allNames[1] = "Canoe"
	// fmt.Println("someNames:", someNames) //someNames: [Paddle Hat]
	// fmt.Println("allNames:", allNames)   //allNames: [Lifejacket Canoe Hat Gloves]

	//Копирование для обеспечения разделения массива среза
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := make([]string, 2)
	// copy(someNames, allNames) // функция(получатель, источник)

	// fmt.Println("someNames:", someNames) //someNames: [Lifejacket Paddle]
	// fmt.Println("allNames:", allNames)   //allNames: [Lifejacket Paddle Hat]

	// //Если срез неинициализирован, то копирование выполнится без ошибок, но массив будет пустым
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// var someNames []string    //нулевая длина, нулевая емкость
	// copy(someNames, allNames) // скопируется ноль элементов

	// fmt.Println("someNames:", someNames) //someNames: []
	// fmt.Println("allNames:", allNames)   //allNames: [Lifejacket Paddle Hat]

	//УКАЗАНИЕ ДИАПАЗОНОВ ПРИ КОПИРОВАНИИ
	//УКАЗАНИЕ ДИАПАЗОНОВ ПРИ КОПИРОВАНИИ
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := []string{"Boots", "Canoe"}
	// copy(someNames[1:], allNames[2:3]) //

	// fmt.Println("someNames:", someNames) //someNames: [Boots Hat]
	// fmt.Println("allNames:", allNames)   //allNames: [Lifejacket Paddle Hat]

	//КОПИРОВАНИЕ СРЕЗОВ РАЗНЫХ РАЗМЕРОВ
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// replacementProducts := []string{"Canoe", "Boots"}

	// copy(products, replacementProducts)
	// fmt.Println("products:", products) //products: [Canoe Boots Paddle Hat]

	// copy(products[0:1], replacementProducts)
	// fmt.Println("products:", products) //products: [Canoe Lifejacket Paddle Hat]

	//УДАЛЕНИЕ ЭЛЕМЕНТОВ СРЕЗОВ
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// deleted := append(products[:2], products[3:]...)

	// fmt.Println("deleted:", deleted) //deleted: [Kayak Lifejacket Hat]

	//ПЕРЕЧИСЛЕНИЕ СРЕЗОВ
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// for index, value := range products[2:] {
	// 	fmt.Println("Index:", index, "Value", value)
	// }
	//ОТВЕТ:
	// 		Index: 0 Value Paddle
	// 		Index: 1 Value Hat

	//СОРТИРОВКА СРЕЗОВ
	// products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// sort.Strings(products)

	// for index, value := range products {
	// 	fmt.Println("Index:", index, "Value", value)
	// }

	//ОТВЕТ:
	// 		Index: 0 Value Hat
	// 		Index: 1 Value Kayak
	// 		Index: 2 Value Lifejacket
	// 		Index: 3 Value Paddle

	//СРАВНЕНИЕ СРЕЗОВ
	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// p2 := p1

	//fmt.Println("Equale:", p1 == p2) //invalid operation: p1 == p2 (slice can only be compared to nil)
	//fmt.Println("Equale:", reflect.DeepEqual(p1, p2)) //Equale: true

	//ПОЛУЧЕНИЕ МАССИВА, ЛЕЖАЩЕГО В ОСНОВЕ СРЕЗА
	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// arrayPtr := (*[3]string)(p1)
	// array := *arrayPtr

	// fmt.Println(array) //[Kayak Lifejacket Paddle]

	//РАБОТА С КАРТАМИ
	// products := make(map[string]float64, 10) // make - ключевое слово; map - ключевое слово; [string] - тип ключа; float64 - тип переменной; 10 - начальный размер)
	// products["Kayak"] = 279
	// products["Lifejacket"] = 48.95

	// fmt.Println("Map size:", len(products))  // Map size: 2
	// fmt.Println("price:", products["Kayak"]) //price: 279
	// fmt.Println("price:", products["Hat"])   //price: 0

	//ЛИТЕРАЛЬНЫЙ СИНТАКСИС КАРТЫ
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// }

	// fmt.Println("Map size:", len(products))  // Map size: 2
	// fmt.Println("price:", products["Kayak"]) //price: 279
	// fmt.Println("price:", products["Hat"])   //price: 0

	//ПРОВЕРКА ЭЛЕМЕНТОВ В КАРТЕ
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }

	// fmt.Println("price:", products["Hat"]) //price: 0 - при отсутствии элемента так же будет ноль
	// value, ok := products["Hat"]
	// if ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	//}

	// ОТВЕТ: Stored value: 0

	//УДАЛЕНИЕ ОБЪЕКТОВ С КАРТЫ
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }

	// delete(products, "Hat") // аргументы функции delete карта и ключ для удаления

	// if value, ok := products["Hat"]; ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// ОТВЕТ: No stored value

	//ПЕРЕЧИСЛЕНИЕ СОДЕРЖИМОГО КАРТЫ
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	//	}
	// for key, value := range products {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	// ОТВЕТ:
	// 		Key: Kayak Value: 279
	// 		Key: Lifejacket Value: 48.95
	// 		Key: Hat Value: 0

	//ПЕРЕЧИСЛЕНИЕ КАРТЫ ПО ПОРЯДКУ
	// products := map[string]float64{
	// 	"Kayak":      279,
	// 	"Lifejacket": 48.95,
	// 	"Hat":        0,
	// }
	// keys := make([]string, 0, len(products))
	// for key, _ := range products {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	// for _, key := range keys {
	// 	fmt.Println("Key:", key, "Value:", products[key])
	// }

	// ОТВЕТ:
	// 		Key: Hat Value: 0
	// 		Key: Kayak Value: 279
	// 		Key: Lifejacket Value: 48.95

	// ПОНИМАНИЕ ДВОЙНОЙ ПРИРОДЫ СТРОК
	// var price string = "$48.95"
	// var currency byte = price[0]
	// var amountString string = price[1:]
	// amount, parseErr := strconv.ParseFloat(amountString, 64)

	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse Error:", parseErr)
	// }
	// ОТВЕТ:
	// 		Currency: 36
	// 		Amount: 48.95

	// var price string = "$48.95"
	// var currency string = string(price[0])
	// var amountString string = price[1:]
	// amount, parseErr := strconv.ParseFloat(amountString, 64)

	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse Error:", parseErr)
	// }
	// ОТВЕТ:
	// 		Currency: $
	// 		Amount: 48.95

	// var price string = "€48.95"
	// var currency string = string(price[0])
	// var amountString string = price[1:]
	// amount, parseErr := strconv.ParseFloat(amountString, 64)

	// fmt.Println("Length:", len(price))
	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse Error:", parseErr)
	// }
	// ОТВЕТ:
	//		Length: 8
	// 		Currency: â
	// 		Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax

	// ПРЕОБРАЗОВАНИЕ СТРОКИ В РУНЫ
	// var price []rune = []rune("€48.95")
	// var currency string = string(price[0])
	// var amountString string = string(price[1:])
	// amount, parseErr := strconv.ParseFloat(amountString, 64)

	// fmt.Println("Length:", len(price))
	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse Error:", parseErr)
	// }
	// ОТВЕТ:
	//		Length: 6
	// 		Currency: €
	// 		Amount: 48.95

	// ПУРУЧИСЛЕНИЕ СТРОК
	var price = "€48.95"

	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
	// ОТВЕТ:
	// 		0 8364 €
	// 		3 52 4
	// 		4 56 8
	// 		5 46 .
	//		6 57 9
	// 		7 53 5

	for index, char := range []byte(price) {
		fmt.Println(index, char)
	}
	// ОТВЕТ:
	// 		0 226
	// 		1 130
	// 		2 172
	// 		3 52
	// 		4 56
	// 		5 46
	// 		6 57
	// 		7 53
}
