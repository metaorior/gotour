// In Go, a name is exported if it begins with a capital letter. For example, Pizza
// is an exporte name, as is Pi, which is exported from the math package
// pizza and pi aren't capital so they aren't exported
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

// Run the code. Notice the error message.

// To fix the error, rename math.pi to math.Pi and try it again

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(math.pi)
// }

package main

import(
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}