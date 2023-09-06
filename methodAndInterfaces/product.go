package main

type Product struct {
	name, category string
	price          float64
}

//func (p Product) getName() string {
//	return p.name
//}
//
//func (p Product) getCost(_ bool) float64 {
//	return p.price
//}

func (p *Product) getName() string {
	return p.name
}

func (p *Product) getCost(_ bool) float64 {
	return p.price
}

/*
func (p *Product) getName() string {
	return p.name
}

func (p *Product) getCost(_ bool) float64 {
	return p.price
}

./main.go:296:24: cannot use product (variable of type Product) as type Expense in variable declaration:
        Product does not implement Expense (getCost method has pointer receiver)

*/
