package main

import(
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Hello Venus.")
}
	
func fizzbuzz(num int) string {
		var name string = ""
		if num<0{
		fmt.Println("Must be Positive Number")
		return "Must be Positive Number"
		}
		if num%15 == 0 {
			fmt.Println("Fizzbuzz")
			return "Fizzbuzz"
		}
		if num%3 == 0{
			fmt.Println("Fizz")
			return "Fizz"
		}
		if num%5 == 0{
			fmt.Println("Buzz")
			return "Buzz"
		}
		if name == "" {
			name = strconv.Itoa(num)
		}
		fmt.Println(name)
		return name	
}


