package main

import "fmt"

func main() {
	
	var number int
	number = 10
	if (number % 2 == 0){

		for x := 100; x >=1; x--{
			fmt.Print(x , " ")
		}
	}else{
		for x:= 0; x <= 100; x++{
			fmt.Print(x , " ")
		}
	}

}