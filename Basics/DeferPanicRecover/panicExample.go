package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey")
	panicExample()
	fmt.Println("This won't run")
}

/*
Panic is a way in which one can stop the program from running.
For eg. If a you can't connect to DB or server, you can use panic to force stop your program
and tell user progam cant connect to DB or server.
*/ 
func panicExample(){
panic("We have paniced!!!")
}