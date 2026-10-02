package spt

import (
	"bufio"
	"encoding/json"
	"math"
	"os"
	"strconv"
	"strings"
)

//Arquivo save.go responsavel por salvar o arquivo. (funcao essencial)
// E mais algumas outras funcoes espalhadas.

/*---------------------------------------------- Avisos --------------------------------------------------/
		Eh importante distinguir quando eu presciso alterar a struct original com ponteiro e distinguir
	quando NAO presciso usar um ponteiro para o original, nesse caso aqui, na hora de salvar o arquivo
	eu nao presciso usar a struct original, por que isso vira o ponteiro do ponteiro na hora de salvar,
	tudo o que eu prescisava era criar uma copia para salvar ela no meu disco.
//-------------------------------------------------------------------------------------------------------/
*/

var NovaEntrada = bufio.NewReader(os.Stdin)

//Salva o ano retornando o caminho do arquivo e um erro caso nao tenha conseguido fazer;
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

func CalcularIDMaior(arq2 *Arquivos) int {

	var MeuID int = 0
	
	for _, Cadernos := range arq2.MeuCaderno {
		for _, gastos := range Cadernos.MeusGastos {
			if MeuID < gastos.ID {
				MeuID = gastos.ID
			}
		}
	}
	
	return MeuID
}

func LerNovaLinha() string {
	linha1, _ := NovaEntrada.ReadString('\n')
	return strings.TrimSpace(linha1)
}