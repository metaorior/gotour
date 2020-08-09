package main 


import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// tell that p value is equal to i value
	p := &i 
	fmt.Println("the value of p is ", *p) // read i through the pointer
	*p = 21 //set i throught the pointer
	fmt.Println("the value of i is ", i) // see the new value of j

	p = &j // p now is the value of j
	fmt.Println("p = &j is " , j)
	*p = *p / 37 //divide j through the pointer
	fmt.Println("the value of j is ", j) // see the new value of j 
	
	z := 6
	var a int
	for  a = 0; a < z; a++ {
		fmt.Println(a)
	}
	

 }