package main

import (
	"fmt"
	"time"
)


func main() {

	//Timers are for when you want to do something once in the future
	// tickers are for when you want to do something many times
	// 
	ticker := time.NewTicker(500 * time.Milisecond)
	done := make(chan bool)

	// tickers use a similiar logic as timers. a channel that
	//is sent values. Here we will use the select builtin
	// on the channel to wait the values as they arrive every 500ms

	

}