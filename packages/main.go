package main

/*
альтернативная запись currencyFmt "packages/fmt" - ". "packages/fmt"", но в пакете не должно быть имен используеиых в пакете
_ "packages/data" - пустой идентификатор (символ подчеркивания), позволяет импортировать пакет, не требуя использования его экспортируемых функций
"github.com/fatih/color"
*/
import (
	"fmt"
	//"github.com/fatih/color"
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

/*
При попытке обращения к полю price возникает ошибка
./main.go:14:3: unknown field 'price' in struct literal of type store.Product
./main.go:18:32: product.price undefined (type store.Product has no field or method price)
*/
func main() {
	//fmt.Println("Hello, Packages and Modules")

	product := store.NewProduct("Kayak", "Watersport", 279)
	cartNew := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	//fmt.Println("Name", product.Name)
	//fmt.Println("Category", product.Category)
	//fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	//fmt.Println("Name", cart.CustomerName)
	//fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))

	//color.Green("Name: " + cartNew.CustomerName)
	//color.Cyan("Total: " + currencyFmt.ToCurrency(cartNew.GetTotal()))

	fmt.Println("Name", cartNew.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cartNew.GetTotal()))
}
