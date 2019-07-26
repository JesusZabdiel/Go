package main

import(
	"fmt"
)

func main(){
	/*
	1.- A pointer is a memory adress
	2.- When we acces a pointer we are accesing to the adress, not the value
	3.- If X,Y in "0x51d52d"
	4.- Then  X = 6
	5.- "Y" value will be 6

	*T => *int, *string, *Struct

	Zero value => nill
	*/

	var x,y *int

	entero := 5
	x = &entero
	y = &entero

	*x = 10  //*pointerName shows us the value in that specific memory adress

	

	fmt.Println(x,y)
	fmt.Println(*x,*y)


}