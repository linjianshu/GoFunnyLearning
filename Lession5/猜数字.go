package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i int
	fmt.Println("请输入100以内要猜的数字")
	fmt.Scanf("%d", &i)
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	for {
		intn := rand.Intn(100)
		if intn == i {
			fmt.Println(i)
			break
		}
		fmt.Println(intn)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("用时: ", time.Now().Sub(now).String())
}
