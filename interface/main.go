package main

import (
	"fmt"
	"math"
)

type quadrado struct{
	lado int
}

type retangulo struct{
	altura int
	largura int
}

//na interface so tem assinatura das funcoes que queremos "abstrair"
type forma interface{
	area() int
}

//essa é a funcao tem que ter o mesmo nome da interface dentro
//e falar o que o metodo tem que fazer
func (r retangulo) area() int{
	return r.altura * r.largura
}

func (q quadrado) area() int{
	return int(math.Pow(float64(q.lado), 2))
}
func escreverArea(f forma){
	fmt.Printf("a área será de %d \n", f.area())
}

func main(){
	experimentacao := retangulo{10, 20}
	experimentacao1 := quadrado{10}

	escreverArea(experimentacao)
	escreverArea(experimentacao1)

}