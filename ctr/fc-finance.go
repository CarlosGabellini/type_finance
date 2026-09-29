package ctr

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"type_finance/spt"
)

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
		fmt.Scan(&AnoInput)
		
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