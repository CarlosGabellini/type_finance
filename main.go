package main

import (
	"fmt"
	"type_finance/ctr"
	"type_finance/spt"
)

/*-------------------------------------------- Sobre a funcao Main --------------------------------------/
	A funcao main eh a logica do arquivo e onde vai funcionar uma boa parte do programa, aqui estou
comecando a fazer o basico sobre a impressao de determinadas mensagens para que o nosso programa com-
-pile corretamente.

	MeuArquivo = Variavel global onde mecheremos nela para ver sobre nossos arquivos;
---------------------------------------------------------------------------------------------------------/
*/

var MeuArquivo spt.Arquivos

func main() {
	
	for i := 0; i < 50; i++ {
		fmt.Printf("=")
	}

	for {
		fmt.Printf("\n")
		fmt.Printf("Seja bem vindo ao seu controle financeiro!\n")
	
		fmt.Printf("Quais opcoes deseja usar para cadastro?\n\n")
		fmt.Printf("00 - Abortar operacao.\n")
		fmt.Printf("01 - Cadastrar novo ano.\n")	//Deve somente abrir e cadastrar um ano, e salvar no disco;
		fmt.Printf("02 - Abrir ano\n")
	
		fmt.Printf("\n")
		
		var controle int

		fmt.Printf("Digite sua opcao aqui - ")
		fmt.Scan(&controle)

		switch controle {

			case 0:
				fmt.Printf("Abortado!")
				return

			case 1:
				ctr.CadastrarAno(&MeuArquivo)

			default:
				ctr.Formatacao()
				fmt.Printf("Opcao invalida! Tente novamente!")
		}

		ctr.Formatacao()			//Limpa o terminal antes de comecar tudo denovo!
	}
}