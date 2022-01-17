package main
 
import (
    "context"
    "fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context){
	rID := ctx.Value("request-id")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timed out")
			return 
		default:
			fmt.Println("Doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(rID)
}

func main(){
	fmt.Println("Go context tutorial")
	// ctx := context.Background()
	// ctx := enrichContext(ctx)
	// doSomethingCool(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select{
	case <-ctx.Done():
		fmt.Println("Oh no, I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)

}