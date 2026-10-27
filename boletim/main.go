package main

import (
	"fmt"
)

func notas() []float64 {
	notas := []float64{7.5, 4.0, 9.2, 6.0, 3.5, 8.8, 5.9}
	return notas
}

func media(notas []float64) (float64, int){
	somatotal := 0.0
	voltas := len(notas)
	for _, nota := range notas {
		somatotal += nota
	}
	calculomedia := somatotal / float64(voltas)
	return calculomedia, voltas
}

func MaioreMenor(notas []float64) (float64, float64) {
	maior := 0.0
	menor := 10.0
	for _, nota := range notas {
		if nota >= maior {
			maior = nota
		}
		if nota <= menor {
			menor = nota
		}
	}
	return maior, menor
}

func situacao(nota float64) string {

	switch {
	case nota >= 7:
		return "APROVADO"
	case nota < 7 && nota >= 5:
		return "RECUPERAÇÃO"
	case nota < 5:
		return "REPROVADO"

	}
	return "aconteceu algum erro"
}

func contador(notas []float64) int {
	
	aprovados := 0
	for _,nota := range notas {
		if nota >= 7 {
			aprovados += 1
		} 
	}
	return aprovados
}


func main() {
	media, voltas := (media(notas()))
	fmt.Printf("A média geral da turma é %.2f\n", media)

	maior, menor := (MaioreMenor(notas()))
	fmt.Printf("A maior nota é : %.1f | A nota menor é: %.1f\n", maior, menor)

	for i, nota := range notas() {
		situacao := situacao(nota)
		i += 1

		fmt.Printf("ALUNO %d: %.2f - %s\n", i, nota, situacao)
	}

	contadorAprovados := contador(notas())
	fmt.Printf("O número de aprovados é %d de %d\n", contadorAprovados, voltas)
}
