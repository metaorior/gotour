// When two or more consecutive named functions parameters share a type, you can ommit
// the type from all but the last

// in this example, we shortened

// x int, y int

// x, y int 

package main

import (
	"fmt"
)
//func add(x int, y int) int 
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	doCheck()
}

func doCheck() {
	checkElo(-132)
}

func checkElo(elo float32) {
	if (elo < 0) {
		fmt.Println("Hero has a negative elo!", elo)

	}
}