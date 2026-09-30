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

func PersonalizarTerminal() {

	for i := 0; i < 50; i++ {
		fmt.Printf("-")
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
		CAMINHO_ARQ, err := MinhasFinancas.SalvarAno()
	
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

		for i := 0; i < 35; i++ {

			fmt.Printf("-")
			
			if i == 16 {
				fmt.Printf(" Menu dos anos ")
			}
		}

		fmt.Printf("\n")
		
		fmt.Printf("Selecione uma das opcoes abaixo: \n")
		fmt.Printf("00 - Abortar Operacao\n")
		fmt.Printf("01 - Mostrar arquivo de anos anteriores\n")
		fmt.Printf("02 - Abrir algum ano.\n")
		fmt.Printf("03 - Salvar o ano.\n")

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

			case 2:
				AbrirAlgumAno(MinhasFinancas)
				time.Sleep(3 * time.Second)

			case 3:
				CAMINHO_ARQ, err2 := MinhasFinancas.SalvarAno()

				if err2 != nil {
					fmt.Println(err2)
					time.Sleep(3 * time.Second)
					return
				}

				fmt.Println("Ano esta salvo! Caminho dele: ", CAMINHO_ARQ)
				time.Sleep(3 * time.Second)
				
			default:
				fmt.Printf("Numero invalido!")
		}

		Formatacao()
	}
}


func AbrirAlgumAno(MinhasFin *spt.Arquivos) {
	Formatacao()

	fmt.Printf("\n")
	fmt.Printf("Digite o ano que gostaria de ver.\n")

	for i := 0; i < 50; i++ {
		fmt.Printf("-")
	}

	fmt.Printf("\n")
	fmt.Printf("Digite sua opcao aqui - ")
	controle, err := strconv.Atoi(LerLinha())

	if err != nil {
		fmt.Println(err)
		return
	}

	err = MinhasFin.CarregarArquivo(controle)

	if err != nil {
		fmt.Println(err)
		return				//Aqui deve ser return por que o ano nao foi encontrado!
	}

	fmt.Printf("Ano carregado com sucesso!")
	
	MenuDoAno(MinhasFin)
}

func MenuDoAno(MinhasFin *spt.Arquivos) {
	Formatacao()

	for {
		fmt.Printf("\n")
		fmt.Printf("Bem vindo! O que deseja fazer?\n")

		fmt.Printf("0 - Abortar\n")
		fmt.Printf("01 - Cadastrar novo mes\n")
		fmt.Printf("02 - Abrir novo mes\n")

		for i := 0; i < 50; i++ {
			fmt.Printf("-")
		}

		fmt.Printf("\n")
		fmt.Printf("Digite sua opcao aqui - ")

		controle, err := strconv.Atoi(LerLinha())

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch controle {

			case 0:
				fmt.Printf("Abortando....")
				return
		
			case 1:
				CadastrarNovoMes(MinhasFin, MinhasFin.Ano)
				time.Sleep(2 * time.Second)
		}
	}
}