package main

import (
	"fmt"
)
type Human struct{
	name string

}


type Tutor struct{
	Human
}

func (this Human) speak() string{
	return "Bla bla bla"
}

func (this Tutor) speak() string{
	
	return this.Human.speak() + " Hello, I'm your tutor"
}
 

func main (){

	tutor := Tutor{Human{"Jesus"}}

	fmt.Println(tutor.speak())

}