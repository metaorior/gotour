package main

import (
"strconv"
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "world")
	fmt.Println(a,b)


	num1, num2, num3 := crazy(700, 200, 10000)
	fmt.Println(num1, num2, num3)
}

func crazy(a,b,c int) (string, string, string) {
  convert := strconv.Itoa
	return (convert(b) + convert(c) + convert(a)),(convert(b) + convert(c) + convert(a)),(convert(b) + convert(c) + convert(a))
}

// func crazyInverter3(john, car, trash, mouse, m90 double) ()
//
