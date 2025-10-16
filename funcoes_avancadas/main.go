package main

import "fmt"

// funcoes nomeadas
func matematica(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

//funcoes variaticas
func soma(numeros ...int) (int){
	//ele entra aqui como um slice
	fmt.Println(numeros)
	var resultado int

	for _, numeros := range(numeros){
		resultado += numeros
	}
	return resultado
}

func main() {
	// soma, subtracao := matematica(10, 20)
	// fmt.Println(soma, subtracao)

	//----------------------------------------
	
	// fmt.Println(soma(12, 10, 14, 16, 17))

	//----------------------------------------

	//funcao anonima
	func(){
		fmt.Println("sdnfsdfds")
	}()

	//ou

	func(nome string){
		fmt.Println(nome)
	}("gustavo da massa")

}
