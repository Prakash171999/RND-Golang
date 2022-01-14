package main

import "fmt"

func receiver(ch chan int){
	fmt.Println(<-ch)
}

func main(){
	fmt.Println("Channels")
	
	ch := make(chan int)

	go receiver(ch) //This goroutine will be blocked until it reads value from channel ch

	ch <- 42

}