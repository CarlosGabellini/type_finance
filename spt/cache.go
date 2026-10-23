package spt

import (
	"os"
	"path/filepath"
)

/*----------------------------------------------- Sobre este Arquivo -------------------------------------/
	Este arquivo tem como objetivo de responsabilidade guardar e mostrar os caminhos para a persistencia
de dados do nosso aplicativo de terminal; ele basicamente faz a manipulacao de caminhos e os respectivos
arquivos.
	A funcao principal eh CaminhoCache(), que basicamente retorna a pasta cache padrao do usuario em um
dispositivo Windows, Linux ou MacOS, depois dessa funcao, conseguimos criar a pasta principal do nosso
aplicativo, onde depois fazemos a manipulacao de outros arquivos vigentes.
----------------------------------------------------------------------------------------------------------/
 */

//Cria o caminho do cache e a pasta principal para persistir nossos dados;
func CaminhoCache() (string, error) {

	CaminhoDoCache, err := os.UserCacheDir()

	if err != nil {
		return "", err
	}

	CaminhoDaPasta := filepath.Join(CaminhoDoCache, "type_finance")
	err = os.Mkdir(CaminhoDaPasta, 0755)

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	//Se chegou ate aqui ou a pasta ja foi criada, da no mesmo!
	return CaminhoDaPasta, nil
}

//Cria o diretorio do JSON, onde os arquivos ficam numa pasta mais organizada;
func CreateDirJSON() (string, error) {

	MeuCaminhoCache, err := CaminhoCache()
	_subpasta := "type_json"
	
	if err != nil {
		return "", err
	}

	CaminhoSUBPASTa := filepath.Join(MeuCaminhoCache, _subpasta)
	err = os.Mkdir(CaminhoSUBPASTa, 0755)

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	return CaminhoSUBPASTa, nil
}

//For future.
func CreateDirBin() (string, error) {
	MeuCaminhoCache, err := CaminhoCache()
	const pastaBin string = "type_bin"
	
	if err != nil {
		return "", err
	}

	CaminhoSUBPASTa := filepath.Join(MeuCaminhoCache, pastaBin)
	err = os.Mkdir(CaminhoSUBPASTa, 0755)

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	return CaminhoSUBPASTa, nil
}

//Cria um caminho de um arquivo JSON e manda em string onde ele fica;
func CaminhoArquivo(setNome string) (string, error) {

	CaminhoDaPasta, err := CreateDirJSON()

	if err != nil || CaminhoDaPasta == "" {
		return "", err
	}

	CAMINHO_ARQ := filepath.Join(CaminhoDaPasta, setNome)

	return CAMINHO_ARQ, nil
}

//Cria um arquivo binario a partir de um diretorio ja feito;
func CaminhoArquivoBIN(setNome string) (string, error) {
	CaminhoDaPasta, err := CreateDirBin()

	if err != nil || CaminhoDaPasta == "" {
		return "", err
	}

	CAMINHO_ARQ := filepath.Join(CaminhoDaPasta, setNome)

	return CAMINHO_ARQ, nil
}