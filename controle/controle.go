package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Gustavo Senador de Faria"

	if strings.Compare(nome, "Guinho") == 0 {
		fmt.Println("você é feio")
	} else {
		fmt.Println("você é bonito")
	}

	//voce pode inicialiar um variavel dentro do if

	if OutroNome := "Guinho"; strings.Compare(OutroNome, "Guinho") == 0 {
		fmt.Println("Como você é feio muleke")
	} else {
		fmt.Println("como você é lindo")
	}

	//porem a variavel criada dentro do if_init, nao sera visivel fora do escopo do if

}
