package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey")
	recoverExample()
}

/*
	Recover lets you recover from panic
*/
func recoverExample(){


	defer func(){
		if r := recover(); r!= nil {
			fmt.Println("The function recovered from panic with reason:", r)
		}
	}()

	fmt.Println("1")
	// panic("Some reason")
	fmt.Println("2")
}