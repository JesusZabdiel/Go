package main

import(
	"fmt"
)

type User struct{
	age int
	name, familyName string
}

func main(){


	//1st way of calling
	zabdiel := User{age: 20, name: "Zabdiel", familyName: "Sanchez"}
	fmt.Println(zabdiel)
	//2nd way

	jesus := User{20,"Jesus", "Sanchez"} //it is necesary to specify all information by order
	fmt.Println(jesus)

	//3rd way

	jesusZ := new(User)

	jesusZ.name = "jesusZ"

	fmt.Println(jesusZ.name)

}