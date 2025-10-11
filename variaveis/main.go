package main

import (
	"errors"
	"fmt"
)

// import "fmt"

func main() {
	//voce pode declarar as variaveis assim
	// var fraseFodas string = "peidei legal"

	//tambem pode ser implicito (mas ainda ele tem o tipo)
	// frasesFodas := "peidei legal"

	//outro jeito de declarar mas tudo de uma vez so
	// var(
	// 	frasesFodas string
	// 	numerosFodas int
	// )
	// frasesFodas = "peidei legal     "
	// numerosFodas = 12
	// fmt.Println(frasesFodas, numerosFodas)

	//outro jeito de declarar mas com inferencia varias vezes
	// frasesFodas, numerosFodas := "peidinho", 144
	// fmt.Println(frasesFodas, numerosFodas)

	//////////////////////////////////////////////////////////////

	// para numeros temos somente os inteiros e os reais (int e float)

	//tipos de inteiros int8, int16, int32, int64 ---> mas pode ser declarado como somente int, baseado na arq da cpu
	//pode ter tambem os unisgned int ---> (uint)
	//tipos de reais float32, float64

	//char normalmente int para representa

	char := 'b'
	fmt.Println(char) //ele solta o numero na tabela ascii

	//toda variavel quando declarada, ela recebe zero/nulo como padrao

	var exemplo int
	fmt.Println(exemplo)

	//tipo booleano, so pode ser true ou false

	//tem tambem o tipo error

	//erro nao é do tipo string, é do tipo error mesmo, temos que usar um biblioteca para especificar o erro
	var erro error = errors.New("erro dos brarbos")
	fmt.Println(erro)
}
