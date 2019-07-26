package main 

import(
"fmt"
"time"
"strings"

)

func main(){
	go SlowName("Jesus")
	fmt.Println("Finally!")
	var wait string
	fmt.Scanln(&wait)
}

func SlowName(name string){

	letras := strings.Split(name, "")

	for _, letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}



}