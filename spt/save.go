package spt

import (
	"encoding/json"
	"math"
	"os"
	"strconv"
)

//Arquivo save.go responsavel por salvar o arquivo. (funcao essencial)
// E mais algumas outras funcoes espalhadas.

func (MeuArquivo Arquivos) SalvarAno() (string, error) {

	var contas string = "ano-" + strconv.Itoa(MeuArquivo.Ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)

	if err1 != nil {
		return "", err1
	}

	Arq2, err := os.OpenFile(CAMINHO_ARQUIVO, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)

	if err != nil {
		return "", err
	}

	defer Arq2.Close()

	//DOC = Documento;
	DOC := json.NewEncoder(Arq2)
	DOC.SetIndent("", "  ")
	err = DOC.Encode(&MeuArquivo)

	if err != nil {
		return "", err
	}

	return CAMINHO_ARQUIVO, err
}

//Retorna a porcentagem de um valor de referencia;
func CalcularPorcentagem(valor float64, total float64) float64 {
	NovoValor := Arredondamento2Digitos(valor)
	NovoTotal := Arredondamento2Digitos(total)

	return (NovoValor * 100) / NovoTotal
}

//Retorna um arrendodamento de 2 digitos para float64, exemplo: 3.1517 -> 3.15;
func Arredondamento2Digitos(valor float64) float64 {
	return (math.Round(valor * 100)) / 100
}