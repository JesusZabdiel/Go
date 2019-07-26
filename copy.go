package main

import (
	"fmt"
)

func main(){
	 
	slice := make([]int, 5,10)
	slice2 := make([]int,len(slice), cap(slice)+1)

	copy(slice2, slice) //func copy copy(destino, fuente)

	fmt.Println(slice)
	fmt.Println(slice2)



}
