package code

import (
	"fmt"
	"time"
)

//import (
//	"fmt"
//	"time"
//)
//
//func goroutinePrint123() {
//	var ch1 = make(chan struct{})
//	var ch2 = make(chan struct{})
//	var ch3 = make(chan struct{})
//
//	go func() {
//		for {
//			<-ch1
//			fmt.Println("1")
//			time.Sleep(time.Second)
//			ch2 <- struct{}{}
//		}
//	}()
//	go func() {
//		for {
//			<-ch2
//			fmt.Println("2")
//			time.Sleep(time.Second)
//			ch3 <- struct{}{}
//		}
//	}()
//	go func() {
//		for {
//			<-ch3
//			fmt.Println("3")
//			time.Sleep(time.Second)
//			ch1 <- struct{}{}
//		}
//	}()
//	ch1 <- struct{}{}
//	time.Sleep(time.Second * 5)
//}

func goroutinePrint123() {
	var ch1 = make(chan struct{})
	var ch2 = make(chan struct{})
	var ch3 = make(chan struct{})

	go func() {
		for {
			<-ch1
			fmt.Println(1)
			time.Sleep(time.Second)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println(2)
			time.Sleep(time.Second)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch3
			fmt.Println(3)
			time.Sleep(time.Second)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	time.Sleep(time.Second * 10)
}
