package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(10) + 1
	fmt.Println(intn)
	fmt.Println("apppppple" > "banana")
}
