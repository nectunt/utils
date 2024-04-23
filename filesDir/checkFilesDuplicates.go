package filesDir

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func checkDuplicates(root string) (map[string][]string, error) {
	// Mapa para armazenar os arquivos por hash
	fileMap := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Ignorar erros de permissão
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			// Calcular hash do arquivo
			hash, err := fileHash(path)
			if err != nil {
				return err
			}

			// Adicionar arquivo ao mapa usando o hash como chave
			fileMap[hash] = append(fileMap[hash], path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Criar mapa para armazenar arquivos duplicados
	duplicates := make(map[string][]string)

	// Iterar sobre o mapa de arquivos por hash
	for _, files := range fileMap {
		if len(files) > 1 {
			// Se houver mais de um arquivo com o mesmo hash, são duplicados
			for _, file := range files {
				duplicates[file] = files
			}
		}
	}

	return duplicates, nil
}

func fileHash(path string) (string, error) {
	// Abrir arquivo
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Criar hash SHA256
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Retornar hash como string hex
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func writeDuplicatesToFile(path string, duplicates map[string][]string) error {
	// Criar ou abrir arquivo para escrita
	file, err := os.Create(path + "arquivos_duplicados.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// Escrever caminhos dos arquivos duplicados no arquivo
	for _, files := range duplicates {
		if len(files) > 1 {
			_, err := file.WriteString("Arquivos Duplicados:\n")
			if err != nil {
				return err
			}
			for _, f := range files {
				_, err := file.WriteString(fmt.Sprintf("%s\n", f))
				if err != nil {
					return err
				}
			}
			_, err = file.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// path = Diretório a ser verificado
func CheckFilesDuplicates(path string) {
	// Verificar arquivos duplicados
	duplicates, err := checkDuplicates(path)
	if err != nil {
		fmt.Println("Erro ao verificar arquivos duplicados:", err)
		return
	}

	// Escrever arquivos duplicados em um arquivo de texto
	err = writeDuplicatesToFile(path, duplicates)
	if err != nil {
		fmt.Println("Erro ao escrever arquivo de texto:", err)
		return
	}

	fmt.Println("Verificação concluída. Arquivos duplicados foram escritos em arquivos_duplicados.txt.")
}
