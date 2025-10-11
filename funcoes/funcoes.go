package main

//assim que funciona a assinatura da funcao --> nome da variavel -- tipo da variavel
//logo quando se fecha os parenteses, temos que colocar o que ele retorna

func soma(n1 int8, n2 int8) int {
	return int(n1 + n2)
}

// funcoes pode ter retorno de duas variaveis
// assim que pode ser declarada
func retornaDuasCoisas(n1 int, n2 int) (int, int) {
	soma := soma(int8(n1), int8(n2))
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	// fmt.Println(soma(3,4))

	//pode tambem colocar dentro de uma variavel um funcao, seria uma variavel do tipo func
	//pode passar parametros tambem tudo que a de direito
	// var aux = func(n1 int8, n2 int8) int8 {
	// 	return n1 + n2
	// }
	// fmt.Println(aux(3,6))

	//para pegar as duas variaveis que retornou, basta declarar duas variaveis e chamar elas
	//tem que colocar duas variaveis para retornar aquilo que vem da funcao
	// resultadoSoma, resultadoSubtracao := retornaDuasCoisas(9, 7)
	// fmt.Println(resultadoSoma, resultadoSubtracao)

	//podemos tambem nao querer um dos parametros que a funcao nos retorna, ai basta usar um underline para especificar que vc nao quer
	// resultadoSoma, _ := retornaDuasCoisas(6,8)
	// fmt.Println(resultadoSoma)
}
