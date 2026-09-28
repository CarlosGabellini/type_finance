package spt

import (
	"encoding/json"
	"io"
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

type MeuCaderno struct {
	Ano int `json:"ano"`
	DataInicio time.Time `json:"data_inicio"`
	DataFim time.Time `json:"data_fim"`
	SaldoInicial float64 `json:"saldo_inicial"`
	SaldoFinal float64 `json:"saldo_final"`
	MeusGastos []Gastos `json:"meus_gastos"`
}


type Gastos struct {
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
func CarregarCaderno(ano int) (MeuCaderno, error) {

	var contas string = "ano-" + strconv.Itoa(ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)
	var MinhasContas MeuCaderno

	if err1 != nil {
		return MeuCaderno{}, err1
	}

	Arq1, err := os.OpenFile(CAMINHO_ARQUIVO, os.O_RDONLY | os.O_CREATE, 0644)

	if err != nil {
		return MeuCaderno{}, err
	}

	defer Arq1.Close()

	MeuDecoder := json.NewDecoder(Arq1)
	err = MeuDecoder.Decode(&MinhasContas)

	if err != nil {
		//Arquivo recem criado (vazio), devolve io.EOF

		if err == io.EOF {
			//Caderno novo, inicia com os valores padrao!
			MinhasContas = MeuCaderno{Ano: ano}
			return MinhasContas, nil
		}

		return MeuCaderno{}, err
	}

	return MinhasContas, nil
}


func ColocarAno(setAno int, MinhasFinancas *MeuCaderno) int {

	if setAno < 1800 || setAno > 2800 {		//Duvido que esse programa sobreviva ate 2800;
		return 0
	} 

	MinhasFinancas.Ano = setAno
	
	return 1
}