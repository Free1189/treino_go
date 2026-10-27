package main
import "fmt"

type Conta struct{
	Numero int
	Saldo float64
}

func (c *Conta) adicionarSaldo(adicionar float64) {
	if adicionar <= 0 {
		return
	}
	c.Saldo += adicionar
}
func (c *Conta) sacar(sacar float64) {
	if sacar > c.Saldo || sacar <= 0 {
		return
	}
	c.Saldo -= sacar
}

func main() {
	
	
	conta := Conta{
		Numero: 23123,
		Saldo: 2000.00,
	}
	var adicionar float64
	var sacar float64
	var escolha int
	

	fmt.Println("=================")
	fmt.Println("1 - Adicionar saldo na conta")
	fmt.Println("2 - Sacar saldo da conta")
	fmt.Println("0 - Sair")
	fmt.Println("=================")
	fmt.Println("Faça sua escolha: ")
	fmt.Scanln(&escolha)

	if escolha == 1 {
		fmt.Println("Quanto deseja adicionar: ")
		fmt.Scanln(&adicionar)
		conta.adicionarSaldo(adicionar)
		fmt.Println(conta.Saldo)

	}
	if escolha == 2 {
		fmt.Println("Quanto você deseja sacar: ")
		fmt.Scanln(&sacar)
		conta.sacar(sacar)
		fmt.Println(conta.Saldo)
	}
	if escolha == 0 {
		return
	}


}