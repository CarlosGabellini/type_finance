package ctr

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	//Essa funcao tem como objetivo abrir o mes e ter algumas funcionalidades, desde mostrar a tabela,
	// cadastrar receitas e despesas, e muito mais!


	//SOMENTE USE ESTA OPCAO PARA O CASE 5, eh para imprimir somente a tabela;
	PlanilhaMes := &MinhasFin.MeuCaderno[Mes]
	
	for {
		Formatacao()
		
		fmt.Printf("\n")
		fmt.Printf("--------------------- Funcionalidades do Mes e cadastro -------------------------\n")

		fmt.Printf("\t0 - Abortar operacao.\n")
		fmt.Printf("\t01 - Definir saldo inicial.\n")
		fmt.Printf("\t02 - Definir valor referencia\n")
		fmt.Printf("\t03 - Cadastrar Receita/Despesa\n")
		fmt.Printf("\t04 - Excluir Receita/Despesa\n")
		fmt.Printf("\t05 - Imprimir tabela de receita e despesa do mes\n")
		fmt.Printf("\t06 - Definir o saldo final\n")
		fmt.Printf("\t07 - Fazer o fechamento do mes\n")
		fmt.Printf("\t08 - Fazer o salvamento || checkpoint\n")

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
					fmt.Println(err1)
					time.Sleep(2 * time.Second)
					continue
				}
				
				error1 := MinhasFin.DefinirSaldoInicial(saldoInicial, Mes)

				if error1 != nil {
					fmt.Println(error1)
					time.Sleep(1 * time.Second)
					continue
				}

				continue

			case 2:
				fmt.Printf("Digite o valor do saldo de referencia - ")
				ref, err2 := strconv.ParseFloat(LerLinha(), 64)

				if err2 != nil {
					fmt.Println(err2)
					time.Sleep(2 * time.Second)
					continue
				}

				//Ja tem um time sleep e uma mensagem pronta a partir desta funcao;
				error2 := MinhasFin.DefinirValorRef(ref, Mes)

				if error2 != nil {
					fmt.Println(error2)
					time.Sleep(1 * time.Second)
					continue
				}

				continue

			case 3:
				CadastrarReceitaDespesa(MinhasFin, Mes)
				continue

			case 5:
				fmt.Printf("Entrando no menu das receitas....")
				time.Sleep(1 * time.Second)
				TabelaReceitasDespesas(*PlanilhaMes)

				fmt.Printf("tabela mostrada!")
				time.Sleep(1 * time.Second)
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
	}
}

func CadastrarReceitaDespesa(MinhasFin *spt.Arquivos, mes int) {
	
	//Criando um nova struct que aponta diretamente para a original, deixa o codigo menos verboso;
	MeuCaderno := &MinhasFin.MeuCaderno[mes]

	//O indice deve ser colocado aqui fora para ele nao sofrer reatribuicao de valor toda hora!
	var indice int = -1

	if len(MeuCaderno.MeusGastos) > 0 {
		indice = 0
	}

	//Funcao anonima usada para nao ter indice invalido! e para nao fechar o programa inesperadamente;
	//NAO MECHER!
	indiceValido := func() bool {
		return indice >= 0 && indice < len(MeuCaderno.MeusGastos)
	}
	
	for {
		Formatacao()
		fmt.Printf("\n")
		
		fmt.Printf("------------ Bem vindo ao cadastro de receitas/despesas -----------------\n\n")
		fmt.Printf("As opcoes encontram-se abaixo:\n")
		fmt.Printf("\t 0 - Abortar\n")
		fmt.Printf("\t 01 - Criar uma nova receita/despesa\n")
		fmt.Printf("\t 02 - alterar uma receita/despesa existente\n")
		fmt.Printf("\t 03 - Colocar gasto ou despesa\n")
		fmt.Printf("\t 04 - Definir valor && data\n")
		fmt.Printf("\t 05 - Colocar descricao\n")
		fmt.Printf("\t 06 - Ver as receitas que tenho\n")
		fmt.Printf("\t 07 - Ver saldo atual\n")
		fmt.Printf("\t 08 - Calcular o pormenor\n")
		fmt.Printf("\nSinta-se livre para alterar quando quiser o valor caso esteja errado.\n")

		//Na opcao 6 eu poderia excluir ela e somente deixar no switch de cima, mas prefiri deixar em ambos
		//os casos para os usuarios nao ficarem prescisando navegar toda hora, tem que ser algo pratico e 
		// intuitivo;

		PersonalizarTerminal()

		if indiceValido() {
			fmt.Printf("\nAtualmente mechendo no ID - %d\n", MeuCaderno.MeusGastos[indice].ID)

		} else {
			fmt.Printf("\nNenhuma transacao selecionada! use a opcao 1 para criar uma transacao.\n")
		}
		
		fmt.Printf("Digite sua opcao - ")
		controle, err := strconv.Atoi(LerLinha())

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
				indice = MeuCaderno.DefinirUmaNovaTransacao(MinhasFin)
				fmt.Printf("Criado o cadastro! - indice %d\n", indice)
				fmt.Printf("Id do cadastro - %d\n", MeuCaderno.MeusGastos[indice].ID)
				
				time.Sleep(2 * time.Second)
				continue

			case 2:
				fmt.Printf("Digite o ID que gostaria de buscar - ")
				busca, err := strconv.Atoi(LerLinha())

				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				//Reatribuindo em outra variavel para nao ter o crash novamente;
				achado := MeuCaderno.BuscarID(busca)

				if indice == -1 {
					fmt.Printf("ID nao encontrado!\n")
					time.Sleep(1 * time.Second)
					continue
				}

				indice = achado

				fmt.Printf("ID encontrado - Voce pode alterar ele agora.")
				continue

			case 3:

				if !indiceValido() {
					fmt.Println("Nenhuma transacao selecionada! Crie ou busque uma primeiro.")
					time.Sleep(2 * time.Second)
					continue
				}

				if MeuCaderno.MeusGastos[indice].Transacao != "" {
					fmt.Printf("Gasto || Receita ja cadastrados! deseja realmente alterar?\n")
					fmt.Printf("1 - Para sim ------ 0 - Para nao\n")
					fmt.Printf("Digite sua opcao - ")

					descisao, err := strconv.Atoi(LerLinha())

					if err != nil {
						fmt.Println(err)
						time.Sleep(1 * time.Second)
						continue
					}

					if descisao == 1 {
						fmt.Printf("Pode continuar!\n")
					}

					if descisao == 0 {
						fmt.Printf("Nao alterado!\n")
						time.Sleep(1 * time.Second)
						continue
					}

					if descisao != 1 && descisao != 0 {
						fmt.Printf("Voce nao digitou nem 0 e nem 1!")
						time.Sleep(1 * time.Second)
						continue
					}
				}

				Formatacao()
				fmt.Printf("\nColocando o gasto/receita\n")
				fmt.Printf("gasto - 1\tReceita - 0\n")
				controle2, err3 := strconv.Atoi(LerLinha())

				if err3 != nil {
					fmt.Println(err3)
					time.Sleep(2 * time.Second)
					continue
				}
				
				mostrador := MeuCaderno.DefinirReceitaGasto(indice, controle2)

				if mostrador == -1 {
					fmt.Println("Numero errado! coloque outro!")
					time.Sleep(2 * time.Second)
					continue
				}

				if mostrador == 0 {
					fmt.Println("Receita feita!")
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Printf("Gasto registrado!\n")
				time.Sleep(2 * time.Second)
				continue

			case 4:
				if !indiceValido() {
					fmt.Println("Nenhuma transacao selecionada! Crie ou busque uma primeiro.")
					time.Sleep(2 * time.Second)
					continue
				}
			
				if MeuCaderno.MeusGastos[indice].Transacao == "" {
					fmt.Println("Tipo de transacao nao definido! coloque ela primeiro!")
					time.Sleep(2 * time.Second)
					continue
				}

				if MeuCaderno.MeusGastos[indice].Transacao != "" {
					fmt.Printf("Gasto || Receita ja cadastrados! deseja realmente alterar?\n")
					fmt.Printf("1 - Para sim ------ 0 - Para nao\n")
					fmt.Printf("Digite sua opcao - ")

					//Meu programa esta fechando inesperadamente, pode ser o nome da variavel que eh igual;
					descisao2, err := strconv.Atoi(LerLinha())

					if err != nil {
						fmt.Println(err)
						time.Sleep(1 * time.Second)
						continue
					}

					if descisao2 == 1 {
						fmt.Printf("Pode continuar!\n")
					}

					if descisao2 == 0 {
						fmt.Printf("Nao alterado!\n")
						time.Sleep(1 * time.Second)
						continue
					}

					if descisao2 != 1 && descisao2 != 0 {
						fmt.Printf("Voce nao digitou nem 0 e nem 1!")
						time.Sleep(1 * time.Second)
						continue
					}
				}

				Formatacao()

				fmt.Printf("\nDefinindo a data && valor\n")
				err4 := MeuCaderno.ColocarDataGastos(indice, MinhasFin.Ano)

				if err4 != nil {
					fmt.Println(err4)
					time.Sleep(2 * time.Second)
					continue
				}

				if MeuCaderno.MeusGastos[indice].Transacao == spt.TipoTransacao(spt.Ganho) {
					fmt.Printf("Digite o valor para a receita - ")
					ColocarValor, err5 := strconv.ParseFloat(LerLinha(), 64)

					if err5 != nil {
						fmt.Println(err5)
						time.Sleep(2 * time.Second)
						continue
					}

					err5 = MeuCaderno.ColocarValorGastos(indice, ColocarValor)

					if err5 != nil {
						fmt.Println(err5)
						time.Sleep(2 * time.Second)
						continue
					}

					fmt.Printf("Cadastro da receita concluido!\n")
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Printf("Digite o valor para gastos - ")
				ColocarValor, err5 := strconv.ParseFloat(LerLinha(), 64)

				if err5 != nil {
					fmt.Println(err5)
					time.Sleep(2 * time.Second)
					continue
				}

				err5 = MeuCaderno.ColocarValorGastos(indice, ColocarValor)

				if err5 != nil {
					fmt.Println(err5)
					time.Sleep(2 * time.Second)
					continue
				}

				err6 := MeuCaderno.AtualizarPorcentagemRef(indice, ColocarValor)

				if err6 != nil {
					fmt.Println(err6)
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Println("Registro feito com sucesso!")
				time.Sleep(2 * time.Second)
				continue

			case 5:
				//Fazendo protecao de indices para nao acessar algo invalido!
				if !indiceValido() {
					fmt.Println("Nenhuma transacao selecionada! Crie ou busque uma primeiro.")
					time.Sleep(2 * time.Second)
					continue
				}
				
				fmt.Printf("Digite uma descricao - ")

				if indice == -1 {
					fmt.Printf("Indice invalido! Cadastre uma transacao!")
					time.Sleep(2 * time.Second)
					continue
				}
			
				if MeuCaderno.MeusGastos[indice].Transacao == "" {
					fmt.Println("Tipo de transacao nao definido! coloque ela primeiro!")
					time.Sleep(2 * time.Second)
					continue
				}

				colocarDescricao := strings.TrimSpace(LerLinha())

				err7 := MeuCaderno.ColocarDescricao(indice, colocarDescricao)

				if err7 != nil {
					fmt.Println(err7)
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Println("A descricao foi devidamente colocada!")
				time.Sleep(1 * time.Second)
				continue

			case 6:
				fmt.Printf("Entrando no menu das receitas....")
				time.Sleep(1 * time.Second)
				TabelaReceitasDespesas(*MeuCaderno)

				fmt.Printf("tabela mostrada!")
				time.Sleep(1 * time.Second)
				continue

			case 7:
				fmt.Printf("Saldo atual encontra-se abaixo: ")
				fmt.Printf("%.2f\n", MeuCaderno.CalcularSaldoFinal())

				time.Sleep(3 * time.Second)
				continue

			case 8:
				myerror := spt.CalcularPormenor(MinhasFin, MeuCaderno)

				if myerror != nil {
					fmt.Println(myerror)
					time.Sleep(2 * time.Second)
					continue
				}
				
				continue
				
			default:
				fmt.Printf("Entrada invalida! digite novamente!")
				time.Sleep(2 * time.Second)
				continue
		}
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