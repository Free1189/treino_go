package main

import (
	"fmt"
)

type Produto struct{
	Nome string
	Preco float64
	Estoque int
}
type Cliente struct{
	ID int
	Nome string
	Idade int
	Email string
}
type Pedido struct{
	ID int
	Cliente *Cliente
	Produto *Produto
	Quantidade int
	Precototal float64
}


func main (){ 
	fmt.Println("Esse sera uma teste de conhecimento para um pedido")


}