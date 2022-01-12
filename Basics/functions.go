package main
import ("fmt"
    "strings"
)

//declaring return value type for named return
func add(x int, y int) (result int){
    // return x + y
    result = x + y
    return
}
//declaring return value type
func sub(a int, b int) int{
    return a - b
}

//Multiple return values
func multipleType(v1 int,v2 string) (num int, text string){
    num = v1 + v1
    text = v2 + " Added Text"
    return
}

//Multiple return values 2
func getNameInitials(name string)(string, string){
    upperCaseName := strings.ToUpper(name)
    names := strings.Split(upperCaseName, " ") //Splits the string into all substrings separated by the given separator and returns a slice which contains these substrings.

    var initials []string
    for _, v := range names{
        initials = append(initials, v[:1]) //use arrange. v[:1] means go from positon 0 upto position 1 but not including position 1.
    }

    if len(initials) > 1 {
        return initials[0], initials[1]
    }

    return initials[0], "_" //If there is no second name return "_" as 2nd string
}

//Recursion function
func countTest(x int) int{
    if x == 10{
        return 0
    }
    fmt.Println(x)
    return countTest(x+1)
}

func greeting(n string){
    fmt.Println("Ohayo",n)
}

func bye(n string){
    fmt.Println("Sayonara",n)
}

//Iterating through the slice and string value is passed as argument in "f" function
func loopNames(n []string, f func(string)){
    for _, v := range n {
        f(v)
    }
}

func main(){
    fmt.Println(add(5,5))
    
    total := add(6,4)
    fmt.Println("total",total)

    fmt.Println("Subtracted Value", sub(9,10))

    num, text :=  multipleType(5, "Your")
    fmt.Println("Multiple return output",num, text)

    //Omitting the returned value
    _, v1 := multipleType(5, "Hello")
    fmt.Println(v1)

    //Call recursive function
    countTest(5)

    //Ad function
    testFunc := func(value int) int{
        return value + 100
    }(8)

    //or

    // testFunc(8)
    fmt.Println(testFunc)

    //function inside another function
    loopNames([]string{"Prakash","Sawal","Reyzan"}, bye)


    //Multiple value return example
    fn1, sn1 := getNameInitials("prakash")
    fmt.Println(fn1, sn1)

    fn2, sn2 := getNameInitials("sawal timsina")
    fmt.Println(fn2, sn2)
}
