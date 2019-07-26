package main

import(
	"fmt"
)

func main(){

	channel := make(chan string)

	go func (channel chan string){ //channel --> name of the channelthe function must wait
		for{
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel) //initialize function


	msg := <- channel
	fmt.Println("Se tecleo: " + msg)

	msg = <- channel
	fmt.Println("Se tecleo: " + msg)

}