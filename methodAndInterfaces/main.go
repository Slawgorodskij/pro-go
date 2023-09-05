package main

import "fmt"

//type Product struct {
//	name, category string
//	price          float64
//}

//func printDetails(product *Product) {
//	fmt.Println("Name", product.name, "Category", product.category, "Price", product.price)
//}
//func main() {
//	products := []*Product{
//		{"Kayak", "Watersport", 275},
//		{"Lifejacket", "Watersport", 48.95},
//		{"Soccer Ball", "Soccer", 19.50},
//	}
//	//I этап
//	//for _, p := range products {
//	//	fmt.Println("Name", p.name, "Category", p.category, "Price", p.price)
//	//}
//
//	//II этап
//
//	for _, p := range products {
//		printDetails(p)
//	}
//
//}

//Определение метода

//func newProduct(name, category string, price float64) *Product {
//	return &Product{name, category, price}
//}
//func (product *Product) printDetails() {
//	fmt.Println("Name", product.name, "Category", product.category, "Price", product.calcTax(0.2, 100))
//}
//func (product *Product) calcTax(rate, threshold float64) float64 {
//	if product.price > threshold {
//		return product.price + (product.price * rate)
//	}
//	return product.price
//}
//func main() {
//	products := []*Product{
//		newProduct("Kayak", "Watersport", 275),
//		newProduct("Lifejacket", "Watersport", 48.95),
//		newProduct("Soccer Ball", "Soccer", 19.50),
//	}
//
//	for _, p := range products {
//		p.printDetails()
//	}
//}

//Перегрузка метода

//type Product struct {
//	name, category string
//	price          float64
//}
//type Supplier struct {
//	name, city string
//}
//
//func newProduct(name, category string, price float64) *Product {
//	return &Product{name, category, price}
//}
//func (product *Product) printDetails() {
//	fmt.Println("Name", product.name, "Category", product.category, "Price", product.calcTax(0.2, 100))
//}
//func (product *Product) calcTax(rate, threshold float64) float64 {
//	if product.price > threshold {
//		return product.price + (product.price * rate)
//	}
//	return product.price
//}
//func (supplier *Supplier) printDetails() {
//	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
//}
//
///*
//В коде у двух методов одиноковое имя (printDetails) и это не вызывет ошибки так как они имеют разных получателей.
//Следующий код уже вызовет ошибку:
//		./main.go:88:27: Supplier.printDetails redeclared in this block
//        ./main.go:80:27: other declaration of printDetails
//*/
//func (supplier *Supplier) printDetails(showName bool) {
//	if showName {
//		fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
//	} else {
//		fmt.Println("Supplier:", supplier.name)
//	}
//
//}
//
//func main() {
//	products := []*Product{
//		newProduct("Kayak", "Watersport", 275),
//		newProduct("Lifejacket", "Watersport", 48.95),
//		newProduct("Soccer Ball", "Soccer", 19.50),
//	}
//	supplier := []*Supplier{
//		{"Acme Co", "New York"},
//		{"BoatCo", "Chicago"},
//	}
//
//	for _, p := range products {
//		p.printDetails()
//	}
//	for _, s := range supplier {
//		s.printDetails()
//	}
//}

// Понимание ПОЛУЧАТЕЛЕЙ, УКАЗАТЕЛЕЙ и ЗНАЧЕНИЙ
//
//type Product struct {
//	name, category string
//	price          float64
//}
//
//func (product *Product) printDetails() {
//	fmt.Println("Name", product.name, "Category", product.category, "Price", product.calcTax(0.2, 100))
//}
//
//func (product *Product) calcTax(rate, threshold float64) float64 {
//	if product.price > threshold {
//		return product.price + (product.price * rate)
//	}
//	return product.price
//}
//
//func main() {
//	kayak := Product{"Kayak", "Watersport", 275}
//	kayak.printDetails()
//}
/*
функция printDetails может быть записана
func (product Product) printDetails() {
	fmt.Println("Name", product.name, "Category", product.category, "Price", product.calcTax(0.2, 100))
}
присваивание переменной в этом случае должна выглядеть:
kayak := &Product{"Kayak", "Watersport", 275}
*/

// Определение методов для псевдонимов типов

//type Product struct {
//	name, category string
//	price          float64
//}
//
//type ProductList []Product
//
//func (products *ProductList) calcCategoryTotals() map[string]float64 {
//	totals := make(map[string]float64)
//	for _, p := range *products {
//		totals[p.category] = totals[p.category] + p.price
//	}
//	return totals
//}
//
//func getProducts() []Product {
//	return []Product{
//		{"Kayak", "Watersport", 275},
//		{"Lifejacket", "Watersport", 48.95},
//		{"Soccer Ball", "Soccer", 19.50},
//	}
//}
//func main() {
//	//products := ProductList{
//	//	{"Kayak", "Watersport", 275},
//	//	{"Lifejacket", "Watersport", 48.95},
//	//	{"Soccer Ball", "Soccer", 19.50},
//	//}
//	products := ProductList(getProducts())
//	for category, total := range products.calcCategoryTotals() {
//		fmt.Println("Category:", category, "Total:", total)
//	}
//}

//Размещение типов и методов в отдельных файлах

/*
созданы два файла product.go и service.go
*/

//func main() {
//	kayak := Product{"Kayak", "Watersport", 275}
//	insurance := Service{"Boat Cover", 12, 89.50}
//
//	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
//	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))
//}

// Определение и использование ИНТЕРФЕЙСОВ

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//
//func main() {
//
//	expenses := []Expense{
//		Product{"Kayak", "Watersport", 275},
//		Service{"Boat Cover", 12, 89.50},
//	}
//
//	for _, expense := range expenses {
//		fmt.Println("expense:", expense.getName(), "Cost:", expense.getCost(true))
//	}
//}

//Использование ИНТЕРФЕЙСА в функции

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//
//func calcTotal(expenses []Expense) (total float64) {
//	for _, item := range expenses {
//		total += item.getCost(true)
//	}
//	return
//}
//func main() {
//
//	expenses := []Expense{
//		Product{"Kayak", "Watersport", 275},
//		Service{"Boat Cover", 12, 89.50},
//	}
//
//	for _, expense := range expenses {
//		fmt.Println("expense:", expense.getName(), "Cost:", expense.getCost(true))
//	}
//	fmt.Println("Total:", calcTotal(expenses))
//}

//Использование ИНТЕРФЕЙСА для полей стуктуры

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//type Account struct {
//	accountNumber int
//	expenses      []Expense
//}
//
//func calcTotal(expenses []Expense) (total float64) {
//	for _, item := range expenses {
//		total += item.getCost(true)
//	}
//	return
//}
//func main() {
//	account := Account{
//		accountNumber: 12345,
//		expenses: []Expense{
//			Product{"Kayak", "Watersport", 275},
//			Service{"Boat Cover", 12, 89.50},
//		},
//	}
//
//	for _, expense := range account.expenses {
//		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
//	}
//	fmt.Println("Total:", calcTotal(account.expenses))
//}

//понимание эффекта приемников метода указателя

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//
//func main() {
//	product := Product{"Kayak", "Watersport", 275}
//	/*
//			при var expense Expense = product
//			ответ:
//		Product field value: 100
//		Expense method result: 275
//
//			при var expense Expense = &product
//			ответ:
//		Product field value: 100
//		Expense method result: 100
//	*/
//	var expense Expense = product
//	product.price = 100
//
//	fmt.Println("Product field value:", product.price)
//	fmt.Println("Expense method result:", expense.getCost(false))
//}

//Сравнение знфчений ИНТЕРФЕЙСА

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//
//func main() {
//	var e1 Expense = &Product{name: "Kayak"}
//	var e2 Expense = &Product{name: "Kayak"}
//	var e3 Expense = Service{description: "Boat Cover"}
//	var e4 Expense = Service{description: "Boat Cover"}
//
//	fmt.Println("e1 == e2:", e1 == e2)
//	fmt.Println("e3 == e4:", e3 == e4)
//}

// Выполнение утверждений типа

//type Expense interface {
//	getName() string
//	getCost(annual bool) float64
//}
//
//func main() {
//	expenses := []Expense{
//		Service{"Boat Cover", 12, 89.50, []string{}},
//		Service{"Paddle Protect", 12, 8, []string{}},
//	}
//
//	for _, expense := range expenses {
//		s := expense.(Service)
//		fmt.Println("Service:", s.description, "Price", s.monthlyFee*float64(s.durationMonths))
//	}
//}

//Тестирование перед выполнеием утверждения типа

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
	}

	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price", s.monthlyFee*float64(s.durationMonths))
	}
}
