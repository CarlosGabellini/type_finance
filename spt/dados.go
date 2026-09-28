package spt

import "time"

//spt significa suporte!

type Categoria string

const (
	Superfluo Categoria = "Superfluo"
	ContaCasa Categoria = "ContaCasa"
	Saude Categoria = "Saude"
	Moradia Categoria = "Moradia"
	Lazer Categoria = "Lazer"
)


type MeuCaderno struct {
	Data time.Time `json:"id"`
	Valor float32 `json:"valor"`
	Descricao float32 `json:"descricao"`
	Categoria Categoria `json:"categoria"`
}