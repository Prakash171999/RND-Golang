package main
import "fmt"
 
 
//Pointers with function
func myFucntionWithPointer(s *string){
   fmt.Println(s, *s) //Here s will output the memory address passed from the function and *s will output the variable which it is pointing to.
}
 
func main(){
   i, j := 100, 200
 
   //& is written before the variable name to get the memory address of that variable
   fmt.Println("Get the memory address of the variable:",&i, &j)
  
   //Assigning the address of i to p. So now p is a pointer
   p:= &i
 
   /*
       If * is infront of a type name (* int) then the whole thing includin the star is a type (pointer type) where int is the base
       But if * is infront of the variable with a pointer type then * operates as an operator which return what p is pointing to.
   */
   fmt.Println(*p) //This will print 42. Because its the value of the variable p is pointing to. Also known as dereferencing or indirecting
 
   *p = 21 //Changing the value of *p will change the value of i
   fmt.Println(i)
 
   p = &j
   *p = *p / 2
   fmt.Println("Value in J:",j)
 
   //Why use Pointer
   /*
   Its efficient to store the variable in one place and access it from here and there.You can share the variable and update it in multiple places
   in the program which is more efficient than copying the variable everytime you need to make changes to them or use them.
 
   If you want to manipulate a value across function calls you need to use pointers.
   */
 
   //Pointer with functions
   myOtherString := "Another message here"
   myFucntionWithPointer(&myOtherString)
 
}
