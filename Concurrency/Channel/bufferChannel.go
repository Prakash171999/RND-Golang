package main
import ("fmt"
"time")

func receiver(ch <-chan int){
	fmt.Println(<-ch)
}

func main(){
	fmt.Println("Channels")

	//Takes two argument type and capacity of the channel
	ch := make(chan int, 2)

	go receiver(ch)
	ch <- 1
	ch <-2
	ch <-3
	// ch <-5

	time.Sleep(1000 * time.Millisecond)
}