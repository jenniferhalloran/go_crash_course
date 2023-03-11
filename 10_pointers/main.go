package main 

import "fmt"

func main() {
	a := 5
	b := &a  //&a is the memory address
	//*b is the pointer
	fmt.Println(a, b, *b, *&a)


	//change val with pointer 

	*b = 10 
	fmt.Println(a)
}