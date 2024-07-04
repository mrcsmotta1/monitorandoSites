package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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
		fmt.Println("")

		return comandoInt
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":")
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}
