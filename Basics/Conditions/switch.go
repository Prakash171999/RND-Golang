package main
import ("fmt";"time")

func switchTest(){
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func main(){
	day := "saturday"

	switch day{
	case "monday","tuesday":
		fmt.Println("Visit Dharan")
	case "wednesday":
		fmt.Println("Visit Kathmandu")
	case "saturday","sunday":
		fmt.Println("It's holiday")
	default:
		fmt.Println("Go to office")
	}
	
	switchTest()
}