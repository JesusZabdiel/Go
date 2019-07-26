package main 

import(
	"fmt"
	"io/ioutil"
)

func main(){

	data, err := ioutil.ReadFile("./readable.txt")
	
	if err != nil{
		fmt.Println("An error ocurred")

	}

	fmt.Println(string(data))

}