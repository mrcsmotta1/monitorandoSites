package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

// restante do código omitido

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("./sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(1)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println("Ocorreu um erro ao criar arquivo: ", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
