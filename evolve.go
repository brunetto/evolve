package main 

import (
	"fmt"
	
	
	
)

func main () () {
	var particle = &Body{}
	fmt.Println(*particle)
	
}



type Body struct {
	Mass float64
	Pos [3]float64
	Vel [3]float64
}





