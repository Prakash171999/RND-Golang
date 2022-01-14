 package main

 import ("fmt"; "time")

 func main(){

	//Goroutine is created by adding keyword go.
	//Goroutine is managed by the go runtime
	go greeter("Hello")

	greeter("World")
 }

 /*Here we added the Sleep() method this function which makes the main Goroutine sleep for 3 millisecond
   In between 3 millisecond the new Goroutine execute, displays "Hello" on the screen, and then terminate after 3 millisecond.
   The process continues until i < 6. After that the main routine executes. Here, both Goroutine and the normal function work concurrently.
 */
 func greeter(s string){
	 for i := 0; i < 6; i++ {
		 time.Sleep(3 * time.Millisecond)
		 fmt.Println(s)
	 }
 }