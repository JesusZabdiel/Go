package main

import(
	"fmt"
)

type User struct{
	age int
	name, familyName string
}

func (this User) nameFormat() string{          
	return this.name + " " + this.familyName
}

/*When we use objects with fucntions we are only creating a copy of the object
If we want to modify it, wee need to pass the pointer
Example:

func (this *User) SetName(s string){
	this.name = s
}


func main(){
	jesus := new(User)
	jesus.name = "Jesus"

	jesus.SetName("Zabdiel")

	fmt.Println(jesus.name) -----> Zabdiel}

*/

func main(){
	jesusZ := new(User)
	jesusZ.name = "Jesus"
	jesusZ.familyName = "Sanchez"

	fmt.Println(jesusZ.nameFormat())

}
