package main

import (
	"fmt"
	"errors"
)

func main(){
	total, err := sumOf(2,10)

	if err != nil{
		fmt.Println("There was an error:", err)
	}else{
		fmt.Println("Total value:", total)
	}
}

func sumOf(start int, end int) (int, error) {

	if start > end {
		return 0, errors.New("Start is greater than end")
	}
	total := 0
	for i:=start; i<= end; i++ {
		total +=i
	}

	return total, nil
}