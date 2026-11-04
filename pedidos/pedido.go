package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Produto struct {
	Nome    string
	Preco   float64
	Estoque int
}
type Cliente struct {
	Nome  string
	Idade int
	Email string
}
type Pedido struct {
	Cliente    *Cliente
	Produto    *Produto
	Quantidade int
	Precototal float64
}

var clientes = []*Cliente{}
var produtos = []*Produto{}
var pedidos = []*Pedido{}



var leitor = bufio.NewReader(os.Stdin)

func LerFloat(mensagem string) (float64, error) {
	fmt.Println(mensagem)
	texto, _ := leitor.ReadString('\n')
	texto = strings.TrimSpace(texto)
	texto = strings.Replace(texto, ",", ".", 2)
	numero, erro := strconv.ParseFloat(texto, 64)
	if erro != nil {
		fmt.Println("O valor possue erro, corrija e tente novamente ! ")
		return 0.0, erro

	}
	if numero <= 0.0 {
		fmt.Println("O valor deve ser maior do que zero ! Tente novamente. ")
		return 0.0, errors.New("Tente novamente, o numero esta menor do que zero ! ")
	}
	return numero, erro

}

func LerInteiro(valor string) (int, error) {
	fmt.Println(valor)
	texto, _ := leitor.ReadString('\n')
	texto = strings.TrimSpace(texto)
	numero, erro := strconv.Atoi(texto)
	if erro != nil {
		fmt.Println("valor inválido ! ")
		return 0, erro
	}
	if numero <= 0 {
		fmt.Println("o valor esta nulo, tente outro valor !")
		return 0, errors.New("o valor não pode ser nulo ou negativo")
	}

	return numero, erro

}

func LerEmail(email string) bool {
	fmt.Println(email)
	leitura := strings.Contains(email, "@")

	return leitura

}

func LerString(script string) (string, error) {
	fmt.Println(script)
	texto, erro := leitor.ReadString('\n')
	if erro != nil {
		return "", erro
	}
	texto = strings.TrimSpace(texto)
	if texto == "" {
		return "", errors.New("O campo não pode ficar vazio")
	}
	return texto, nil
}

func cadastrarCliente() {
	for {
		nome, err := LerString("Qual o nome do cliente : ")
		if err != nil {
			fmt.Println(err)
			continue

		}
		if nome != "" {
			fmt.Println("o nome foi cadastrado")

		}

		email, err := LerString("Qual o email do cleinte: ")
		if err != nil {
			fmt.Println("o email esta errado, digite novamente")
			continue
		}
		confereEmail := LerEmail(email)
		if confereEmail == true {
			fmt.Println("o email foi cadastrado ")

		} else {
			fmt.Println("o email possue erro valide")
			continue

		}
		idade, erro := LerInteiro("Qual a idade do cliente : ")

		if erro != nil {
			fmt.Println(erro)
			continue
		}
		if idade <= 14 {
			fmt.Println("Não possui idade o suficiente para fazer o cadastro na loja ! ")
			continue
		}
		if idade >= 15 {
			fmt.Println("idade cadastrada com sucesso")

		}
		informacoesCliente := Cliente{Nome: nome, Idade: idade, Email: email}
		clientes = append(clientes, &informacoesCliente)
		return

	}

}
func cadastrarProduto() {
	for {
		nomeProduto, err := LerString("Qual o nome do produto: ")
		if err != nil {
			fmt.Println("o nome do produto é inválido, tente novamente")
			continue
		}
		if nomeProduto == "" {
			fmt.Println("O nome do produto não pode ser vázio")
			continue
		} else {
			fmt.Println("nome cadastrado !")
		}

		preco, err := LerFloat("Qual o valor que será cobrado pelo produto: ")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if preco <= 0.0 {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("O valor foi cadastrado")
		}

		estoque, err := LerInteiro("Qual a quantidade desse produto no estoque: ")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if estoque <= 0 {
			fmt.Println("O valor do estoque não pode ser menor ou igual a zero! ")
			continue
		} else {
			fmt.Println("O valor foi cadastrado ! ")
			
		}
		informacoeProduto := Produto{Nome: nomeProduto, Preco: preco, Estoque: estoque}
		produtos = append(produtos, &informacoeProduto)
		fmt.Println("Produto cadastrado com sucesso")
		return
	}

}
func cadastrarPedido() {
	for{
	fmt.Println("Os clientes cadastrados são esses : ")
	listarClientes()
	clienteID, erro := LerInteiro("Qual o numero do cliente que fez o pedido: ")
	if erro != nil {
		fmt.Println(erro)
		continue
	}
	if clienteID > len(clientes) || clienteID < 0 {
		fmt.Println("O ID do cliente selecionado é inválido: ")
		continue
	} 
	
	IformacoesClientePedido := clientes[clienteID]
	fmt.Println("o ID foi cadastrado ! ") 
	


	fmt.Println("Os produtos cadastrados são esses : ")
	listarProdutos()
	produtoID,erro := LerInteiro("Qual o ID do produto que irá no pedido ?")
	if erro != nil{
		fmt.Println(erro)
		continue
	}
	if produtoID > len(produtos) || produtoID < 0 {
		fmt.Println("ID do produto é inválidom, verifique se o ID está na lista !")
		continue

	}
	InformacoesProdutoPedido := produtos[produtoID]
	fmt.Println("ID cadastrado")

	quantidade, erro := LerInteiro("Qual a quantidade do produto que irá no pedidos : ")
	if erro != nil {
		fmt.Println(erro)
		continue
	}
	if quantidade <= 0 || quantidade > InformacoesProdutoPedido.Estoque {
		fmt.Println("A quantiade selecionada é menor do que zero ou é maior do que estoque disponivel")
		continue
	}
	fmt.Println("A quantidade foi cadastrada")

	PrecoTotal := float64(quantidade) * InformacoesProdutoPedido.Preco
	fmt.Printf("Preço total = %.2f", PrecoTotal)
	todasIformacoesPedido := Pedido{Cliente: IformacoesClientePedido, Produto: InformacoesProdutoPedido, Quantidade: quantidade, Precototal: PrecoTotal}
	pedidos = append(pedidos, &todasIformacoesPedido)
	fmt.Println("Cadastro de pedido confirmado !")
	return

}

}
func listarPedidos() {
	if len(pedidos)> 0{
		for i, pedido := range pedidos {
			fmt.Printf("[%d] - Cliente = %s - Produto = %s - Quantidade = %d Preço total = %.2f",i , pedido.Cliente.Nome, pedido.Produto.Nome, pedido.Quantidade ,pedido.Precototal)
		}
	} 

}
func listarClientes() {
	if len(clientes) > 0 {
		for i, cliente := range clientes {

			fmt.Printf("[%d] - %s - %d anos - referente ao Email = %s\n", i, cliente.Nome, cliente.Idade, cliente.Email)
		}
	} else {
		fmt.Println("Para listar os clientes precisamos ter pelo menos um cliente cadastrado! ")
	}

}
func listarProdutos() {	if len(produtos) > 0{
		for i, produto := range produtos {
			fmt.Printf("[%d] - %s - Estoque = %d - Preco = %.2f\n", i, produto.Nome, produto.Estoque, produto.Preco)
		}
	}else {
		fmt.Println("Para listar os produtos deve ter apenas um produto cadastrado !")
	}

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
		fmt.Println("7 - Sair")
		escolha, err := LerInteiro("DIGITE QUAL O SEU OBJETIVO NO MENU ! ")
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
		case 7:
			return
		default:
			fmt.Println("Esse número é invalido tente outro")
			continue
		}

	}

}
