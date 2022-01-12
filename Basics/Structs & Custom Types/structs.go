package main
import "fmt"

//Note: Structs are used to store multiple values of different data types into a single variable.
//Struts is the collection of fields

type Person struct{
	name string
	age int
	job string
	salary int
}


//Print person info
func printPersonInfo(p Person){
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Job: ", p.job)
	fmt.Println("Salary: ", p.salary)
}


func main(){

	//Basic 
	var per1 Person
	var per2 Person

	//per1 data
	per1.name = "Prakash Rai"
	per1.age = 23
	per1.job = "Software Engineer"
	per1.salary = 30000

	//per2 data
	per2.name = "Nikita Rai"
	per2.age = 23
	per2.job = "Frontend Developer"
	per2.salary = 25000

	//Accessing and printing the per1 and per2 info
	fmt.Println("Per1 Name:", per1.name)
	fmt.Println("Per2 Job:", per2.job)

	//Passing struct as functional arguments
	printPersonInfo(per1)

}

