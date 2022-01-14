/*
	Anonymous functions allows us to inline code by creating a function with no name.
	Combined with closures, this can simply function signatures for goroutines.

	Can be used rather than having a separate function
*/
package main

import ("fmt"
"time")


func main(){

	fmt.Println("Welcome!! to the new world")

	//Anonymous function
	go func(){
		fmt.Println("Welcome!! to the Earth")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! See you again")
	
}