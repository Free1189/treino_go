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

func LerString(promt string) (string, erro) {
	if promt == ""{
		return fmt.Println("Nome inválisdo, digite o nome novamnete")
	}

	return promt  
}

func cadastrarCliente() {

}

func cadastrarProduto(){

}
func cadastrarPedido() {

}
func listarPedidos(){

}
func listarClientes(){

}
func listarProdutos(){

}


func main (){ 
	fmt.Println("Esse sera uma teste de conhecimento para um pedido")

	for {
		fmt.Println("========================")
		fmt.Println("1 - Cadastrar cliente")
		fmt.Println("2 - Cadastrar produto")
		fmt.Println("3 - Cadastrar pedido")
		fmt.Println("4 - Listar Pedidos")
		fmt.Println("5 - Listar Clientes")
		fmt.Println("6 - Listar Produtos")
		fmt.Println("0 - Sair")
		fmt.Println("DIGITE O QUE VOCÊ QUER FAZER")
		escolha := fmt.Scnaln(escolha)

		switch escolha{
		case 1: cadastrarCliente(fmt.Scanln("Qual o nome do cliente : "))
		case 2: cadastrarProduto(fmt.Scanln("Qual o nome do Produto : "))
		case 3: cadastrarPedido(fmt.Scanln("Qual o nome do Cliente para fazer o peido : "))
		case 4: listarPedidos()
		case 5: listarClientes()
		case 6: listarProdutos()
		case 7: return 
		}

		




	}


}
