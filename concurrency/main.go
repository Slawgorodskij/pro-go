package main

import "fmt"

func main() {
	//fmt.Println("main function started")
	//CalcStoreTotal(Products)
	//time.Sleep(time.Second * 5)
	//fmt.Println("main function complete")

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrder(dispatchChannel)
	//for {
	//	if details, open := <-dispatchChannel; open {
	//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	//	} else {
	//		fmt.Println("Channel has been closed")
	//		break
	//	}
	//
	//}
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}
