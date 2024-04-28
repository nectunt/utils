package utils

import (
	"os"
)

func GetWorkDir() (string, error) {
	// Obter o diretório de trabalho atual
	dir, err := os.Getwd()
	if err != nil {
		//fmt.Println("Erro ao obter o diretório atual:", err)
		return "", err
	}
	return dir, nil
}
