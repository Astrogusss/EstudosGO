package pacoteauxiliar

import "fmt"

//quando a funcao tem em sua primeira letra, sendo maiuscula, quer dizer que ele pode ser visivel em outras parte
//qunado se tem uma funcao que comeca com e letra minuscula, ela só e visivel dentro do pacote aonde ela esta

//funcao que fala bom dia
func BomDia(){
	fmt.Println("Bom dia")
	estouComCaganeira()
}