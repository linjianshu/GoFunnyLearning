package main

import "fmt"

func main() {
	drawTable(drawF())
}

func drawTable(CorF func()) {
	CorF()
	drawLine()
}

func drawLine() {
	fmt.Println("=========================")
}
func drawFill(s, s2 interface{}) {
	fmt.Printf("|%11v|%11v|\n", s, s2)
}

func drawC() func() {
	return func() {
		drawLine()
		drawFill("°C", "°F")
		drawLine()
		for i := -40; i <= 100; i += 5 {
			drawFill(float32(i), float32(i)+(33.8))
		}
	}
}

func drawF() func() {
	return func() {
		drawLine()
		drawFill("°F", "°C")
		drawLine()
		for i := -40; i <= 100; i += 5 {
			drawFill(float32(i)+(33.8), float32(i))
		}
	}
}
