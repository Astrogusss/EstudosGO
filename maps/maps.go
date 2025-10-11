package main

import "fmt"

func main() {
	//assim que se declara um map
	//colchete é pra falar o tipo do campo
	//depois dos dois pontos, é o que tipo do argumento
	usuario := map[string]string{
		"nome":  "Gustavo",
		"idade": "21",
	}

	fmt.Println(usuario)

	//pode alinhas maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gustavo",
			"segundo":  "senador",
		},
	}

	fmt.Println(usuario2)
	//tem como deletar campos em maps
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//voce pode adicionar chave ao seu map
	usuario2["caganeira"] = map[string]string{
		"tipo":        "cocozao Gigante",
		"viscosidade": "espessa",
	}
	fmt.Println(usuario2)

}
