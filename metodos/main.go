package main

import "fmt"

//um principio do poo, falando que todas as intancias de usuario pode invocar uma funcao
type usuario struct{
	nome string
	idade uint8
}

//o "u" pode ser qualquer coisa, é somente uma variavel que pode referenciar os campos
func (u usuario) salvar(){
	fmt.Println(u.nome)
}

//podemos tambem alterar os campos da struct, só que temos que passar como ponteiro
//aqui ele vai alterar dentro do endereco de memoria
func (u *usuario) fazerAniversario(){
	u.idade++
}

func (u usuario) getIdade(){
	fmt.Println(u.idade)
}

func main(){
	gustavo := usuario{"gustavo", 21}
	// fmt.Println(gustavo)
	gustavo.getIdade()
	gustavo.fazerAniversario()
	gustavo.getIdade()


}