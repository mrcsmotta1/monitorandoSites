package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opção invalida, tente novamente!")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido string

	for {
		fmt.Scan(&comandoLido)
		comandoInt, err := strconv.Atoi(comandoLido)
		if err != nil {
			fmt.Println("O valor digitado não é um inteiro válido.")
			continue
		}

		if comandoInt < 0 || comandoInt > 9 {
			fmt.Println("O valor digitado está fora do intervalo (0-9).")
			continue
		}

		fmt.Println("O comando escolhido foi:", comandoLido)

		return comandoInt
	}
}
