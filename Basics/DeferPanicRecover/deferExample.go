package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey")
	deferExample()
}

/*
Defer allows you to run a function where the function stops.
You ca use multiple deffers in a function
*/ 
func deferExample(){

	/*the print statement will fun after the function stops.
	  Defer is used because it will run even after the use of panic
	*/
	defer fmt.Println("0") 
	
	fmt.Println("1")
	panic("We paniced for some reason")
	fmt.Println("2")
}

/*
	For example: Assume you opened a file and you want to close it no matter what happerns.
	Even if the function panics you want to close the file
*/