package main

import (
	"fmt"
)

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	// a return statement without args returns the named return vals
	// this is known as a "naked return"
	return // same as return x,y
}

func main() {
	fmt.Println(split(17))
	
	fmt.Println(salarioPorEstado(1045))
}

func salarioPorEstado(salarioBase int) (sp,mg,rj,df int) {
	sp = salarioBase + 200
	mg = salarioBase + 50
	rj = salarioBase + 150
	df = salarioBase + 200

	return sp,mg,rj,df
}