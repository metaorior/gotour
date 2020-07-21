// functions in go are abreviated to func, for reading purposes
// a function can take zero or more arguments
// in this example. add takes two parameters of type int
package main 

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(mult(4, 37))
}

func divide(a int, b int) int {
	return a / b

}

func mult(d int, e int) int {
	return d * e
}