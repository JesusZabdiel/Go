package main

import (

	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader (os.Stdin)
	fmt.Print("Teclea tu nomnbre: ")
	nombre,err := reader.ReadString('\n')

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Hello, ", nombre)
	}


	var edad int

	fmt.Print("¿Que edad tienes?: ")
	fmt.Scanf("%d", &edad)
	fmt.Print(nombre, "," , edad)
}