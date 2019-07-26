package main

import (
	"fmt"
)

func main (){

	matriz := []int{1,2,3,4,} //when using slices, it is not necesary to determine size
	
	var matrizC []int  //when just initialize a slice, the zero value is 'nil'
	
	slice := matriz[0:3]

	fmt.Println(matriz)

	if matrizC == nil{
		fmt.Println("Nothing")
	}else{
		fmt.Print(matrizC)
	}

	fmt.Println(slice)
}

