package ctr

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"
	"type_finance/spt"
)

//Aqui eh para mostrar a tabela formatada das despesas, eh melhor usar uma copia do que a struct original;
func TabelaReceitasDespesas(MeuCaderno spt.MeuCaderno) {

	Formatacao()

	Mes := MeuCaderno.Mes
	var formatGasto string
	var formatPorcentagem string

	fmt.Printf("\nBem vindo a tabela de gastos && receitas!\n")
	fmt.Printf("Saldo inicial - %.2f\n", MeuCaderno.SaldoInicial)
	fmt.Printf("Valor referencia - %.2f\n", MeuCaderno.ValorRef)
	
	if MeuCaderno.SaldoFinal != 0 {
		fmt.Printf("Saldo final: %.2f\n", MeuCaderno.SaldoFinal)
	
	} else {
		fmt.Printf("Saldo final (ate o momento): %.2f\n", MeuCaderno.CalcularSaldoFinal())
	}
	
	if len(MeuCaderno.MeusGastos) <= 0 {
		fmt.Println("\nNenhuma transacao ainda!");
	}

	//Criando um slice para ordenar sem mecher no original;
	transacoes := make([]spt.MinhasTransacoes, len(MeuCaderno.MeusGastos))
	copy(transacoes, MeuCaderno.MeusGastos)

	sort.Slice(transacoes, func(i, j int) bool {

		//Desempate pelo ID, caso tenha duas transacoes;
		if transacoes[i].Data.Equal(transacoes[j].Data) {
				return transacoes[i].ID < transacoes[j].ID
		}
		
		return transacoes[i].Data.Before(transacoes[j].Data)
	})
	

	fmt.Printf("\n--------------------- Transacoes do mes %s ---------------------\n\n", Mes)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tDESCRICAO\tTIPO\tVALOR\tPORCENTAGEM\tDATA")
	fmt.Fprintln(w, "--\t---------\t----\t-----\t-----------\t----")

	for _, t := range transacoes {

		if t.Transacao == spt.Ganho {
			formatGasto = "+"
			formatPorcentagem = "---"
		}

		if t.Transacao == spt.Gasto {
			formatGasto = "-"
			formatPorcentagem = fmt.Sprintf("%.2f%%", t.Porcentagem)
		}
		
		fmt.Fprintf(w, "%02d\t%s\t%s\tR$ %.2f\t%s\t%s\n",
				t.ID,
				t.Descricao,
				formatGasto,
				t.Valor,
				formatPorcentagem,
				t.Data.Format("02/01"),		//Alterando o format para nao prescisar mostrar o ano toda vez;
		)
	}

	w.Flush()
	fmt.Println()

	fmt.Println("---------------------------------")
	fmt.Println("Digite enter para sair.....")
	LerLinha()
}

// Monta a tabela do mes em texto, sem mecher na funcao original;
func FormatarTabela(MeuCaderno spt.MeuCaderno, Ano int) string {

	var sb strings.Builder
	var formatGasto string
	var formatPorcentagem string

	fmt.Fprintf(&sb, "Relatorio de gastos && receitas - %s/%d\n\n", MeuCaderno.Mes, Ano)
	fmt.Fprintf(&sb, "Saldo inicial - %.2f\n", MeuCaderno.SaldoInicial)
	fmt.Fprintf(&sb, "Valor referencia - %.2f\n", MeuCaderno.ValorRef)

	if MeuCaderno.SaldoFinal != 0 {
		fmt.Fprintf(&sb, "Saldo final: %.2f\n", MeuCaderno.SaldoFinal)
		
	} else {
		fmt.Fprintf(&sb, "Saldo final (ate o momento): %.2f\n", MeuCaderno.CalcularSaldoFinal())
	}

	if len(MeuCaderno.MeusGastos) <= 0 {
		fmt.Fprintf(&sb, "\nNenhuma transacao ainda!\n")
	}

	transacoes := make([]spt.MinhasTransacoes, len(MeuCaderno.MeusGastos))
	copy(transacoes, MeuCaderno.MeusGastos)

	sort.Slice(transacoes, func(i, j int) bool {
		if transacoes[i].Data.Equal(transacoes[j].Data) {
			return transacoes[i].ID < transacoes[j].ID
		}
		return transacoes[i].Data.Before(transacoes[j].Data)
	})

	fmt.Fprintf(&sb, "\n--------------------- Transacoes do mes %s ---------------------\n\n", MeuCaderno.Mes)

	w := tabwriter.NewWriter(&sb, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tDESCRICAO\tTIPO\tVALOR\tPORCENTAGEM\tDATA")
	fmt.Fprintln(w, "--\t---------\t----\t-----\t-----------\t----")

	for _, t := range transacoes {

		if t.Transacao == spt.Ganho {
			formatGasto = "+"
			formatPorcentagem = "---"
		} else {
			formatGasto = "-"
			formatPorcentagem = fmt.Sprintf("%.2f%%", t.Porcentagem)
		}

		fmt.Fprintf(w, "%02d\t%s\t%s\tR$ %.2f\t%s\t%s\n",
			t.ID,
			t.Descricao,
			formatGasto,
			t.Valor,
			formatPorcentagem,
			t.Data.Format("02/01"),
		)
	}

	w.Flush()
	fmt.Fprintf(&sb, "\n---------------------------------\n")

	return sb.String()
}

//Funcao que salva o nosso relatorio em TXT na pasta de downloads;
func SalvarTXT_emDowloads(Texto string, Nome string) (string, error) {
	CAMINHO_DIR, err := spt.CaminhoArquivoDownloads()

	if err != nil {
		return "", err
	}

	CAMINHO_aRQ := filepath.Join(CAMINHO_DIR, Nome)

	//Write cria o arquivo se nao existir, e sobreescreve se ja existir;
	err = os.WriteFile(CAMINHO_aRQ, []byte(Texto), 0644)

	if err != nil {
		return "", err
	}

	return CAMINHO_aRQ, err
}