package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is functions 3")
	callAddString()
}


func addString(word1 string, word2 string) string {
	return word1 + " " + word2
}

func callAddString() {
	fmt.Println(addString("caio" ,"sarah"))
}
