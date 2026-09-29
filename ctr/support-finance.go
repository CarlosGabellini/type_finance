package ctr

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
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