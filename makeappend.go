package main

import(
	"fmt"
)


func main(){

	slice := make([]int, 5, 10) //make([]type, len, cap)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice,5)
	fmt.Println(slice)
}
	