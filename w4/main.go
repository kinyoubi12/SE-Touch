package main

import (
	"fmt"
)

func cal (x int) (result int) {
	result = x + 2
	return result
}

func main(){
	fmt.Println("Hello Venus.")
	result := cal(2)
	fmt.Println(result)
}