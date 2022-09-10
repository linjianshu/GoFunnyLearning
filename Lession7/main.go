package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%5.2f\n", third)
	fmt.Printf("%05.2f\n", third)

	//复用格式化变量
	day := 365
	fmt.Printf("one year has %05d days , day is %[1]T type \n", day)
	fmt.Printf("day is %T type , one year has %05[1]d days\n", day)

	//十六进制
	fmt.Printf("%05b \n", 0xa)

	//十进制转二进制 用0填充并限定位数
	var green uint8 = 3
	fmt.Printf("%08b\n", green)

	//使用int64解决int32只能表示到2038年的问题
	unix := time.Unix(12622780800, 0)
	fmt.Println(unix)

	//利用指数来减少0的键入次数
	fmt.Println(10e2)

	//大数
	b := new(big.Int)
	b.SetString("2400000000", 10)
	fmt.Println(b)

	newInt := big.NewInt(240000000)
	fmt.Println(newInt)

	b.Div(b, newInt)
	fmt.Println(b)

	const name = 1 //不会自动推断类型 类型是untyped
	//字面量和常量底层将由big包提供支持 不会为常量和字面量推断类型
	//常量和bigint无法转换

	const distance = 236000000000000000
	const guangnian = 100000000
	//b2 := new(big.Int)
	//setString, _ := b2.SetString("236000000000000000", 10)
	//setString.Div(setString,guangnian)
	fmt.Println(distance / guangnian)

	//使用反引号包围的叫做原始字符字面量
	fmt.Printf("hello world \t\t")
	fmt.Printf(`hello world \t\t`)

	fmt.Printf("\n%c", 65)
	fmt.Printf("\n%c", 128515)
	var rune = 65
	fmt.Printf("\n%c", rune)
	fmt.Println(strconv.Atoi(string('😃'))) //这样不行 只能给string进行转换成int 比如"110"->110
	fmt.Printf("%d\n", '😃')                //这样可以给字符转成int

	s := "shalom"
	for _, c := range s {
		fmt.Printf("%d\n", c)
	}

	fmt.Println('A' - 'a')
	fmt.Printf("%c\n", 'x'+3-26)

	//forrange可以解析UTF-8的编码 所以这里遍历是没有问题的
	for _, v := range "镜中花,水中月" {
		fmt.Printf("%c %[1]T \n", v)
	}

	//fori不能解析UTF-8编码 求len的时候就已经出问题了
	shi := "镜中花,水中月"
	for i := 0; i < len(shi); i++ {
		fmt.Printf("%c %[1]T \n", shi[i])
	}
	fmt.Println(len("镜中花,水中月")) //len返回的是字符串占位的字节长度 这里返回19

	//使用utf8包和函数
	//解析第一个字占几字节  具体值是多少
	decodeRune, size := utf8.DecodeRune([]byte(shi))
	fmt.Println(size)
	fmt.Printf("%c \n", decodeRune)
	//查看一共解析了几个字
	inString := utf8.RuneCountInString(shi)
	fmt.Println(inString)
	//解析第一个字占几字节  具体值是多少
	runeInString, size1 := utf8.DecodeRuneInString(shi)
	fmt.Printf("%c %d bytes\n", runeInString, size1)

	s3 := "abcdefghijklmnopqrstuvwxyz"
	decodeRuneInString, i := utf8.DecodeRuneInString(s3)
	fmt.Println(decodeRuneInString, " ", i)

	//判断它占了几个字节
	r, size2 := utf8.DecodeRuneInString("¿")
	fmt.Println(r, " ", size2)

	//go使用UTF-8可变长度编码,每个字符根据需要占用1-4个字节的内存空间
	//可能会问fori和forr的区别吧

	println(caesar("L fdph,L vdz,L frqtxhuhg."))

	println(strings.Repeat("helloworld helloworld hello", 3))

	var bh float64 = 32768
	if bh > math.MaxInt16 || bh < math.MinInt16 {
		fmt.Println(bh > math.MaxInt16 || bh < math.MinInt16)
	}
}

func caesar(s string) (final string) {
	for _, v := range s {
		if v >= 'A' && v <= 'z' {
			if v <= 'c' && v >= 'a' {
				v = v + 26
			}
			final += string(v - 3)
		} else {
			final += string(v)
		}
	}
	return final
}
