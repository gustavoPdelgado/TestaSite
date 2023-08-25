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

const qtdMonitoramento = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			monitoramento()
			fmt.Println()
		case 2:
			fmt.Println("Exibindo logs...")
			fmt.Println()
			imprimeRegistros()
			fmt.Println()
		case 0:
			fmt.Println("Saindo.")
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido, o programa sera fechado.")
			fmt.Println()
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Gustavo"
	var versao float32 = 1.1

	fmt.Println("Olá! Sr.", nome)
	fmt.Println("A versão do programa é:", versao)

}

func exibeMenu() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func monitoramento() {
	fmt.Println("Monitorando...")
	//urls := []string{"https://www.recmais.com.br", "https://www.alura.com.br", "https://random-status-code.herokuapp.com/"}
	urls := leArquivo()
	for i := 0; i < qtdMonitoramento; i++ {
		for i, url := range urls {
			fmt.Println("Testando site", i, ":", url)
			testaSite(url)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func testaSite(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		os.Exit(-1)
	} else {

		if resp.StatusCode == 200 {
			fmt.Println("O site:", url, "foi carregado com sucesso! -> StatusCode:", resp.StatusCode)
			escreveLogArquivo(url, true)
		} else {
			fmt.Println("O site", url, "não está funcionando! -> StatusCode:", resp.StatusCode)
			escreveLogArquivo(url, false)
		}
	}
}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("links.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
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

func escreveLogArquivo(site string, status bool) {

	arquivo, err := os.OpenFile("Registros.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("O-correu um erro:", err)
		os.Exit(-1)
	}

	arquivo.WriteString(time.Now().Format("02/Jan/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	fmt.Println(arquivo)

	arquivo.Close()
}

func imprimeRegistros() {

	arquivo, err := os.ReadFile("Registros.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
	}
	fmt.Println(string(arquivo))
}
