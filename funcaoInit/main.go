package main

import "fmt"

var numeroImportante int
//funcao init significa que é uma funcao que sera executada antes da funcao main
//uso muito recomendado para qunado se cria uma variavel global e vc tem que setar ela dentro da funcao init
func init(){
	fmt.Println("primeiro")
	numeroImportante = 2025
}


func main(){
	fmt.Println("segundo")
	fmt.Println(numeroImportante)
}