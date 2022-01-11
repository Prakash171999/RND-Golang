package main
import ("fmt")

func main(){
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)
  
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("Length:",len(myslice2))
	fmt.Println("Capicity:",cap(myslice2))
	fmt.Println("myslice:",myslice2)

	//Create slice with make()
	myslice3 := make([]int,5,10)
	
	fmt.Println("myslice3 value:", myslice3)
	fmt.Println("myslice3 length:", len(myslice3))
	fmt.Println("myslice3 cap:", cap(myslice3))

	//Modify Slice
	prices := []int{10,20,30,40,50}
	prices[4] = 100

	fmt.Println("Modified prices:", prices)

	//Append new element to slice
	prices = append(prices,200,300,400,500,600,700,800,900)
	
	fmt.Println("Appended value in prices:", prices)

	//Append one slice to another slice
	mySlice4 := []int{1,2,3}
	mySlice5 := []int{10,20,30,40,50,-50}
	mySlice6 := append(mySlice4, mySlice5...)

	fmt.Println("mySlice6:", mySlice6)
}