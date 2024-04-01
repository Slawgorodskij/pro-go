package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello, Operations")

	// price, tax := 275.00, 27.40
	// sum := price + tax
	// difference := price - tax
	// product := price * tax
	// quotient := price / tax

	// fmt.Println(sum)        // 302.4
	// fmt.Println(difference) // 247.6
	// fmt.Println(product)    // 7535
	// fmt.Println(quotient)   // 10.036496350364963

	// Арифмети́ческое переполне́ние — специфичная для компьютерной арифметики ситуация, когда при арифметическом действии результат становится больше максимально возможного значения для переменной, использующейся для хранения результата.

	// var inVal = math.MaxInt64                  //  MaxInt64 - константа представляющая наибольшее значение типа - int64
	// var floatVal = math.MaxFloat64             // MaFloat64 - константа представляющая наибольшее значение типа - float64
	// fmt.Println(inVal * 2)                     // -2
	// fmt.Println(floatVal * 2)                  // +Inf
	// fmt.Println(math.IsInf((floatVal * 2), 0)) // true IsInf - функция для обнаружения бесконечности

	// оператор "%" - возвращает остаток от деления целочисленного значения (в этом языке это не оператор по модулю так как возвращает и отрицательные значения)

	// posResult := 3 % 2
	// negResult := -3 % 2
	// absResult := math.Abs(float64(negResult)) // Abs функция возвращающая абсолютное значение float64

	// fmt.Println(posResult) // 1
	// fmt.Println(negResult) // -1
	// fmt.Println(absResult) // 1

	// операторы инкремента и декремента
	// value := 10.2
	// value++
	// fmt.Println(value) // 11.2
	// value += 2
	// fmt.Println(value) //13.2
	// value -= 2
	// fmt.Println(value) //11.2
	// value--
	// fmt.Println(value) //10.2

	//Объединение строк в языке GO возможно объединение только строк

	// greeting := "Hello"
	// language := "GO"
	// combinedString := greeting + "," + language

	// fmt.Println(combinedString) //Hello,GO

	// Операторы сравнение

	//	first := 100
	// const second = 200.00

	// equal := first == second
	// notEqual := first != second
	// lessThan := first < second
	// lessThanOrEqual := first <= second
	// greaterThan := first > second
	// greaterThanOrEqual := first >= second
	// fmt.Println(equal)              // false
	// fmt.Println(notEqual)           // true
	// fmt.Println(lessThan)           // true
	// fmt.Println(lessThanOrEqual)    // true
	// fmt.Println(greaterThan)        // false
	// fmt.Println(greaterThanOrEqual) // false

	// В GO нет тернарных сравнений

	//Сравнение указателей

	// first := 100
	// second := &first
	// third := &first

	// alpha := 100
	// beta := &alpha

	// fmt.Println(second == third)   // true  ссылается на одну ячейку
	// fmt.Println(second == beta)    // false указывает на другую...
	// fmt.Println(*second == *third) // true сравнение идет значения 100
	// fmt.Println(*second == *beta)  // true сравнение идет значения 100

	// Логические операторы
	// maxMph := 50
	// passengerCapacity := 4
	// airbags := true
	// familyCar := passengerCapacity > 2 && airbags
	// sportsCar := maxMph > 100 || passengerCapacity == 2
	// canCategorize := !familyCar && !sportsCar

	// fmt.Println(familyCar)     // true
	// fmt.Println(sportsCar)     // false
	// fmt.Println(canCategorize) // false

	// Преобразование, анализ и форматирование значений
	// Go не позволяет смешивать типы

	// kayak := 275
	// soccerBall := 19.50

	// total := kayak + soccerBall
	// fmt.Println(total) // invalid operation: kayak + soccerBall (mismatched types int and float64)

	// total := float64(kayak) + soccerBall
	// fmt.Println(total) // после явного преобразования значения переменной "kayak" код выполнится - 294.5

	// необхдимо помнит, что явное преобразование может привести к потере точности или переполнению

	// total := kayak + int(soccerBall)
	// fmt.Println(total) // код выполнится, но - 294
	// fmt.Println(int8(total)) // код выполнится, но - 38

	// totalCeil := kayak + int(math.Ceil(soccerBall))               //возвращает наименьшее целое число
	// totalFloor := kayak + int(math.Floor(soccerBall))             //возвращает наибольшее целое число
	// totalRound := kayak + int(math.Round(soccerBall))             //округляет до ближайшего целого числа
	// totalRoundToEven := kayak + int(math.RoundToEven(soccerBall)) //округляет до ближайшего четного целого числа
	// fmt.Println(totalCeil)                                        // код выполнится - 295
	// fmt.Println(totalFloor)                                       // код выполнится - 294
	// fmt.Println(totalRound)                                       // код выполнится - 295
	// fmt.Println(totalRoundToEven)                                 // код выполнится - 295

	// val1 := "true"
	// val2 := "false"
	// val3 := "not true"

	// bool1, b1err := strconv.ParseBool(val1)
	// bool2, b2err := strconv.ParseBool(val2)
	// bool3, b3err := strconv.ParseBool(val3)

	// fmt.Println("Bool 1", bool1, b1err) // Bool 1 true <nil>
	// fmt.Println("Bool 2", bool2, b2err) // Bool 2 false <nil>
	// fmt.Println("Bool 3", bool3, b3err) // Bool 3 false strconv.ParseBool: parsing "not true": invalid syntax

	// val1 := "0"
	// bool1, b1err := strconv.ParseBool(val1)
	// if b1err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // результат Parsed value: false

	// // второй вариант записи

	// if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
	// 	fmt.Println("Parsed value:", bool1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }

	// Разбор целых чисел

	// val1 := "100"

	// val1 := "500"

	// if int1, int1err := strconv.ParseInt(val1, 0, 8); int1err == nil {
	// 	fmt.Println("Parsed value:", int1)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }

	// ответ - Parsed value: 100
	// ответ - Cannot parse 500

	// val1 := "100"

	// if int1, int1err := strconv.ParseInt(val1, 0, 8); int1err == nil {
	// 	smallInt := int8(int1)
	// 	fmt.Println("Parsed value:", smallInt)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // ответ - Parsed value: 100

	// val1 := "100"

	// if int1, int1err := strconv.ParseInt(val1, 2, 8); int1err == nil {
	// 	smallInt := int8(int1)
	// 	fmt.Println("Parsed value:", smallInt)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // ответ - Parsed value: 4

	// val1 := "0b1100100"

	// if int1, int1err := strconv.ParseInt(val1, 0, 8); int1err == nil {
	// 	smallInt := int8(int1)
	// 	fmt.Println("Parsed value:", smallInt)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // ответ - Parsed value: 100

	// val1 := "100"

	// if int1, int1err := strconv.ParseInt(val1, 10, 0); int1err == nil {
	// 	var intResult int = int(int1)
	// 	fmt.Println("Parsed value:", intResult)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // ответ - Parsed value: 100

	// val1 := "100"

	// int1, int1err := strconv.Atoi(val1) //принимает только значение и не поддерживает анализ не десятичных значений
	// if int1err == nil {
	// 	var intResult int = int1
	// 	fmt.Println("Parsed value:", intResult)
	// } else {
	// 	fmt.Println("Cannot parse", val1)
	// }
	// // ответ - Parsed value: 100

	// val1 := "48.95"

	// float1, float1err := strconv.ParseFloat(val1, 64) //принимает только значение и не поддерживает анализ не десятичных значений
	// if float1err == nil {
	// 	fmt.Println("Parsed value:", float1)
	// } else {
	// 	fmt.Println("Cannot parse", val1, float1err)
	// }
	// // ответ - Parsed value: 48.95

	// val1 := "4.895e+01"

	// float1, float1err := strconv.ParseFloat(val1, 64) //принимает только значение и не поддерживает анализ не десятичных значений
	// if float1err == nil {
	// 	fmt.Println("Parsed value:", float1)
	// } else {
	// 	fmt.Println("Cannot parse", val1, float1err)
	// }
	// // ответ - Parsed value: 48.95

	//Форматирование значений как строк

	// val1 := true
	// val2 := false

	// str1 := strconv.FormatBool(val1)
	// str2 := strconv.FormatBool(val2)

	// fmt.Println("Formatted value 1: " + str1) // Formatted value 1: true
	// fmt.Println("Formatted value 2: " + str2) // Formatted value 2: false

	//val := 275

	// base10String := strconv.FormatInt(int64(val), 10)
	// base2String := strconv.FormatInt(int64(val), 2)

	// fmt.Println("Base 10: " + base10String) // Base 10: 275
	// fmt.Println("Base 2: " + base2String)   // Base 2: 100010011

	// base10String := strconv.Itoa(val)
	// base2String := strconv.FormatInt(int64(val), 2)

	// fmt.Println("Base 10: " + base10String) // Base 10: 275
	// fmt.Println("Base 2: " + base2String)   // Base 2: 100010011

	val := 49.95

	Fstring := strconv.FormatFloat(val, 'f', 2, 64) // val - обрабатываемое значение; 'f'- byte значение указывающее формат строки; 2 - указывает количество цифр после запятой
	Estring := strconv.FormatFloat(val, 'e', -1, 64)

	fmt.Println("Format F: " + Fstring) // 49.95
	fmt.Println("Format E: " + Estring) // 4.995e+01
}
