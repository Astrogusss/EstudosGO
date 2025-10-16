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

func media(n1, n2 int) (bool){
	defer recurarPanic()
	media := (n1 + n2) / 2

	if media < 6{
		return true
	}else if media > 6{
		return false
	}

	//logo apos o panic, não será executado nada, o sistema morrerá a partir do panic
	panic("exatemente 6")
}

func recurarPanic(){
	//recura o que foi interrompido pelo panic
	if r := recover(); r != nil{
		fmt.Println("vai tomar no cu piranha")
	}
}

func main() {
	// soma, subtracao := matematica(10, 20)
	// fmt.Println(soma, subtracao)

	//----------------------------------------
	
	// fmt.Println(soma(12, 10, 14, 16, 17))

	//----------------------------------------

	//funcao anonima
	// func(){
	// 	fmt.Println("sdnfsdfds")
	// }()

	// //ou

	// func(nome string){
	// 	fmt.Println(nome)
	// }("gustavo da massa")

	//----------------------------------------

	//funcao defer, quer dizer que ela sera a ultima executada na hora de ver as funcoes

	//----------------------------------------
	media(6,6)

}
