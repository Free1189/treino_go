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
	ID         int
	Cliente    *Cliente
	Produto    *Produto
	Quantidade int
	Precototal float64
}
var novoCliente = []*Cliente{}

var leitor = bufio.NewReader(os.Stdin)

func LerFloat(mensagem string) (float64,error){
	fmt.Println(mensagem)
	texto,_ := leitor.ReadString('\n')
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
 
func LerInteiro(valor string) (int,error){
	fmt.Println(valor) 
	texto,_ := leitor.ReadString('\n')
	texto = strings.TrimSpace(texto)
	numero,erro := strconv.Atoi(texto)
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

func LerEmail (email string) bool {
	fmt.Println(email)
	leitura := strings.Contains(email, "@")
	
	return leitura 
	
}

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
		
	}else {
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
	novoCliente = append(novoCliente, &informacoesCliente)
	return

}


	

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
