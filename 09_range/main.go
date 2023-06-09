package main 

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	//loop through ids 

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	//loop through ids not using index

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	// add ids together 
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map 
	emails := map[string]string{"Bob":"bobbo@gmail.com", "Sharon":"sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s's email is %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}