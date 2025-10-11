package main

import "fmt"

func main(){
	//como que é feito a alocacao dinamica
	//primeiro argumento é o tipo do que voce quer alocar
	//segundo é o tamanho do array
	// terceiro é a capacidade maxima do array
	array := make([]int, 10, 11)
	fmt.Println("quantidade que tem no vetor antes", len(array))
	fmt.Println("capacidade que tem no vetor antes", cap(array))

	//quando usamos slice, temos que o o vetor criado é de 15 posicoes, só que a fatia que 
	//teremos pra ver é as 10 primeiras posicoes desse vetor que criamos
	array = append(array, 12)
	array = append(array, 13)
	fmt.Println("quantidade que tem no vetor depois", len(array))
	fmt.Println("capacidade que tem no vetor depois", cap(array))
	//se analisar a capacidade, ele dobra de tamanho quando se estoura a capacidade do vetor


	
}
