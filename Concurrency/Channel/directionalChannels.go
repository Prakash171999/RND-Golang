package main

import ("fmt")

//To make channel uni directional. add <- infromt of chan.
//The receiver will only receive values from channel but can't send value to channel
func receiver(ch <-chan int){

	fmt.Println("Printed in receiver function:",<-ch)
	// ch <- 12342  //Sending value to channel ch. This line will be invalid as the receiver function is made unidirectional channel

	fmt.Println("Channel closed by sender")

}

//
func sender(ch chan<- int){
	ch <- 12342
	fmt.Println("Channel closed by sender")
}

func main(){
	fmt.Println("Channels")
	
	ch := make(chan int)

	// go receiver(ch) //This goroutine will be blocked until it reads value from channel ch

	go sender(ch)

	// ch <- 45
	fmt.Println("Printed in main function:",<-ch)
	close(ch) //closing channel


}