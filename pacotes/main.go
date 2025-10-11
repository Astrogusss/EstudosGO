package main

import (
	"fmt"
	//para importar modulos que estao dentro do seu projeto, temos que colocar o nome do modulo e a pasta aonde que esta essas funcoes
	"pacotes/pacoteAuxiliar"

	//para importar esses pacotes, temos que mandar um go get [NOME DO PACOTE IMPORTADO]
	"github.com/skip2/go-qrcode"
)

//quando voce tem que criar um projeto, temos que fazer o go mod init [NOME DO MODULO QUE VOCE DESEJA] -> como se fosse o nome do seu projeto
//quando damos um go build, fazemos que o programa crie um executavel possa ser executado pelo terminal, com o mesmo nome do MODULO que voce criou
func main() {
	fmt.Println("vai tomar no cu piranha")
	pacoteauxiliar.BomDia()

  	/*err := */ qrcode.WriteFile("https://discord.com/", qrcode.Medium, 256, "qr.png")

	// fmt.Println(err)
  }