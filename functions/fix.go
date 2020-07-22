package main

import (
	"fmt"
	//"math"
)

func main() {
	fmt.Println("hello world")

	var x string
	x = "caio"
	fmt.Println(x)

	z := "caio"

	fmt.Println(z)
	//testNames()
	var result int
	result = validatePlayer(30, 1001, 2.000)
	fmt.Println(result)
}

func 

//exported names

// func testNames() {
// 	//fmt.Println(math.pi)
// 	fmt.Println(math.Pi)
	
// }




// in this case the function will verify the level int , the elo int and the xp float and return a bool which should only be 1 or 0
func validatePlayer(level int, elo int32, xp float32) int {
	
	if (level <= 0) {
		fmt.Println("Hero is below level 0! error")
		return 1
	}
	
	if (elo < 1000 || elo > 1500) {
		fmt.Println("Player is noob or hacker ! ")
		return 1

	}

	if (xp < 0 ) {
		fmt.Println("Player has negative xp !")
		return 1
	}
	

	//else
	fmt.Println("Player is allowed :) ")
	return 0

}

