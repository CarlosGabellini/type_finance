package ctr

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
	"time"
	"type_finance/spt"
)

//Esse arquivo serve de suporte para o outro arquivo fc-finance.go, onde as funcoes dos
// switchs ficam exatamente aqui;

func MostrarAnosAnteriores() {
	ARQUIVOS_DISPONIVEIS, err := spt.ListarArquivos()

	Formatacao()
	fmt.Printf("\n")
	formatado_tabela := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Fprintln(formatado_tabela, "Arquivo\t|Diretorio")
	fmt.Fprintln(formatado_tabela, "---------\t---------")

	for _, caminho := range ARQUIVOS_DISPONIVEIS {
		fmt.Fprintf(formatado_tabela, "%s\t%s\n", filepath.Base(caminho), filepath.Dir(caminho))
	}

	formatado_tabela.Flush()

	fmt.Printf("\nPressione enter para sair...")
	LerLinha()
}

func CadastrarNovoMes(MinhasFin *spt.Arquivos, ano int) {

	var proximo time.Time

	if len(MinhasFin.MeuCaderno) == 0 {
		proximo = time.Date(ano, time.January, 1, 0, 0, 0, 0, time.Local)

	} else {
		// Pega o último caderno e avança 1 mês a partir do início dele
		ultimo := MinhasFin.MeuCaderno[len(MinhasFin.MeuCaderno) - 1]

		if ultimo.DataInicio.Month() == time.December {
			fmt.Println("Todos os meses ja estao cadastrados!")
			return
		}

		proximo = time.Date(ultimo.DataInicio.Year(), ultimo.DataInicio.Month() + 1, 1, 0, 0, 0, 0, time.Local)
	}

	fim := proximo.AddDate(0, 1, -1)

	indice, err := MinhasFin.NovoCaderno(proximo, fim)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("Mês %s cadastrado (índice %d)\n", proximo.Format("01/2006"), indice)
} 