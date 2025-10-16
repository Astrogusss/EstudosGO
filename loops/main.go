package main

import "fmt"

func main(){
	i := 0

	// a maneira de fazer while
	for i < 10{
		i++
	}

	//maneira de fazer um for do jeito c de ser
	// for j := 0; j < 10; j++{
	// 	fmt.Println(j)
	// }
	//o J nao fica visivel fora do for

	//agora aprendendo fazer um "for in" do jeito python de ser

	// cidades := [3]string{"aiuruoca", "caxambu", "juiz de fora"}

	//sempre vai devolver dois valores, a posicao aonde esta e o valor da posicao aonde esta
	//primeiro parametro por padrao, é o indice
	// for indice, valor := range(cidades){
	// 	fmt.Println(indice, valor)
	// }

	//quando nao se que ro indece, mas so o valor, fazemos isso
	// for _, valor := range(cidades){
	// 	fmt.Println(valor)
	// }

	//tambem podemos iterar por letras
	for indice, valor := range "AIURUOCA"{
		fmt.Println(indice, string(valor))
	}
	//tem que colocar string se nao, ele solta numeros da tabela ascii

	//para map, é um pouco diferente

	usuario := map[string]string{
		"nome": "gustavao Gostoso",
		"idade": "17",
	}

	for chave, valor := range(usuario){
		fmt.Println(chave, valor)
	}
}