package ctr

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"type_finance/spt"
)

var entrada = bufio.NewReader(os.Stdin)

func LerLinha() string {
	linha, _ := entrada.ReadString('\n')
	return strings.TrimSpace(linha)
}

func LimparTela() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")

	} else {							//Para Linux e MacOS caso prescise!
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Formatacao() {
	LimparTela()
	
	for i := 0; i < 50; i++ {
		fmt.Printf("=")
	}
}

func CadastrarAno(MinhasFinancas *spt.Arquivos) {

	var Controle int
	var AnoInput int

	Formatacao()
	fmt.Printf("\n")

	for {
		fmt.Printf("Digite o ano que deseja colocar: ")
		ano, err := strconv.Atoi(LerLinha())
		
		if err != nil {
			fmt.Println("Digite apenas numeros!")
			continue
		}
		AnoInput = ano
		Controle = spt.ColocarAno(AnoInput, MinhasFinancas)

		if Controle == 0 {
			fmt.Println("Ano invalido! tente novamente!")
		}

		if Controle == 1 {
			fmt.Println("Salvo com sucesso!")
			break
		}

		if Controle == 2 {
			fmt.Printf("Caminho JA EXISTE!")
			break
		}
	}

	if Controle == 1 {
		CAMINHO_ARQ, err := spt.SalvarArquivo(*MinhasFinancas)
	
		if err != nil {
			fmt.Printf("Nao foi possivel criar o arquivo!")
		}
		
		fmt.Printf("Ano criado! Caminho do arquivo: \n")
		fmt.Printf("%s", CAMINHO_ARQ)
	
		//Colocando um time somente para aparecer a mensagem de sucesso!
	}

	time.Sleep(2 * time.Second)
}

func AbrirAnos(MinhasFinancas *spt.Arquivos) {
	Formatacao()

	for {
		fmt.Printf("\n")
		
		fmt.Printf("Selecione uma das opcoes abaixo: \n")
		fmt.Printf("00 - Abortar Operacao\n")
		fmt.Printf("01 - Mostrar arquivo de anos anteriores\n")
		fmt.Printf("02 - Abrir algum ano.\n")
		fmt.Printf("03 - Imprimir tabela de algum ano && mes\n")

		for i := 0; i < 50; i++ {
			fmt.Printf("-")
		}

		fmt.Printf("\n")
		fmt.Printf("Digite sua opcao aqui - ")
		controle, err := strconv.Atoi(LerLinha())

		if err != nil {
			fmt.Printf("Digite apenas numeros!")
			continue
		}

		switch controle {

			case 0:
				fmt.Printf("Abortando operacao aqui!")
				time.Sleep(3 * time.Second)
				return
			
			case 1:
				MostrarAnosAnteriores()
				
			default:
				fmt.Printf("Numero invalido!")
		}

		Formatacao()
	}
}