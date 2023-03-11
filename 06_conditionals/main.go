package main 

import "fmt"

func main() {
	x := 15 
	y := 10 

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
		} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "green"
	
	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}