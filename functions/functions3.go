package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is functions 3")
	callAddString()

	a,b := swap("Hello", "World")
	fmt.Println(a,b)
}


func addString(word1 string, word2 string) string {
	return word1 + " " + word2
}

func callAddString() {
	fmt.Println(addString("caio" ,"sarah"))
}

func swap(x, y string) (string, string) {
	return y, x
}




