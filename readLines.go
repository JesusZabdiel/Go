package main


import (
	"fmt"
	"bufio"
	"os"
)

func main (){

	file, err := os.Open("./readablee.txt")

	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan(){
		
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	file.Close()
}