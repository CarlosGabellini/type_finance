package spt

import "time"

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