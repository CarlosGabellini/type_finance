package spt

import (
	"encoding/json"
	"os"
	"strconv"
)

//Arquivo save.go responsavel por salvar o arquivo. (funcao essencial)
// E mais algumas outras funcoes espalhadas.

func SalvarArquivo(MeuArquivo Arquivos) error {

	var contas string = "ano-" + strconv.Itoa(MeuArquivo.Ano) + ".json"
	CAMINHO_ARQUIVO, err1 := CaminhoArquivo(contas)

	if err1 != nil {
		return err1
	}

	Arq2, err := os.OpenFile(CAMINHO_ARQUIVO, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer Arq2.Close()

	//DOC = Documento;
	DOC := json.NewEncoder(Arq2)

	err = DOC.Encode(&MeuArquivo)

	if err != nil {
		return err
	}

	return nil
}

func CalcularPorcentagem(valor float64, total float64) float64 {
	return (valor * 100) / total
}