package main

import "fmt"

func closure() func(){
	texto := "dentro da funcao closure"

	funcao := func(){
		fmt.Println(texto)
	}

	return funcao
}


// referenciam variaveis que estao fora do seu corpo	
func main() {
	//repare que criei uma variavel de mesmo nome que esta dentro da funcao closure
	texto := "dentro da funcao main"
	fmt.Println(texto)

	//aqui ele recebe o que a funcao closure recebe, ou seja, uma nova outra funcao
	funcaonova := closure()

	//sera que ele vai printar o texto da main ou o texto da funcao closure?
	funcaonova()
	//por mais que a mesma variavel tenha o mesmo nome que esta na funcao, ele ainda vai referenciar o que esta dentro da funcao que ela foi criada
	//nao se perde a referencaia 
}
