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