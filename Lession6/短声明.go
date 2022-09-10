package main

import "fmt"

func main() {
	//短声明就是在var不能使用的时候,替代var来声明变量,同时提高可读性,另外缩小变量的作用域范围,提高变量名的复用性
	//可以在case里放置if语句诶
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	if i := 1; i > 0 {
		fmt.Println("大于0")
	}

	switch i := 1; i {
	case 1:
		fmt.Println("这是1诶!")
	}

	switch i := 1; i > 0 {
	case true:
		fmt.Println("真的诶!")
	}

	switch i := 1; {
	case i == 1:
		fmt.Println("真的是1诶!!!")
	}
}
