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
				time.Sleep(2 * time.Second)

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

	for {
		Formatacao()
		
		fmt.Printf("\n")
		fmt.Printf("Bem vindo! O que deseja fazer?\n")

		fmt.Printf("0 - Abortar\n")
		fmt.Printf("01 - Cadastrar novo mes\n")
		fmt.Printf("02 - Abrir novo mes\n")
		fmt.Printf("03 - Ver meses cadastrados\n")
		fmt.Printf("04 - Ver a tabela de um mes fechado.\n")
		fmt.Printf("05 - Exportar um mes para arquivo TXT.\n")

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

			case 2:
				fmt.Printf("Digite o numero do mes que deseja abrir: ")
				Mes_selecionado, err := strconv.Atoi(LerLinha())

				if err != nil {
					fmt.Println(err)
					continue
				}
				
				AbrirNovoMes(MinhasFin, Mes_selecionado)
				time.Sleep(2 * time.Second)

			case 3:
				VerMesesCadastrados(MinhasFin)

			case 4:
				fmt.Printf("Digite o numero do mes fechado que deseja ver: ")
				Mes, err := strconv.Atoi(LerLinha())

				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				VerMesFechado(MinhasFin, Mes)

				time.Sleep(1 * time.Second)
				continue

			case 5:
				fmt.Printf("Digite o numero do mes que deseja exportar: ")
				Mes, err := strconv.Atoi(LerLinha())

				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				if Mes <= 0 || Mes > 12 {
					fmt.Printf("Mes deve estar entre 1 a 12!\n")
					time.Sleep(1 * time.Second)
					continue
				}

				MesCOPIADO := MinhasFin.MeuCaderno[Mes - 1]
				TEXTO_FORMATADO := FormatarTabela(MesCOPIADO, MinhasFin.Ano)

				NOME_TABELA_METADE := strings.ToLower(MesCOPIADO.Mes) + "-"
				NOME_TABELA_COMPLETO := NOME_TABELA_METADE + strconv.Itoa(MinhasFin.Ano) + ".txt"

				Caminho, err := SalvarTXT_emDowloads(TEXTO_FORMATADO, NOME_TABELA_COMPLETO)

				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				fmt.Printf("Tabela exportada!\n");
				fmt.Printf("Caminho da tabela: %s\n", Caminho)
				time.Sleep(1 * time.Second)
				continue

			default:
				fmt.Printf("Entrada errada! tente novamente!")
				continue
		}
	}
}

func VerMesFechado(MinhasFin *spt.Arquivos, MesInput int) {
	Mes := MesInput - 1

	//Validacao de erros abaixo a partir dessa linha;
	if Mes > 13 || Mes < 0 || Mes >= len(MinhasFin.MeuCaderno) {
		fmt.Printf("Este mes ainda nao foi criado || mes invalido!")
		time.Sleep(1 * time.Second)
		return
	}

	if MinhasFin.MeuCaderno[Mes].Aberto {
		fmt.Printf("Mes aberto! impossivel visualizar!")
		time.Sleep(1 * time.Second)
		return
	}

	if MinhasFin.MeuCaderno[Mes].Mes == "" {
		fmt.Printf("Este mes ainda nao foi criado || mes invalido!!")
		time.Sleep(1 * time.Second)
		return
	}

	PlanilhaMes := &MinhasFin.MeuCaderno[Mes]

	fmt.Printf("Entrando no menu de receitas....")
	time.Sleep(1 * time.Second)

	TabelaReceitasDespesas(*PlanilhaMes)

	fmt.Printf("Tabela foi mostrada!")
	time.Sleep(1 * time.Second)
}