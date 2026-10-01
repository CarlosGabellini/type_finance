package ctr

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func AbrirNovoMes(MinhasFin *spt.Arquivos, MesInput int) {

	//Isso daqui serve para caso digitar janeiro, ao inves de digitar 0 posso digitar 1;
	Mes := MesInput - 1

	//A variavel de indice se encontra abaixo para manipular os slices do caderno devidamente!
	//var indice int
	
	//Essa condicao eh necessaria para que nao de panic ou slice fora do range;
	if Mes > 13 || Mes < 0 || Mes >= len(MinhasFin.MeuCaderno) {
		fmt.Printf("Este mes ainda nao foi criado || mes invalido!")
		return
	}

	if !MinhasFin.MeuCaderno[Mes].Aberto {
		fmt.Printf("Mes fechado! Impossivel fazer alteracoes nele!")
		return
	}

	if MinhasFin.MeuCaderno[Mes].Mes == "" {
		fmt.Printf("Este mes ainda nao foi criado || mes invalido!!")
		return
	}
	
	Formatacao()

	//Essa funcao tem como objetivo abrir o mes e ter algumas funcionalidades, desde mostrar a tabela,
	// cadastrar receitas e despesas, e muito mais!

	for {
		fmt.Printf("\n")
		fmt.Printf("--------------------- Funcionalidades do Mes e cadastro -------------------------\n")

		fmt.Printf("\t0 - Abortar operacao.\n")
		fmt.Printf("\t01 - Definir saldo inicial.\n")
		fmt.Printf("\t02 - Definir valor referencia\n")
		fmt.Printf("\t03 - Cadastrar Receita/Despesa\n")
		fmt.Printf("\t04 - Excluir Receita/Despesa\n")
		fmt.Printf("\t05 - Imprimir tabela de receita e despesa do mes\n")
		fmt.Printf("\t06 - Calcular o pormenor\n")		//Descricao de itens que nao queremos registrar.
		fmt.Printf("\t07 - Definir o saldo final\n")
		fmt.Printf("\t08 - Fazer o fechamento do mes\n")
		fmt.Printf("\t09 - Fazer o salvamento || checkpoint\n")

		PersonalizarTerminal()
		fmt.Printf("\n")

		fmt.Printf("Digite sua opcao - ")
		controle, err := strconv.Atoi(LerLinha())

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch controle {

			case 0:
				fmt.Printf("Abortando....")
				time.Sleep(2 * time.Second)
				return

			case 1:
				fmt.Printf("Digite o valor do saldo inicial do mes - ")
				saldoInicial, err1 := strconv.ParseFloat(LerLinha(), 64)

				if err1 != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}
				
				MinhasFin.DefinirSaldoInicial(saldoInicial, Mes)

			case 2:
				fmt.Printf("Digite o valor do saldo de referencia - ")
				ref, err2 := strconv.ParseFloat(LerLinha(), 64)

				if err2 != nil {
					fmt.Println(err2)
					time.Sleep(2 * time.Second)
					continue
				}

				MinhasFin.DefinirValorRef(ref, Mes)
				continue

			case 3:
				CadastrarReceitaDespesa(MinhasFin, Mes)
				continue

			case 9:
				Caminho, err := MinhasFin.SalvarAno()

				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Printf("Arquivo salvo!\n")
				fmt.Printf("Caminho: %s\n", Caminho)
				time.Sleep(2 * time.Second)

			default:
				fmt.Printf("Entrada invalida!")
				continue
		}

		Formatacao()
	}
}

func CadastrarReceitaDespesa(MinhasFin *spt.Arquivos, mes int) {

	Formatacao()

	//Criando um nova struct que aponta diretamente para a original, deixa o codigo menos verboso;
	MeuCaderno := &MinhasFin.MeuCaderno[mes]

	for {
		fmt.Printf("\n")
		
		fmt.Printf("------------ Bem vindo ao cadastro de receitas/despesas -----------------\n\n")
		fmt.Printf("As opcoes encontram-se abaixo:\n")
		fmt.Printf("\t 0 - Abortar\n")
		fmt.Printf("\t 01 - Criar uma nova receita/despesa\n")
		fmt.Printf("\t 02 - alterar uma receita/despesa existente\n")
		fmt.Printf("\t 03 - Colocar gasto ou despesa\n")
		fmt.Printf("\t 04 - Colocar data\n")
		fmt.Printf("\t 05 - Colocar o valor\n")
		fmt.Printf("\t 06 - Colocar descricao\n")
		fmt.Printf("\t 07 - Ver as receitas que tenho\n")
		fmt.Printf("\t 08 - Salvar\n")
		fmt.Printf("\nSinta-se livre para alterar quando quiser o valor caso esteja errado.\n")

		PersonalizarTerminal()

		fmt.Printf("\nDigite sua opcao - ")
		controle, err := strconv.Atoi(LerLinha())

		var indice int = -1

		if err != nil {
			fmt.Println(err)
			time.Sleep(2 * time.Second)
			continue
		}

		switch controle {

			case 0:
				fmt.Printf("Abortando operation.....")
				time.Sleep(2 * time.Second)
				return

			case 1:
				fmt.Printf("Fazendo o seu cadastro...")
				indice = MeuCaderno.DefinirUmaNovaTransacao()
				fmt.Printf("Criado o cadastro! - indice %d\n", indice)
				
				time.Sleep(2 * time.Second)
				continue
				
			default:
				fmt.Printf("Entrada invalida! digite novamente!")
				time.Sleep(2 * time.Second)
				continue				
		}

		Formatacao()
	}
}

func VerMesesCadastrados(MinhasFin *spt.Arquivos) {

	Formatacao()
	fmt.Printf("\nAqui estao os meses cadastrados: \n")

	for _, meses := range MinhasFin.MeuCaderno {
		fmt.Printf("Mes - %s\t\t|Aberto: %t\n", meses.Mes, meses.Aberto)
	}

	fmt.Printf("\n\n")
	fmt.Printf("Este sao os meses cadastrados!\n")
	fmt.Printf("Aperte enter para voltar....")

	LerLinha()
}