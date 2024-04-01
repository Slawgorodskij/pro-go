package main

import "fmt"

//func main() {
//fmt.Println("Hello, Structs")
//type Product struct {
//	name, category string
//	price          float64
//}
//kayak := Product{
//	name:     "Kayak",
//	category: "Watersports",
//	price:    275,
//}
//fmt.Println(kayak.name, kayak.category, kayak.price)
//kayak.price = 300
//fmt.Println("Changed price:", kayak.price)
//
//kayakNew := Product{
//	name:     "Kayak NEW",
//	category: "Watersports",
//}
//fmt.Println(kayakNew.name, kayakNew.category, kayakNew.price)
//kayakNew.price = 350
//fmt.Println("Changed price:", kayakNew.price)
//
//var lifejacket Product
//fmt.Println("Name is zero value", lifejacket.name == "")
//fmt.Println("Category is zero value", lifejacket.category == "")
//fmt.Println("Price is zero value", lifejacket.price == 0)

// Использование имен не обязательно

//var kayak = Product{"Kayak", "Watersports", 275.00}
//fmt.Println("Name:", kayak.name)
//fmt.Println("Category", kayak.category)
//fmt.Println("Price", kayak.price)

//Определение встроенных полей

//type StockLevel struct {
//	Product
//	count int
//}
//stockItem := StockLevel{
//	Product: Product{"Kayak", "Watersports", 275.00},
//	count:   100,
//}
//fmt.Println("Name", stockItem.Product.name)
//fmt.Println("Count", stockItem.count)
//fmt.Println(fmt.Sprint("name:", stockItem.Product.name))

//Имена полей должнф быть уникальны. При необходимости создания двух полей одного типа:

//type StockLevel struct {
//	Product
//	Alternate Product
//	count     int
//}
//stockItem := StockLevel{
//	Product:   Product{"Kayak", "Watersports", 275.00},
//	Alternate: Product{"Lifejacket", "Watersports", 48.95},
//	count:     100,
//}
//fmt.Println("Name", stockItem.Product.name)
//fmt.Println("Alt Name", stockItem.Alternate.name)

//Сравнение значений структуры

//p1 := Product{"Kayak", "Watersports", 275.00}
//p2 := Product{"Kayak", "Watersports", 275.00}
//p3 := Product{"Kayak", "Boats", 275.00}
//
//fmt.Println("p1 == p2", p1 == p2)
//fmt.Println("p1 == p3", p1 == p3)

//Нельзя сравнивать, если тип структуры определяет поля с несравнимыми типами!

//type Product struct {
//	name, category string
//	price          float64
//	otherNames     []string
//}

//p1 := Product{"Kayak", "Watersports", 275.00}
//p2 := Product{"Kayak", "Watersports", 275.00}
//p3 := Product{"Kayak", "Boats", 275.00}
//
//fmt.Println("p1 == p2", p1 == p2) //invalid operation: p1 == p2 (struct containing []string cannot be compared)
//fmt.Println("p1 == p3", p1 == p3) // invalid operation: p1 == p3 (struct containing []string cannot be compared)

//Преобразования между типами структур

//type Product struct {
//	name, category string
//	price          float64
//}
//type Item struct {
//	name     string
//	category string
//	price    float64
//}
//prod := Product{
//	name:     "Kayak",
//	category: "Watersports",
//	price:    275,
//}
//item := Item{
//	name:     "Kayak",
//	category: "Watersports",
//	price:    275,
//}
//fmt.Println("prod == item", prod == Product(item)) // prod == item true
//}

//Определение аннонимных типов структур

//func writeName(val struct {
//	name, category string
//	price          float64
//}) {
//	fmt.Println("Name:", val.name)
//}
//func main() {
//
//	type Product struct {
//		name, category string
//		price          float64
//	}
//	type Item struct {
//		name     string
//		category string
//		price    float64
//	}
//	prod := Product{
//		name:     "Kayak",
//		category: "Watersports",
//		price:    275,
//	}
//	item := Item{
//		name:     "Kayak",
//		category: "Watersports",
//		price:    275,
//	}
//	writeName(prod)
//	writeName(item)
//}

//func main() {
//	type Product struct {
//		name, category string
//		price          float64
//	}
//	prod := Product{
//		name:     "Kayak",
//		category: "Watersports",
//		price:    275,
//	}
//
//	var builder strings.Builder
//	json.NewEncoder(&builder).Encode(struct {
//		ProductName  string
//		ProductPrice float64
//	}{
//		ProductName:  prod.name,
//		ProductPrice: prod.price,
//	})
//
//	fmt.Println(builder.String())
//}

//Создание массивов, срезов и карт, содержащих структурные значения

//func main() {
//	type Product struct {
//		name, category string
//		price          float64
//	}
//	type StockLevel struct {
//		Product
//		Alternate Product
//		count     int
//	}
//	array := [1]StockLevel{
//		{
//			Product: Product{
//				name:     "Kayak",
//				category: "Watersports",
//				price:    275,
//			},
//			Alternate: Product{
//				name:     "Lifejacket",
//				category: "Watersports",
//				price:    48.95,
//			},
//			count: 100,
//		},
//	}
//	fmt.Println("Array:", array[0].Product.name)
//
//	slice := []StockLevel{
//		{
//			Product: Product{
//				name:     "Kayak",
//				category: "Watersports",
//				price:    275,
//			},
//			Alternate: Product{
//				name:     "Lifejacket",
//				category: "Watersports",
//				price:    48.95,
//			},
//			count: 100,
//		},
//	}
//	fmt.Println("Slice:", slice[0].Product.name)
//
//	kvp := map[string]StockLevel{
//		"kayak": {
//			Product: Product{
//				name:     "Kayak",
//				category: "Watersports",
//				price:    275,
//			},
//			Alternate: Product{
//				name:     "Lifejacket",
//				category: "Watersports",
//				price:    48.95,
//			},
//			count: 100,
//		},
//	}
//	fmt.Println("Map:", kvp["kayak"].Product.name)
//}

//Понимание СТРУКТУР и УКАЗАТЕЛЕЙ

//func main() {
//	type Product struct {
//		name, category string
//		price          float64
//	}
//
//	p1 := Product{
//		name:     "Kayak",
//		category: "Watersports",
//		price:    275,
//	}
//	//p2 := p1
//
//	//p1.name = "Original KAYAK"
//	//fmt.Println("P1:", p1.name) // Original KAYAK
//	//fmt.Println("P2", p2.name) // Kayak
//
//	p3 := &p1
//	p1.name = "Original KAYAK"
//	fmt.Println("P1:", p1.name)   // Original KAYAK
//	fmt.Println("P3", (*p3).name) // Original KAYAK
//
//}

//Понимание удобного синтаксиса указателя структуры

//type Product struct {
//	name, category string
//	price          float64
//}

//Вместо
//func calcTax(product *Product) {
//	if (*product).price > 100 {
//		(*product).price += (*product).price * 0.2
//	}
//}

//Можно записать:
//func calcTax(product *Product) {
//	if product.price > 100 {
//		product.price += product.price * 0.2
//	}
//}

//Финальный вариант может выглядеть:
//func calcTax(product *Product) *Product {
//	if product.price > 100 {
//		product.price += product.price * 0.2
//	}
//	return product
//}

//func main() {
//Вместо
//kayak := Product{
//	name:     "Kayak",
//	category: "Watersports",
//	price:    275,
//}
//calcTax(&kayak)

//Можно записать:
//kayak := &Product{
//	name:     "Kayak",
//	category: "Watersports",
//	price:    275,
//}
//calcTax(kayak)

//Следовательно запись вызова будет выглядеть:

//	kayak := calcTax(&Product{
//		name:     "Kayak",
//		category: "Watersports",
//		price:    275,
//	})
//
//	fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
//}

// Понимание функций контсруктора структуры

//type Product struct {
//	name, category string
//	price          float64
//}
//
//func newProduct(name, category string, price float64) *Product {
//	return &Product{name, category, price}
//}
//
//func main() {
//	products := [2]*Product{
//		newProduct("Kayak", "Watersport", 275),
//		newProduct("Hat", "Skiing", 42.50),
//	}
//
//	for _, p := range products {
//		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
//	}
//}

//Использование типов указателей для полей структуры

//type Product struct {
//	name, category string
//	price          float64
//	*Supplier
//}

//type Supplier struct {
//	name, city string
//}

//func newProduct(name, category string, price float64, supplier *Supplier) *Product {
//	return &Product{name, category, price - 10, supplier}
//}

//Для глубоко копирования (не предусмотрен языком GO)
//func copyProduct(product *Product) Product {
//	p := *product
//	s := *product.Supplier
//	p.Supplier = &s
//	return p
//}
//func main() {
//	acme := &Supplier{"Acme Co", "New York"}
//products := [2]*Product{
//	newProduct("Kayak", "Watersport", 275, acme),
//	newProduct("Hat", "Skiing", 42.50, acme),
//}
//for _, p := range products {
//	fmt.Println("Name:", p.name, "Price:", p.price, "Supplier:", p.Supplier.name, p.Supplier.city)
//}
//Копирование поля указателя
//p1 := newProduct("Kayak", "Watersport", 275, acme)
//p2 := *p1
//
//p1.name = "Original KAYAK"
//p1.Supplier.name = "BoatCo"
//
//for _, p := range []Product{*p1, p2} {
//	fmt.Println("Name:", p.name, "Price:", p.price, "Supplier:", p.Supplier.name, p.Supplier.city)
//}
/*
	результатом будет
	Name: Original KAYAK Price: 265 Supplier: BoatCo New York
	Name: Kayak Price: 265 Supplier: BoatCo New York
	так как "продукт" с копируется, но поставщик будет ссылаться на одну и туже ящейку памяти
*/

//При использовании созданой функции копировавния
//p1 := newProduct("Kayak", "Watersport", 275, acme)
//p2 := copyProduct(p1)
//
//p1.name = "Original KAYAK"
//p1.Supplier.name = "BoatCo"
//
//for _, p := range []Product{*p1, p2} {
//	fmt.Println("Name:", p.name, "Price:", p.price, "Supplier:", p.Supplier.name, p.Supplier.city)
//}
/*
	результатом будет
	Name: Original KAYAK Price: 265 Supplier: BoatCo New York
	Name: Kayak Price: 265 Supplier: Acme Co New York
*/
//}

//Понимание нулевого значения для структур и указателей на структуры

//type Product struct {
//	name, category string
//	price          float64
//}
//
//func main() {
//	var prod Product
//	var prodPtr *Product
//	fmt.Println("Value:", prod.name, prod.category, prod.price)
//	fmt.Println("Pointer:", prodPtr)
//
//	/*
//	результатом будет
//	Value:   0
//	Pointer: <nil>
//	*/
//}

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func main() {
	//var prod Product
	//var prodPtr *Product
	//fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	//fmt.Println("Pointer:", prodPtr)

	/*
		результатом будет
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47e028]
		Избежать ошибки позволит:
	*/

	var prod Product = Product{Supplier: &Supplier{}}
	var prodPtr *Product
	fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	fmt.Println("Pointer:", prodPtr)
	/*
		результатом будет
		Value:   0
		Pointer: <nil>
	*/
}
