package utils

import (
	"os"
)

func DirectoryExists(path string) bool {
	// Stat retorna informações sobre o arquivo/diretório
	// Se o diretório não existir, uma mensagem de erro será retornada
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false // O diretório não existe
	}
	return true // O diretório existe
}
