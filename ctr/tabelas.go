package ctr

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"type_finance/spt"
)

//Aqui eh para mostrar a tabela formatada das despesas, eh melhor usar uma copia do que a struct original;
func TabelaReceitasDespesas(MeuCaderno spt.MeuCaderno) {

	Formatacao()

	Mes := MeuCaderno.Mes
	var formatGasto string
	var formatPorcentagem string

	fmt.Printf("\nBem vindo a tabela de gastos!\n")
	fmt.Printf("Saldo inicial - %.2f\n", MeuCaderno.SaldoInicial)
	fmt.Printf("Saldo final (Ate o momento) - %.2f\n", MeuCaderno.CalcularSaldoFinal())
	
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
				t.Data.Format("02/01/2006"),
		)
	}

	w.Flush()
	fmt.Println()

	fmt.Println("---------------------------------")
	fmt.Println("Digite enter para sair.....")
	LerLinha()
}