package main

import "fmt"

//func (slice ProductSlice) TotalPrice(category string) (total float64) {
//	for _, p := range slice {
//		if p.Category == category {
//			total += p.Price
//		}
//	}
//	return
//}

/*ГЕНЕРАЦИЯ ОШИБОК*/

//type CategoryError struct {
//	requestedCategory string
//}
//
//func (e *CategoryError) Error() string {
//	return "Category " + e.requestedCategory + " does not exist"
//}
//func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
//	productCount := 0
//	for _, p := range slice {
//		if p.Category == category {
//			total += p.Price
//			productCount++
//		}
//	}
//	if productCount == 0 {
//		err = &CategoryError{requestedCategory: category}
//	}
//	return
//}

/*	 СООБЩЕНИЕ ОБ ОШИБКАХ ЧЕРЕЗ КАНАЛЫ  */

//type CategoryError struct {
//	requestedCategory string
//}
//
//func (e *CategoryError) Error() string {
//	return "Category " + e.requestedCategory + " does not exist"
//}
//
//type ChannelMessage struct {
//	Category string
//	Total    float64
//	*CategoryError
//}
//
//func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
//	productCount := 0
//	for _, p := range slice {
//		if p.Category == category {
//			total += p.Price
//			productCount++
//		}
//	}
//	if productCount == 0 {
//		err = &CategoryError{requestedCategory: category}
//	}
//	return
//}
//
//func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
//	for _, c := range categories {
//		total, err := slice.TotalPrice(c)
//		channel <- ChannelMessage{
//			Category:      c,
//			Total:         total,
//			CategoryError: err,
//		}
//	}
//	close(channel)
//}

/*  ИСПОЛЬЗОВАНИЕ УДОБНЫХ ФУНКЦИЙ ОБРАБОТОК ОШИБОК  */

type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		//err = errors.New("Cannot find category")
		err = fmt.Errorf("Cannot find category: %v", category)
	}
	return
}

func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
