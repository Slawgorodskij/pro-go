package main

import (
	"fmt"
	"reflect"
	"strings"
)

//func printDetails(values ...Product) {
//	for _, elem := range values {
//		Printfln("Product: Name: %v, Category: %v, Price: %v", elem.Name, elem.Category, elem.Price)
//	}
//}

//для решения проблемы получения ошибки
/*
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v", val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}

	//printDetails(product)
	// Product: Name: Kayak, Category: Watersports, Price: 279

	customer := Customer{Name: "Alice", City: "New York"}
	printDetails(product, customer)
	//  cannot use customer (variable of type Customer) as Product value in argument to printDetails

	//после изменения функции printDetails
	// Product: Name: Kayak, Category: Watersports, Price: 279
	// Customer: Name: Alice, City: New York
}
*/
// У выше описанного проблема все равно останется так как функции printDetails придется постоянно расширять для исключения
//в Go предусмотренна "Рефлексия" - которая позволяет приложению работать с типами, неизвестными при компиляции проекта

// Использование рефлексии
/*
func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}
	printDetails(product, customer, payment, 10, true)
}
*/
// Использование основных функций типа

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}
	printDetails(product, customer, payment, 10, true)
}
