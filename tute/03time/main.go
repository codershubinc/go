package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("time to study time")
	presentTime := time.Now()
	fmt.Println("present time is: ", presentTime)
	
	// Formatting time
	fmt.Println("formatted time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))
	
}
