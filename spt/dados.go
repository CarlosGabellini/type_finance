package spt

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//spt significa suporte!

type TipoTransacao string


//Nao altere este slice!
var NOMES_DOS_MESES = [...]string{
	"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
	"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
}

const (
	Ganho TipoTransacao = "ganho"
	Gasto TipoTransacao = "gasto"
)

type Arquivos struct {
	Ano int `json:"ano"`
	MeuCaderno []MeuCaderno `json:"cadernos"` 
}

type MeuCaderno struct {
	DataInicio time.Time `json:"data_inicio"`
	DataFim time.Time `json:"data_fim"`
	Mes string `json:"mes"`
	SaldoInicial float64 `json:"saldo_inicial"`
	SaldoFinal float64 `json:"saldo_final"`
	ValorRef float64 `json:"valorRef"`			//Valor de referencia usado para as porcentagem;
	Aberto bool `json:"aberto"`					//Aberto ou fechado significa que o mes nao pode mais
	MeusGastos []MinhasTransacoes `json:"meus_gastos"`		//ser alterado;
}


type MinhasTransacoes struct {
	Data time.Time `json:"data"`
	Valor float64 `json:"valor"`	
	Descricao string `json:"descricao"`			//Somente descricao eh importante, sem prescisar de categoria;
	Porcentagem float64	`json:"porcentagem"` 	//Em relacao a algum salario fixo!
	Transacao TipoTransacao `json:"transacao"`
	ID int `json:"ID"`
}

//Comecando a fazer as funcoes agora!

//Aqui vamos abrir o arquivo JSON com o Ano correspondente;
func (MeuArquivo1 *Arquivos) CarregarArquivo(ano int) error {
	
	//Aqui prescisa mecher com um ponteiro para alterar a struct original;
	var contas string = "ano-" + strconv.Itoa(ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)

	if err1 != nil {
		return err1
	}

	//ReadOnly por que nao prescisamos alterar o arquivo, somente ler ele;
	Arq1, err := os.OpenFile(CAMINHO_ARQUIVO, os.O_RDONLY, 0644)

	if err != nil {
		return err
	}

	defer Arq1.Close()

	var CARREGADO Arquivos
	
	MeuDecoder := json.NewDecoder(Arq1)
	err = MeuDecoder.Decode(&CARREGADO)

	if err != nil {
		//Arquivo recem criado (vazio), devolve io.EOF

		if err == io.EOF {
			//Caderno novo, inicia com os valores padrao!
			*MeuArquivo1 = Arquivos{Ano: ano}
			return nil
		}

		return err
	}

	//Aqui atribuimos a nossa struct principal para alterar a variavel global!
	*MeuArquivo1 = CARREGADO

	return nil
}

func ListarArquivos() ([]string, error) {

	LISTA_DIRETORIOS := make([]string, 0, 15)
	Caminho_Do_Cache, err := CaminhoCache()

	if err != nil {
		return nil, err
	}
	
	err = filepath.WalkDir(Caminho_Do_Cache, func(root string, d os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !d.IsDir() {
			LISTA_DIRETORIOS = append(LISTA_DIRETORIOS, root)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return LISTA_DIRETORIOS, nil
}

func ColocarAno(setAno int, MinhasFinancas *Arquivos) int {
	if setAno < 1800 || setAno > 2800 {		//Duvido que esse programa sobreviva ate 2800;
		return 0
	}

	var contas string = "ano-" + strconv.Itoa(setAno) + ".json"
	CAMINHO_ARQ, err := CaminhoArquivo(contas)
	ARQUIVOS_DIRETORIO, err1 := ListarArquivos()

	if err != nil {
		return 0
	}

	if err1 != nil {
		return 0
	}

	//Validando para ver se ja tem um arquivo ja existente;
	for _, ArqDir := range ARQUIVOS_DIRETORIO {
		if filepath.Base(ArqDir) == filepath.Base(CAMINHO_ARQ) {
			return 2
		}
	}

	MinhasFinancas.Ano = setAno

	return 1
}

func (MinhasFinancas *MeuCaderno) ColocarData1(setInicio, setFim time.Time) error {

	if setInicio.IsZero() || setFim.IsZero() {
		PersonError := fmt.Errorf("Data invalida!")
		return PersonError
	}

	if setFim.Before(setInicio) {
		PersonError := fmt.Errorf("Fim antes de inicio! invalido!")
		return PersonError
	}

	if setInicio.Year() < 1800 || setFim.Year() > 2800 {
		PersonError := fmt.Errorf("Data invalida!")
		return PersonError
	}

	if setFim.Year() < 1800 || setInicio.Year() > 2800 {
		PersonError := fmt.Errorf("Data invalida!")
		return  PersonError
	}

	MinhasFinancas.DataInicio = setInicio
	MinhasFinancas.DataFim = setFim

	return nil
}

//Retorna um novo caderno para a struct devovendo o ultimo indice dele;
func (a *Arquivos) NovoCaderno(inicio, fim time.Time) (int, error) {

	NOVO_CADERNO := MeuCaderno{}

	if err := NOVO_CADERNO.ColocarData1(inicio, fim); err != nil {
		return -1, err
	}

	NOVO_CADERNO.Aberto = true
	NOVO_CADERNO.Mes = NOMES_DOS_MESES[inicio.Month() - 1]
	//Time.Month vai de 1 a 12, entao prescisa ser -1 para que o array funcione corretamente;
	
	a.MeuCaderno = append(a.MeuCaderno, NOVO_CADERNO)

	return len(a.MeuCaderno) - 1, nil
}

//Define o saldo inicial do nosso arquivo em que estamos trabalhando.
func (MinhasFin *Arquivos) DefinirSaldoInicial(setSaldo float64, MesIndice int) {

	if MinhasFin.MeuCaderno[MesIndice].SaldoInicial != 0 {
		fmt.Printf("Saldo ja cadastrado! deseja realmente alterar?")
		time.Sleep(1 * time.Second)
		return
	}
	
	MinhasFin.MeuCaderno[MesIndice].SaldoInicial = setSaldo
	NovoSaldo := MinhasFin.MeuCaderno[MesIndice].SaldoInicial
	
	fmt.Printf("Saldo guardado com sucesso! Valor: %.2f\n", NovoSaldo)

	time.Sleep(1 * time.Second)
}

func (MinhasFin *Arquivos) DefinirSaldoFinal(setSaldo float64, indice int) {
	MinhasFin.MeuCaderno[indice].SaldoFinal = Arredondamento2Digitos(setSaldo)
}

func (MinhasFin *Arquivos) DefinirValorRef(setValor float64, indice int) {

	if MinhasFin.MeuCaderno[indice].ValorRef != 0 {
		fmt.Printf("Deseja alterar o valor de referencia? ja tem um cadastrado.")
		time.Sleep(2 * time.Second)
		return
	}

	MinhasFin.MeuCaderno[indice].ValorRef = setValor
	MeuValor := MinhasFin.MeuCaderno[indice].ValorRef

	fmt.Printf("Valor de referencia salvo - %.2f\n", MeuValor)
	time.Sleep(2 * time.Second)
}

func (MeuCaderno *MeuCaderno) DefinirUmaNovaTransacao(arq1 *Arquivos) (int) {
	var MeuID int
	MeuID = CalcularIDMaior(arq1)
	Nova := MinhasTransacoes{ID: MeuID + 1}
	
	MeuCaderno.MeusGastos = append(MeuCaderno.MeusGastos, Nova)

	return len(MeuCaderno.MeusGastos) - 1
}

func (MeuCaderno *MeuCaderno) DefinirReceitaGasto(indice int, gastoOuReceita int) int{
	if gastoOuReceita == 0 {
		MeuCaderno.MeusGastos[indice].Transacao = TipoTransacao(Ganho)
		return 0
	}

	if gastoOuReceita == 1 {
		MeuCaderno.MeusGastos[indice].Transacao = TipoTransacao(Gasto)
		return 1
	}

	return -1
}

func (MeuCaderno *MeuCaderno) ColocarDataGastos(indice int, ano int) error {
	
	if indice < 0 || indice >= len(MeuCaderno.MeusGastos) {
		return fmt.Errorf("indice invalido: %d", indice)
	}

	for {
		fmt.Printf("Digite a data (dia/mes, ex: 05/03) - ")
		entrada := strings.TrimSpace(LerNovaLinha())

		data, err := time.Parse("02/01/2006", fmt.Sprintf("%s/%d", entrada, ano))
		
		if err != nil {
			fmt.Println("Data invalida! Use o formato dia/mes (ex: 05/03).")
			continue
		}

		if data.Before(MeuCaderno.DataInicio) || data.After(MeuCaderno.DataFim) {
			fmt.Printf("Data fora do periodo do caderno(%s a %s)!\n", 
				MeuCaderno.DataInicio.Format("02/01"),
				MeuCaderno.DataFim.Format("02/01"))

			continue
		}

		MeuCaderno.MeusGastos[indice].Data = data
		return nil
	}
}

func (MeuCaderno *MeuCaderno) ColocarValorGastos(indice int, setValor float64) error {
	
	if indice < 0 || indice >= len(MeuCaderno.MeusGastos) {
		return fmt.Errorf("indice invalido: %d\n", indice)
	}

	if setValor < 0 {
		return fmt.Errorf("Valor invalido, deve ser positivo.\n")
	}

	MeuCaderno.MeusGastos[indice].Valor = setValor

	return nil
}

func (MeuCaderno *MeuCaderno) ColocarDescricao(indice int, descricao string) error {

	var maximaDescricao int = 20
	
	if indice < 0 || indice >= len(MeuCaderno.MeusGastos) {
		return fmt.Errorf("indice invalido!\n")
	}

	if descricao == "" {
		return fmt.Errorf("A descricao nao pode ser vazia!")
	}

	if n := utf8.RuneCountInString(descricao); n > maximaDescricao {
		return fmt.Errorf("Descricao muito longa, tem %d caracteres, maximo: %d\n", n, maximaDescricao)
	}

	MeuCaderno.MeusGastos[indice].Descricao = descricao
	return nil
}

func (MeuCard *MeuCaderno) AtualizarPorcentagemRef(indice int, setvalor float64) error {

	if indice < 0 || indice >= len(MeuCard.MeusGastos) {
		return fmt.Errorf("Indice invalido!")
	}

	if MeuCard.MeusGastos[indice].Transacao != Gasto {
		return fmt.Errorf("Isso nao eh um gasto para calcular!")
	}

	Total := MeuCard.ValorRef
	MeuCard.MeusGastos[indice].Porcentagem = CalcularPorcentagem(setvalor, Total)

	return nil
}

func (MeuCard *MeuCaderno) CalcularSaldoFinal() float64 {

	MeuSaldoAtual := Arredondamento2Digitos(MeuCard.SaldoInicial)

	for _, calc := range MeuCard.MeusGastos {

		if calc.Transacao == Ganho {
			MeuSaldoAtual += Arredondamento2Digitos(calc.Valor)
		}

		if calc.Transacao == Gasto {
			MeuSaldoAtual -= Arredondamento2Digitos(calc.Valor)
		}

		if calc.Transacao == "" {
			continue
		}
	}

	return Arredondamento2Digitos(MeuSaldoAtual)
}