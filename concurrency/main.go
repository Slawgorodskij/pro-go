package main

import (
	"fmt"
	"time"
)

//func receiveDispatches(chanhel <-chan DispatchNotification) {
//	for details := range chanhel {
//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//	}
//	fmt.Println("Channel has been closed")
//}
//func main() {
//fmt.Println("main function started")
//CalcStoreTotal(Products)
//time.Sleep(time.Second * 5)
//fmt.Println("main function complete")
//
//dispatchChannel := make(chan DispatchNotification, 100)
//go DispatchOrder(dispatchChannel)
//for {
//	if details, open := <-dispatchChannel; open {
//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//	} else {
//		fmt.Println("Channel has been closed")
//		break
//	}
//
//}
//for details := range dispatchChannel {
//	fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//}
//fmt.Println("Channel has been closed")

//dispatchChannel := make(chan DispatchNotification, 100)
//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
//go DispatchOrder(sendOnlyChannel)
//receiveDispatches(receiveOnlyChannel)

//go DispatchOrder(chan<- DispatchNotification(dispatchChannel))
//receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

//	dispatchChannel := make(chan DispatchNotification, 100)
//	go DispatchOrder(chan<- DispatchNotification(dispatchChannel))
//
//	for {
//		select {
//		case details, ok := <-dispatchChannel:
//			if ok {
//				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//			} else {
//				fmt.Println("Channel has been closed")
//				goto alldone
//			}
//		default:
//			fmt.Println("== No message ready to be received")
//			time.Sleep(time.Millisecond * 500)
//		}
//	}
//alldone:
//	fmt.Println("all values received")
//}
//
//func enumerateProducts(channel chan<- *Product) {
//	for _, p := range ProductList[:3] {
//		channel <- p
//		time.Sleep(time.Millisecond * 800)
//	}
//	close(channel)
//}
//
//func main() {
//	dispatchChannel := make(chan DispatchNotification, 100)
//	go DispatchOrder(chan<- DispatchNotification(dispatchChannel))
//
//	productChannel := make(chan *Product)
//	go enumerateProducts(productChannel)
//	openChannels := 2
//	for {
//		select {
//		case details, ok := <-dispatchChannel:
//			if ok {
//				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//			} else {
//				fmt.Println("Dispatch channel has been closed")
//				dispatchChannel = nil
//				openChannels--
//			}
//		case product, ok := <-productChannel:
//			if ok {
//				fmt.Println("Product:", product.Name)
//			} else {
//				fmt.Println("Product channel has been closed")
//				productChannel = nil
//				openChannels--
//			}
//
//		default:
//			if openChannels == 0 {
//				goto alldone
//			}
//			fmt.Println("-- No message ready to be received")
//			time.Sleep(time.Millisecond * 500)
//		}
//	}
//alldone:
//	fmt.Println("all values received")
//}

/*Отправка без блокировки*/

//func enumerateProducts(channel chan<- *Product) {
//	for _, p := range ProductList {
//		select {
//		case channel <- p:
//			//fmt.Println("Send product:", p.Name)
//		default:
//			fmt.Println("Discarding product:", p.Name)
//			time.Sleep(time.Second)
//		}
//	}
//	close(channel)
//}
//
//func main() {
//	dispatchChannel := make(chan DispatchNotification, 100)
//	go DispatchOrder(chan<- DispatchNotification(dispatchChannel))
//
//	productChannel := make(chan *Product, 5)
//	go enumerateProducts(productChannel)
//	time.Sleep(time.Second)
//	for p := range productChannel {
//		fmt.Println("Received product:", p.Name)
//	}
//}

/* ОТПРАВКА НА НЕСКОЛЬКО КАНАЛОВ*/

func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel1")
		case channel2 <- p:
			fmt.Println("Send via channel2")
		}
	}
	close(channel1)
	close(channel2)
}

func main() {
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProducts(c1, c2)
	time.Sleep(time.Second)

	for p := range c1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range c2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}
}
