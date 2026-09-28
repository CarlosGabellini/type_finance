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

func CaminhoArquivo(setNome string) (string, error) {

	CaminhoDaPasta, err := CaminhoCache()

	if err != nil || CaminhoDaPasta == "" {
		return "", err
	}

	CAMINHO_ARQ := filepath.Join(CaminhoDaPasta, setNome)

	return CAMINHO_ARQ, nil
}