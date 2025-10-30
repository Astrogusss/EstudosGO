package main

import "fmt"

func inverterSinal (numero int) (int){
	return numero * -1
}

func inverteSinalPonteiro (numero *int) (){
	*numero = (*numero) * (-1)
}

func main(){
	numero := 20
	// invertido := inverterSinal(numero)
	inverteSinalPonteiro(&numero)

	//mesma merda de c para referenciar numeros com ponteiro
	fmt.Println(numero)

}