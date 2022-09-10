package main

import (
	"fmt"
	"time"
)

var c = make(chan int)

func main() {
	//go add()
	//time.Sleep(time.Second)
	//for v := range c {
	//	fmt.Println(v)
	//}

	c1 := make(chan int)
	c2 := make(chan int)
	var a = func() {
		select {
		case <-c1:
			fmt.Println("A ...")
			time.Sleep(time.Second)
			c2 <- 1
		}
	}

	var b = func() {
		select {
		case <-c2:
			fmt.Println("B ...")
			time.Sleep(time.Second)
			c1 <- 1
		}
	}

	go func() {
		c1 <- 1
	}()
	go func() {
		for {
			a()
		}
	}()
	go func() {
		for {
			b()
		}
	}()

	time.Sleep(4 * time.Second)
}

func add() {
	for i := 0; i < 10; i++ {
		//time.Sleep(1*time.Second)
		c <- i
	}
	close(c)
}
