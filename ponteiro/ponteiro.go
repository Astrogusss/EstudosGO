package main

import "fmt"

func main(){
	//mesma merda de c
	var valor1 int = 10
	var valoreferenia *int = &valor1
	fmt.Println(*valoreferenia)
}
