package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 10
	for count > 0 {
		intn := rand.Intn(100)
		if intn == 0 {
			fmt.Println("发射失败")
			break
		} else {
			fmt.Println(count)
			count--
			time.Sleep(1 * time.Second)
		}
	}
}
