package main

import (
	"fmt"
)

func main() {
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("other")
	// }

	//mulitple case switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("It is an Integer")
		case string:
			fmt.Println("It is a string")
		case bool:
			fmt.Println("It is a boolean")
		default:
			fmt.Println("other", v)
		}
	}
	whoAmI(42)
	whoAmI("hello")
	whoAmI(true)
}
