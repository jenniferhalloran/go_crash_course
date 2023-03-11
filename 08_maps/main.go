package main 

import "fmt"

func main() {
	//define map 
	// emails := make(map[string]string)

	// //Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Marge"] = "marge@gmail.com"

	// fmt.Println(emails["Bob"])

	// delete(emails, "Bob")

	// declare and assign at same time
	emails := map[string]string{"Bob":"bobbo@gmail.com", "Sharon":"sharon@gmail.com"}

	fmt.Println(emails)

}