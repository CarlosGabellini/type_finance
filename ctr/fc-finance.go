package ctr

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"type_finance/spt"
)

func Saudacao() {
	
	for i := 0; i < 50; i++ {
		fmt.Printf("=")
	}

	fmt.Printf("\n")
	fmt.Printf("Seja bem vindo ao seu controle financeiro!\n")

	fmt.Printf("Quais opcoes deseja usar para cadastro?\n")
	fmt.Printf("00 - Abortar operacao.\n")
	fmt.Printf("01 - Cadastrar novo ano.\n")
	fmt.Printf("02 - Abrir ano\n")
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

	fmt.Printf("\n\n")
}

func MenuPrincipal(MinhasFinancas *spt.MeuCaderno) {

	for {
		Saudacao()
		var controle int

		fmt.Scan(&controle)

		switch controle {

			case 0:
				fmt.Printf("Abortado!")
				return
			
			case 1:
				CadastrarAno(MinhasFinancas)

			default:
				Formatacao()
				fmt.Printf("Opcao invalida!! Tente novamente!\n")
		}
	}
}

func CadastrarAno(MinhasFinancas *spt.MeuCaderno) {

	var Controle int
	var AnoInput int

	Formatacao()

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
	}

	//Colocando um time somente para aparecer a mensagem de sucesso!
	time.Sleep(2 * time.Second)
}