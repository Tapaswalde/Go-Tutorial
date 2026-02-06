package main
import "fmt"

func main(){
	var name string="GO"
	// once a variable is declared, it cannot be redeclared in the same scope.
	// we can assign a new value to the variable but we cannot change its type.
	// name= "Golang"
	fmt.Println(name)

	var isAdult bool=true
	fmt.Println(isAdult)

	//walrus operator is used to declare and initialize a variable in a single line.
	age := 25
	fmt.Println(age, "is of type", fmt.Sprintf("%T", age))

	var z = "Go"
	fmt.Println(z)
}