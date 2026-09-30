package spt

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
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
	ValorRef float64 `json:"valorRef"`			//Valor de referencia usado para as porcentagem;	
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
func (MeuArquivo1 *Arquivos) CarregarArquivo(ano int) error {

	var contas string = "ano-" + strconv.Itoa(ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)

	if err1 != nil {
		return err1
	}

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
func (a *Arquivos) NovoCaderno(inicio, fim time.Time) (int, error) {

	NOVO_CADERNO := MeuCaderno{}

	if err := NOVO_CADERNO.ColocarData(inicio, fim); err != nil {
		return -1, err
	}
	
	a.MeuCaderno = append(a.MeuCaderno, NOVO_CADERNO)

	return len(a.MeuCaderno) - 1, nil
}

//Define o saldo inicial do nosso arquivo em que estamos trabalhando.
func (MinhasFin *Arquivos) DefinirSaldoInicial(setSaldo float64, indice int) {
	MinhasFin.MeuCaderno[indice].SaldoInicial = Arredondamento2Digitos(setSaldo)
}

func (MinhasFin *Arquivos) DefinirSaldoFinal(setSaldo float64, indice int) {
	MinhasFin.MeuCaderno[indice].SaldoFinal = Arredondamento2Digitos(setSaldo)
}