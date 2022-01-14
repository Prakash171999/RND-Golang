package main

import ("fmt"
"time")

func receiver(ch chan int){
	//Infinite for loop.It will read values as long as we are sending values
	// for {
	// 	i, ok := <-ch
	// 	if ok == false {
	// 		fmt.Println(i, ok, "<-- loop broker!")
	// 		break  //exit break loop
	// 	}else {
	// 		fmt.Println(i, ok)
	// 	}
	// }

	//Better approach
	for val := range ch{
		fmt.Println(val)
	}

	fmt.Println("Channel closed by sender")


	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

func main(){
	fmt.Println("Channels")
	
	ch := make(chan int)

	go receiver(ch) //This goroutine will be blocked until it reads value from channel ch

	ch <- 42
	ch <- 43
	ch <- 44
	ch <- 45
	close(ch) //closing channel
	// ch <- 25

	time.Sleep(100 * time.Millisecond)

}