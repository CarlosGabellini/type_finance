package spt

import "time"

//spt significa suporte!

type Categoria string
type Transacao string

const (
	Superfluo Categoria = "superfluo"
	ContaCasa Categoria = "conta_casa"
	Saude Categoria = "saude"
	Moradia Categoria = "moradia"
	Lazer Categoria = "lazer"
	Ganho Transacao = "ganho"
	Gasto Transacao = "gasto"
)

type MeuCaderno struct {
	Ano int `json:"ano"`
	DataInicio time.Time `json:"data_inicio"`
	DataFim time.Time `json:"data_fim"`
	SaldoInicial float64 `json:"saldo_inicial"`
	SaldoFinal float64 `json:"saldo_final"`
	MeusGastos []Gastos
}


type Gastos struct {
	Data time.Time `json:"data"`
	Valor float64 `json:"valor"`
	Descricao string `json:"descricao"`
	Categoria Categoria `json:"categoria"`
	Parcelado bool `json:"parcelado"`
	NumeroParcelas int `json:"numero_parcelas"`
	Porcentagem float64	`json:"porcentagem"` 	//Em relacao a algum salario fixo!
	Transacao Transacao `json:"transacao"`
	ID int `json:"ID"`
}

//Comecando a fazer as funcoes agora!

//Aqui vamos abrir o arquivo JSON com o Ano correspondente;
func CarregarAno(ano int) (MeuCaderno, error) {}