package main

import "fmt"

type pessoa struct{
	nome string
	idade uint8
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main(){

	//mesmo jeito de C para declarar uma variavel do tipo struct etc
	// var gustavo usuario
	// gustavo.idade = 15
	// gustavo.nome = "gustavao da massa"
	// fmt.Println(gustavo.idade, gustavo.nome)

	//podemos tambem ter inferencias de tipos para struct
	// cinara := usuario{"cinara vida loka", 47}
	//fmt.Println(cinara)

	//tem tambem de como tipar a entrada de dados da struct
	// edvaldo := usuario{nome: "edvaldo"}
	// fmt.Println(edvaldo)

	//heranca em go
	gustavo := estudante{pessoa{"gustavo", 19}, "ciencia da computacao", "unifei"}
	fmt.Println(gustavo)
}
