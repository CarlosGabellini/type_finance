package spt

import (
	"os"
	"path/filepath"
)

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

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	return CaminhoSUBPASTa, nil
}

func CaminhoArquivo(setNome string) (string, error) {

	CaminhoDaPasta, err := CreateDirJSON()

	if err != nil || CaminhoDaPasta == "" {
		return "", err
	}

	CAMINHO_ARQ := filepath.Join(CaminhoDaPasta, setNome)

	return CAMINHO_ARQ, nil
}