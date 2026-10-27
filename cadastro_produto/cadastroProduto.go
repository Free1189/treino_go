package main

import (
	"fmt"
	
)

type Produto struct{
	Nome string
	Preco float64
	Quantidade int
}

var produtos = []Produto{}



func cadastroProduto() Produto {
	var nome string
	var preco float64
	var quantidade int 

	fmt.Print("Qual o nome do seu produto: ")
	fmt.Scanln(&nome)

	fmt.Print("Qual o preço do novo produto: ")
	fmt.Scanln(&preco)

	fmt.Print("Qual a quantidade desse produto: ")
	fmt.Scanln(&quantidade)

	produto := Produto{
		Nome: nome,
		Preco: preco,
		Quantidade: quantidade,
	}
	
	return produto

}


func (p Produto) Deletar(nome string){
	for i, p := range produtos {
		if nome == p.Nome {
			produtos = append(produtos[:i], produtos[i+1:] ...)
			fmt.Println(produtos)
			
			break
		
		}

	}
}

func (p *Produto) adicionarEstoque(estoque int) {
	p.Quantidade += estoque

}


  

func main() {

	produto := Produto{
			Nome: "notebook",
			Preco: 5000,
			Quantidade: 30,
			
	}

	produtos = append(produtos, produto)

	
	var escolha int
	fmt.Println("Qual ação você deseja: ")
	fmt.Println("=======================")
	fmt.Println("1 - Cadastrar Produto ")
	fmt.Println("2 - Listar Produtos ")
	fmt.Println("3 - Adicionar estoque ")
	fmt.Println("4 - Remover Produtos ")
	fmt.Println("0 - Sair ")
	fmt.Print("Qual ação você deseja: ")
	fmt.Scanln(&escolha)

	if escolha == 1 {
		produto := cadastroProduto()
		produtos = append(produtos, produto)
		
	}
	if escolha == 2 {
		fmt.Println(produtos)

	}
	if escolha == 3 {

		var estoque int 
		fmt.Print("Quantos itens quer adicionar no estoque: ")
		fmt.Scanln(&estoque)
		produto.adicionarEstoque(estoque)
		fmt.Println(produto.Quantidade)
		
	}
	
	if escolha == 4 {
		var deletar string

		fmt.Print("Qual o nome do item de deseja excluir: ")
		fmt.Scanln(&deletar)

		produto.Deletar(deletar)

		
	}

	if escolha == 0{
		fmt.Println("O sistema foi encerrado.")
		
	}

}