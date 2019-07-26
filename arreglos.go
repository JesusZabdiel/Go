package main


import "fmt"

func main () {

	var arreglo [10]int
	var  matriz [5][2]int

	matriz[2][1] = 10

	arreglo2 := [3]int{1,2}
	fmt.Println(arreglo)
	fmt.Println(arreglo2)
	fmt.Println(matriz)
	fmt.Println(matriz[2][1])

	arreglo[5] = 50

	for i := 0; i < len(arreglo); i++ {

		fmt.Println(arreglo[i])
		
	}

}