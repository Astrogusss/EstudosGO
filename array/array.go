package main

import (
	"fmt"
)

func main() {

	//jeito comum de declarar vetores
	var vetorzao [6]int
	vetorzao[0] = 12
	fmt.Println(vetorzao)

	//jeito de declarar implicitamente
	vetorzao1 := [6]string{"vai tomar no cu", "vai se foder", "sdlgd"}
	fmt.Println(vetorzao1)

	//jeito mais foda de declarar um vetor "mais dinaimico"
	//nao precisa especificar o tamanho do vetor, to povoando o vetor, depois ele da o tamanho do vetor
	vetorzao2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(vetorzao2)

	//mas vetor nao é muito usual, melhor fazer um slice
	//de tamanho mais dinamico
	vetorzaoSlice := []int{10, 10, 12}
	vetorzaoSlice = append(vetorzaoSlice, 13)
	fmt.Println(vetorzaoSlice)

	//isso é um ponteiro, entao o vetor original nao pode ser alterado, porque muda o que esta apontando pra ele tambem
	vetorzaoSlice1 := vetorzao1[1:3]
	fmt.Println(vetorzaoSlice1)

}
