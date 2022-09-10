package main

import (
	"fmt"
)

func main() {
	//说明fallthrough只能将下降延续到下一个case语句中 下下个就不行了 并且啊 fallthrough的时候还不需要判断条件就可以执行
	score := 79
	switch {
	case score < 100:
		fmt.Println("fuck")
		fallthrough
	case score == 90:
		fmt.Println("okok")
		fallthrough
	case score < 80:
		fmt.Println("just so so")
	case score < 60:
		fmt.Println("shit shit shit")
	}
}
