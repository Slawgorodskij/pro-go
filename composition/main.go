package main

import (
	"composition/store"
	"fmt"
)

func main() {
	//fmt.Println("Hello, Composition")
	//kayak := store.NewProduct("Kayak", "Watersports", 275)
	//lifejacket := &store.Product{Name: "lifejacket", Category: "Watersports"}
	//for _, p := range []*store.Product{kayak, lifejacket} {
	//	fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	//}

	//boats := []*store.Boat{
	//	store.NewBoat("Kayak", 275, 1, false),
	//	store.NewBoat("Canoe", 400, 3, false),
	//	store.NewBoat("Tender", 650.25, 2, true),
	//}
	/*
	   При использовании литерального синтаксиса присвоение значения будет выглядеть более громозким т.к.
	   boat := store.Boat {Name: "Kayak", Category: "Watersport", Capacity: 1, Motorized: false}
	   выдаст ошибку о неизвестном поле следовательно присвоение будет выглядеть:
	   boat := store.Boat{ Product: &store.Product{Name: "Kayak", Category: "Watersport"}, Capacity: 1, Motorized: false}
	*/

	//for _, b := range boats {
	//	fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name, "Price:", b.Price(0.2))
	//}
	//
	//rentals := []*store.RentalBoat{
	//	store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
	//	store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
	//	store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	//}
	//for _, r := range rentals {
	//	fmt.Println("Rental Boats:", r.Name, "Rental Price:", r.Price(0.2), "Capitan:", r.Capitan)
	//
	//}

	//product := store.NewProduct("Kayak", "Watersports", 279)
	//deal := store.NewSpecialDeal("Weekend Special", product, 50)
	//Name, price, Price := deal.GetDetails()
	//
	//fmt.Println("Name:", Name)
	//fmt.Println("Price field:", price)
	//fmt.Println("Price method:", Price)

	/*
		Попытка использовать в срезе два разных типа выдаст результат:
			products := map[string]store.Product{
				"Kayak": store.NewBoat("Kayak", 279, 1, false),
				"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
			}

		for _, p := range products {
			fmt.Println("Name:", p.Name, "Category:", p.Category, "Price", p.Price(0.2))
		}

		./main.go:51:12: cannot use store.NewBoat("Kayak", 279, 1, false) (value of type *store.Boat) as type store.Product in map literal
		./main.go:52:12: cannot use store.NewProduct("Soccer Ball", "Soccer", 19.50) (value of type *store.Product) as type store.Product in map literal
	*/

	/*
		При создании и использовании интерфейса создание среза с разными типами завершится успехом
	*/
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	//for key, p := range products {
	//	fmt.Println("Key:", key, "Price", p.Price(0.2))
	//}

	/*
		Key: Kayak Price 334.8
		Key: Ball Price 23.4
	*/

	/*
			Вернется с ошибкой попытка обратиться к значениям
			for key, p := range products {
			switch item := p.(type) {
							case *store.Product, *store.Boat:
								fmt.Println("Name:", item.Name, "Cetegory:", item.Category, "Price:", item.Price(0.2))
							default:
								fmt.Println("Key:", key, "Price", p.Price(0.2))
							}
						}
					./main.go:89:30: item.Name undefined (type store.ItemForSale has no field or method Name)
					./main.go:89:54: item.Category undefined (type store.ItemForSale has no field or method Category)
			так как у интерфейса нет полей которые мы запрашиваем

		При обращении к конкретному типу
			for key, p := range products {
			switch item := p.(type) {
			case *store.Product:
				fmt.Println("Name:", item.Name, "Cetegory:", item.Category, "Price:", item.Price(0.2))
			case *store.Boat:
				fmt.Println("Name:", item.Name, "Cetegory:", item.Category, "Price:", item.Price(0.2))
			default:
				fmt.Println("Key:", key, "Price", p.Price(0.2))
			}
			}
			мы получим ожидаемый результат
				Name: Kayak Cetegory: Watersports Price: 334.8
				Name: Soccer Ball Cetegory: Soccer Price: 23.4

	*/

	/*
				Вместо полей попробуем использовать интерфейсы

				for key, p := range products {
					switch item := p.(type) {
					case store.Describable:
						fmt.Println("Name:", item.GetName(), "Cetegory:", item.GetCategory(), "Price:", item.(store.ItemForSale).Price(0.2))
					default:
						fmt.Println("Key:", key, "Price", p.Price(0.2))
					}
				}
			Получи ожидаемый результат:
				Key: Kayak Price 334.8
				Key: Ball Price 23.4
		    Но мы обращались отдельно к интерфейсу - ItemForSale.

			Эта проблема устраняется добавлением интерфейса ItemForSale в интерфейс Describable
	*/
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name:", item.GetName(), "Cetegory:", item.GetCategory(), "Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price", p.Price(0.2))
		}
	}
}
