package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

type Produto struct {
	Nome    string
	Preco   float64
	Estoque int
}
type Cliente struct {
	ID    int
	Nome  string
	Idade int
	Email string
}
type Pedido struct {
	ID         int
	Cliente    *Cliente
	Produto    *Produto
	Quantidade int
	Precototal float64
}

var leitor = bufio.NewReader(os.Stdin)

func LerString(script string) (string,error){
	fmt.Println(script)
	texto, erro := leitor.ReadString('\n') 
 	if erro != nil {
		return "", erro 
	}
	texto = strings.TrimSpace(texto)
	if texto == ""{
		return "", errors.New("O campo não pode ficar vazio")
	}
	return texto, nil 
}


func cadastrarCliente() {

}

func cadastrarProduto() {

}
func cadastrarPedido() {

}
func listarPedidos() {

}
func listarClientes() {

}
func listarProdutos() {

}

func main() {
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
		escolha := 0
		_, err := fmt.Scanln(&escolha)
		if err != nil {
			fmt.Println("número inserido é invalido:", err)
			continue

		}

		switch escolha {
		case 1:
			cadastrarCliente()
		case 2:
			cadastrarProduto()
		case 3:
			cadastrarPedido()
		case 4:
			listarPedidos()
		case 5:
			listarClientes()
		case 6:
			listarProdutos()
		case 0:
			return
		default:
			fmt.Println("Esse número é invalido tente outro")
			continue
		}

	}

}
