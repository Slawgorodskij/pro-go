package main

import "fmt"

type calcFunc func(float64) float64

// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// func printPrice(product string, price float64, calculator func(float64) float64) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) func(float64) float64 {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

//добавление псевдонима
func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		var withTax calcFunc = func(price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 		return withTax
// 	}
// 	withoutTax := func(price float64) float64 {
// 		return price
// 	}
// 	return withoutTax
// }

// функция withTax доступна только в блоке if - в результате код не выполнится - undefined: withTax
// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		var withTax calcFunc = func(price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 		return withTax
// 	} else if price < 10 {
// 		return withTax
// 	}
// 	withoutTax := func(price float64) float64 {
// 		return price
// 	}
// 	return withoutTax
// }

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return func(price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 	}
// 	return func(price float64) float64 {
// 		return price
// 	}
// }

//ЗАМЫКАНИЯ
// func priceCalcFactory(threshold, rate float64) calcFunc {
// 	return func(price float64) float64 {
// 		if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

var prizeCiveaway = false

// цены упадут до нуля так как использоваться будет при выполнении priceCalcFactory (в обоих случаях) переменная prizeCiveaway = true
// func priceCalcFactory(threshold, rate float64) calcFunc {
// 	return func(price float64) float64 {
// 		if prizeCiveaway {
// 			return 0
// 		} else if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

// значение prizeCiveaway будет использоваться текущее на момент создания функции
// func priceCalcFactory(threshold, rate float64) calcFunc {
// 	return func(price float64) float64 {
// 		fixedPrizeCiveaway := prizeCiveaway
// 		if fixedPrizeCiveaway {
// 			return 0
// 		} else if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

// Ещё один вариант
// func priceCalcFactory(threshold, rate float64, zeroPrices bool) calcFunc {
// 	return func(price float64) float64 {
// 		if zeroPrices {
// 			return 0
// 		} else if price > threshold {
// 			return price + (price * rate)
// 		}
// 		return price
// 	}
// }

//использование указателей для предотвращения ранней оценки
func priceCalcFactory(threshold, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {
	//fmt.Println("Hello, Functions Types")

	// products := map[string]float64{
	// 	"Kayak":      275,
	// 	"Lifejacket": 48.95,
	// }
	// for product, price := range products {
	// 	var calcFunc func(float64) float64
	// 	fmt.Println("Function assigned:", calcFunc == nil)
	// 	if price > 100 {
	// 		calcFunc = calcWithTax
	// 	} else {
	// 		calcFunc = calcWithoutTax
	// 	}
	// 	fmt.Println("Function assigned:", calcFunc == nil)
	// 	totalPrice := calcFunc(price)
	// 	fmt.Println("Product:", product, "Price:", totalPrice)
	// }

	// for product, price := range products {
	// 	if price > 100 {
	// 		printPrice(product, price, calcWithTax)
	// 	} else {
	// 		printPrice(product, price, calcWithoutTax)
	// 	}
	// }

	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }

	// for product, price := range products {
	// 	printPrice(product, price, func(price float64) float64 {
	// 		return price + (price * 0.2)
	// 	})
	// }

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	// calc := func(price float64) float64 {
	// 	if price > 100 {
	// 		return price + (price * 0.2)
	// 	}
	// 	return price
	// }
	// for product, price := range watersportsProducts {
	// 	printPrice(product, price, calc)
	// }

	// calc = func(price float64) float64 {
	// 	if price > 50 {
	// 		return price + (price * 0.1)
	// 	}
	// 	return price
	// }
	// for product, price := range soccerProducts {
	// 	printPrice(product, price, calc)
	// }

	// waterCalc := priceCalcFactory(100, 0.2)
	// soccerCalc := priceCalcFactory(50, 0.1)

	// prizeCiveaway = false
	// waterCalc := priceCalcFactory(100, 0.2)
	// prizeCiveaway = true
	// soccerCalc := priceCalcFactory(50, 0.1)

	// prizeCiveaway = false
	// waterCalc := priceCalcFactory(100, 0.2, prizeCiveaway)
	// prizeCiveaway = true
	// soccerCalc := priceCalcFactory(50, 0.1, prizeCiveaway)

	prizeCiveaway = false
	waterCalc := priceCalcFactory(100, 0.2, &prizeCiveaway)
	prizeCiveaway = true
	soccerCalc := priceCalcFactory(50, 0.1, &prizeCiveaway)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}
