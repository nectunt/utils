package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var dir string
var workers int

type Result struct {
	file   string
	sha256 [32]byte
}

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
func CheckFilesDuplicates(path string) error {
	// Verificar arquivos duplicados
	duplicates, err := checkDuplicates(path)
	if err != nil {
		// fmt.Println("Erro ao verificar arquivos duplicados:", err)
		return err

	}

	// Escrever arquivos duplicados em um arquivo de texto
	err = writeDuplicatesToFile(path, duplicates)
	if err != nil {
		//fmt.Println("Erro ao escrever arquivo de texto:", err)
		return err
	}

	fmt.Println("Verificação concluída. Arquivos duplicados foram escritos em arquivos_duplicados.txt.")

	return nil
}

// checkFilesDuplicatesV2 - uso de todos os processadores disponíveis
func worker(input chan string, results chan<- *Result, wg *sync.WaitGroup) {
	for file := range input {
		var h = sha256.New()
		var sum [32]byte
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if _, err = io.Copy(h, f); err != nil {
			fmt.Fprintln(os.Stderr, err)
			f.Close()
			continue
		}
		f.Close()
		copy(sum[:], h.Sum(nil))
		results <- &Result{
			file:   file,
			sha256: sum,
		}
	}
	wg.Done()
}

func search(input chan string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else if info.Mode().IsRegular() {
			input <- path
		}
		return nil
	})
	close(input)
}

func CheckFilesDuplicatesV2(directory string) error {

	if !DirectoryExists(directory) {
		return fmt.Errorf("Diretório informado não existe!")
	}

	flag.StringVar(&dir, "dir", directory, "directory to search")
	flag.IntVar(&workers, "workers", runtime.NumCPU(), "number of workers")
	flag.Parse()

	fmt.Printf("Buscando em %s usando %d processos...\n", dir, workers)

	input := make(chan string)
	results := make(chan *Result)

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(input, results, &wg)
	}

	go search(input)
	go func() {
		wg.Wait()
		close(results)
	}()

	counter := make(map[[32]byte][]string)
	for result := range results {
		counter[result.sha256] = append(counter[result.sha256], result.file)
	}

	// Busca o diretório atual de execução do programa
	directoryWorkDir, err := GetWorkDir()
	if err != nil {
		return err
	}

	// Criar ou abrir arquivo para escrita
	file, err := os.Create(directoryWorkDir + "arquivos_duplicados.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for sha, files := range counter {
		if len(files) > 1 {

			//fmt.Printf("Encontrado %d duplicados de %s: \n", len(files), hex.EncodeToString(sha[:]))

			_, err := file.WriteString(fmt.Sprintf("Encontrado %d duplicados de %s: \n", len(files), hex.EncodeToString(sha[:])))
			if err != nil {
				return err
			}

			for _, f := range files {
				//fmt.Println("-> ", f)

				_, err := file.WriteString(fmt.Sprintf("-> %s\n", f))
				if err != nil {
					return err
				}
			}
		}
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil

}
