package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
// }

//Product: Kayak Price: 275

/*Написание строк*/

// func main() {
// 	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
// 	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
// }

/*
Product: Kayak Price: 275
Product:KayakPrice:275
*/

/*Форматирование строк*/
// func main() {
// 	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
// }
// Product: Kayak, Price: $275.00

// func getProductName(index int) (name string, err error) {
// 	if len(Products) > index {
// 		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
// 	} else {
// 		err = fmt.Errorf("error for index %v", index)
// 	}
// 	return
// }

// func main() {
// 	name, _ := getProductName(1)
// 	fmt.Println(name)
// 	_, err := getProductName(10)
// 	fmt.Println(err)
// }

/*
Name of product: Lifejacket
error for index 10
*/

/*Понимание глаголов форматирования*/

/*
Использовнаие глаголов форматирования общего назначения

%v - отображает формат значения по умолчанию
%+v - включает имена полей при записи значений структуры
%#v - можно использовать для повторного создания значения в файле кода
%T - отображает тип значения
*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	Printfln("Value: %v", Kayak)
// 	Printfln("Go syntax: %#v", Kayak)
// 	Printfln("Type: %T", Kayak)
// }
/*
Value: {Kayak Watersports 275}
Go syntax: main.Product{Name:"Kayak", Category:"Watersports", Price:275}
Type: main.Product
*/

/*Управление форматирование структуры*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	Printfln("Value: %v", Kayak)
// 	Printfln("Value with fields: %+v", Kayak)
// }

/*
Value: {Kayak Watersports 275}
Value with fields: {Name:Kayak Category:Watersports Price:275}
*/

/*
После добавления:
func (p Product) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}
в product.go
Value: Product: Kayak, Price: $275.00
Value with fields: Product: Kayak, Price: $275.00
*/

/*Использование команд целочисленного формативрования*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	number :=250

// 	Printfln("Binary: %b", number)
// 	Printfln("Decimal: %d", number)
// 	Printfln("Octal: %o, %O", number, number)
// 	Printfln("Hexadecimal: %x, %X", number, number)
// }
/*
Binary: 11111010
Decimal: 250
Octal: 372, 0o372
Hexadecimal: fa, FA
*/

/*Использование глаголов форматирования значений с плавающей запятой*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	number := 279.00

// 	Printfln("Decimalless with exponent: %b", number)
// 	Printfln("Decimal with exponent: %e", number)
// 	Printfln("Decimal without exponent: %f", number)
// 	Printfln("Hexadecimal: %x, %X", number, number)

// 	// форматом можно уапрвлять указав ширину и точность
// 	Printfln("Decimal without exponent: >>%8.2f<<", number)
// 	// ширина может быть опущена
// 	Printfln("Decimal without exponent: >>%.2f<<", number)

// 	/*
// 		Вывод можно изменить с помощью модификаторов
// 		"+" - указание знака (положительный или отрицательный),
// 		"0" - ширина заполняется не пустыми полями, а нолями,
// 		"-" - добавляет отступ справа, а не слева
// 	*/
// 	Printfln("Sign: >>%+.2f<<", number)
// 	Printfln("Zeros for Padding: >>%010.2f<<", number)
// 	Printfln("Right Padding: >>%.2f<<", number)
// }

/*
Decimalless with exponent: 4908219906392064p-44
Decimal with exponent: 2.790000e+02
Decimal without exponent: 279.000000
Hexadecimal: 0x1.17p+08, 0X1.17P+08

Decimal without exponent: >>  279.00<<

Decimal without exponent: >>279.00<<

Sign: >>+279.00<<
Zeros for Padding: >>0000279.00<<
Right Padding: >>279.00<<
*/

/*
использование глаголов форматирования строк и символов
%s - отображает строку (применяется по умолчанию при использовании %v)
%c - отображает характер
%U - отображает символы в формате Unicode
*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	name := "Kayak"

// 	Printfln("String: %s", name)
// 	Printfln("Character: %c", []rune(name)[0])
// 	Printfln("Unicode: %U", []rune(name)[0])
// }
/*
String: Kayak
Character: K
Unicode: U+004B
*/

/*Использвание глагола форматирования значений*/
// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	name := "Kayak"

// 	Printfln("Bool: %t", len(name)>1)
// 	Printfln("Bool: %t", len(name)>100)
// }

/*
Bool: true
Bool: false
*/

/*Использвание глагола форматирования указателя*/
// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	name := "Kayak"

// 	Printfln("Pointer: %p", &name)
// }

/*
Pointer: 0xc000010250
*/

/*Сканирование строк*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	var name string
// 	var category string
// 	var price float64

// 	fmt.Print("Enter text to scan: ")
// 	n, err := fmt.Scan(&name, &category, &price)

// 	if err==nil {
// 		Printfln("Scanned %v values", n)
// 		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*
Enter text to scan: Kayak Watersports 279
Scanned 3 values
Name: Kayak, Category: Watersports, Price: 279.00

Enter text to scan: Kayak Watersports Zero
Error: strconv.ParseFloat: parsing "": invalid syntax

Enter text to scan: Kayak Watersports
279
Scanned 3 values
Name: Kayak, Category: Watersports, Price: 279.00
*/

/*Работа с символами новой строки*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	var name string
// 	var category string
// 	var price float64

// 	fmt.Print("Enter text to scan: ")
// 	n, err := fmt.Scanln(&name, &category, &price)

// 	if err==nil {
// 		Printfln("Scanned %v values", n)
// 		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*
Enter text to scan: Kayak Watersports
Error: unexpected newline
*/

/*Использование другого источника строк*/

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	var name string
// 	var category string
// 	var price float64

// 	source := "Lifejacket Watersport 48.95"
// 	n, err := fmt.Sscan(source, &name, &category, &price)

// 	if err == nil {
// 		Printfln("Scanned %v values", n)
// 		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*
Scanned 3 values
Name: Lifejacket, Category: Watersport, Price: 48.95
*/

/*Использование шаблона сканирования*/

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	var name string
	var category string
	var price float64

	source := "Product Lifejacket Watersport 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)

	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

/*
Scanned 3 values
Name: Lifejacket, Category: Watersport, Price: 48.95
*/
