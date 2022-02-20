package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle Switch")
	dia := diaDaSemana(1)
	fmt.Println(dia)

	fmt.Println("---------")
	diaSegundaManeira := diaDaSemanaSegundaManeira(1)
	fmt.Println(diaSegundaManeira)
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia não encontrado"
	}
}

func diaDaSemanaSegundaManeira(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia não encontrado"
	}

	return diaDaSemana
}
