package main

import "fmt"

//funcao que retorna os dias da sema
func DiasDaSemana(numero int) string{
	switch numero{
	case 1:
		return "segunda-feira"
	case 2:
		return "terca-feira"
	case 3:
		return "quarta-feira"
	case 4:
		return "quinta-feira"
	case 5:
		return "sexta-feira"
	case 6:
		return "sabado"
	case 7:
		return "domingo"
	default:
		return "vai tomar no cu"
	}
}

func main(){
	fmt.Println(DiasDaSemana(10))
}