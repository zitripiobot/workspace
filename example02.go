package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Printf("Now you %g problems. \n", math.Sqrt(7))
	fmt.Println("Now you have %g problems.",math.Nextafter(2,3))
	fmt.Println("Happy", math.Pi,"Day")
}



