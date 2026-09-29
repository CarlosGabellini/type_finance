package spt

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

//spt significa suporte!

type Categoria string
type TipoTransacao string

const (
	Superfluo Categoria = "superfluo"
	ContaCasa Categoria = "conta_casa"
	Saude Categoria = "saude"
	Moradia Categoria = "moradia"
	Lazer Categoria = "lazer"
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
	SaldoInicial float64 `json:"saldo_inicial"`
	SaldoFinal float64 `json:"saldo_final"`
	MeusGastos []MinhasTransacoes `json:"meus_gastos"`
}


type MinhasTransacoes struct {
	Data time.Time `json:"data"`
	Valor float64 `json:"valor"`
	Descricao string `json:"descricao"`
	Categoria Categoria `json:"categoria"`
	Parcelado bool `json:"parcelado"`
	NumeroParcelas int `json:"numero_parcelas"`
	Porcentagem float64	`json:"porcentagem"` 	//Em relacao a algum salario fixo!
	Transacao TipoTransacao `json:"transacao"`
	ID int `json:"ID"`
}

//Comecando a fazer as funcoes agora!

//Aqui vamos abrir o arquivo JSON com o Ano correspondente;
func CarregarCaderno(ano int) (Arquivos, error) {

	var contas string = "ano-" + strconv.Itoa(ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)
	var MinhasContas Arquivos

	if err1 != nil {
		return Arquivos{}, err1
	}

	Arq1, err := os.OpenFile(CAMINHO_ARQUIVO, os.O_RDONLY | os.O_CREATE, 0644)

	if err != nil {
		return Arquivos{}, err
	}

	defer Arq1.Close()

	MeuDecoder := json.NewDecoder(Arq1)
	err = MeuDecoder.Decode(&MinhasContas)

	if err != nil {
		//Arquivo recem criado (vazio), devolve io.EOF

		if err == io.EOF {
			//Caderno novo, inicia com os valores padrao!
			MinhasContas = Arquivos{Ano: ano}
			return MinhasContas, nil
		}

		return Arquivos{}, err
	}

	return MinhasContas, nil
}


func ColocarAno(setAno int, MinhasFinancas *Arquivos) int {

	if setAno < 1800 || setAno > 2800 {		//Duvido que esse programa sobreviva ate 2800;
		return 0
	} 

	MinhasFinancas.Ano = setAno
	
	return 1
}

func (MinhasFinancas *MeuCaderno) ColocarData(setInicio, setFim time.Time) error {

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
func (a *Arquivos) NovoCaderno(inicio, fim time.Time, saldoInicial float64) (int, error) {

	NOVO_CADERNO := MeuCaderno{}

	if err := NOVO_CADERNO.ColocarData(inicio, fim); err != nil {
		return -1, err
	}

	NOVO_CADERNO.DefinirSaldoInicial(saldoInicial)

	a.MeuCaderno = append(a.MeuCaderno, NOVO_CADERNO)

	return len(a.MeuCaderno) - 1, nil
}

func (MinhasFin *MeuCaderno) DefinirSaldoInicial(setSaldo float64) {
	NovoSaldo := (math.Round(setSaldo * 100)) / 100
	MinhasFin.SaldoInicial = NovoSaldo
}

func (MinhasFin *MeuCaderno) DefinirSaldoFinal(setSaldo float64) {
	_NovoSaldoFinal := (math.Round(setSaldo * 100)) / 100
	MinhasFin.SaldoFinal = _NovoSaldoFinal
}