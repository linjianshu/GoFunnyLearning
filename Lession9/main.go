package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	//直接对数字执行string 不会把33=>"33" 而是会把33代表的字符char找出来
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))
	fmt.Printf("%c %c %c %c \n", pi, alpha, omega, bang)

	countdown := 10
	//将数字直接变成字符串 而不是char 使用itoa或者sprint
	fmt.Println("launch in T minus " + strconv.Itoa(countdown) + " seconds.")
	fmt.Println("launch in T minus " + fmt.Sprintf("%v", countdown) + " seconds.")

	a := 42
	b := &a
	fmt.Printf("%T", b)

	var administrator *string
	//这个指针类型没有指向任何地方 会报错
	//fmt.Println(*administrator)

	sclese := "Christpher J. Scolese"
	administrator = &sclese
	fmt.Println(*administrator)

	*administrator = "hello world"
	fmt.Println(*administrator)
	fmt.Println(sclese)

	another := administrator
	//他们存的地址是不是一样的
	fmt.Println(another)
	fmt.Println(another == administrator)
	//他们自身地址是不是一样的
	fmt.Println(&another, "  ", &administrator)
	fmt.Println(&another == &administrator)
	//他们存的地址指向的值是不是一样的
	fmt.Println(*another)
	fmt.Println(*another == *administrator)

	c := [...]int{1, 2, 3}
	c[1] = 3
	fmt.Println(c)

	d := [8][8]rune{{1, 2, 3, 4, 5, 6, 7, 8}, {}, {}, {}, {}, {}, {}, {}}
	fmt.Println(d)
	d[0] = [...]rune{2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(d)

	e := "hello world"
	//不能对字符串的索引进行操作 只能读不能写
	//e[5]='h'
	fmt.Println(e)

	var slice []int
	fmt.Println(slice == nil)

}

type sudokuError []error

//方法不是结构体独有的 类型也可以有方法 例如string.contains
func (se sudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}
