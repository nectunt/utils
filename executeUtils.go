package main

import (
	"fmt"
	"os"
	"time"

	filesDir "github.com/nectunt/utils/filesDir"
	scrTerm "github.com/nectunt/utils/screenTerminal"
)

func main() {
	// Limpa a tela no início para uma melhor visualização
	scrTerm.ClearScreen()

	// Splash screen com uma coruja ASCII art
	fmt.Println("=====================================================")
	fmt.Println("=                                                   =")
	fmt.Println("=             ,___,                                 =")
	fmt.Println("=           ( o,o )               NECTUNT           =")
	fmt.Println("=            (   )                 UTILS            =")
	fmt.Println("=             '- '                                  =")
	fmt.Println("=                                                   =")
	fmt.Println("=====================================================")

	// Aguarda por 5 segundos
	time.Sleep(5 * time.Second)

	// Limpa a tela novamente para exibir o menu
	scrTerm.ClearScreen()

	// Menu principal
	for {
		fmt.Println("### Menu ###")
		fmt.Println("1 - checkFilesDuplicates")
		fmt.Println("2 - checkFilesDuplicatesV2")
		fmt.Println("3 - listTop10BigestDirectories")
		fmt.Println("4 - listTop10BigestFiles")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		var escolha int
		_, err := fmt.Scanln(&escolha)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			os.Exit(1)
		}

		switch escolha {
		case 1:
			scrTerm.ClearScreen()

			var directory string
			fmt.Print("Digite o diretório a ser verificado: ")
			_, err := fmt.Scanln(&directory)
			if err != nil {
				fmt.Println("Erro ao ler entrada:", err)
			} else {
				if filesDir.DirectoryExists(directory) {
					fmt.Println("Executando checkFilesDuplicates...")
					err := filesDir.CheckFilesDuplicates(directory)
					if err != nil {
						fmt.Println("Erro ao verificar os arquivos duplicados:", err)
					}
				} else {
					fmt.Println("O diretório informado, não existe!")
				}
			}
			os.Exit(1)

		case 2:
			scrTerm.ClearScreen()

			var directory string
			fmt.Print("Digite o diretório a ser verificado: ")
			_, err := fmt.Scanln(&directory)
			if err != nil {
				fmt.Println("Erro ao ler entrada:", err)
			} else {
				if filesDir.DirectoryExists(directory) {
					fmt.Println("Executando checkFilesDuplicates...")
					err := filesDir.CheckFilesDuplicatesV2(directory)
					if err != nil {
						fmt.Println("Erro ao verificar os arquivos duplicados:", err)
					}
				} else {
					fmt.Println("O diretório informado, não existe!")
				}
			}
			os.Exit(1)

		case 3:
			scrTerm.ClearScreen()
			fmt.Println("Executando listTop10BigestDirectories...")
			// Aqui você chamaria a função listTop10BigestDirectories()
			os.Exit(1)

		case 4:
			scrTerm.ClearScreen()
			fmt.Println("Executando listTop10BigestFiles...")
			// Aqui você chamaria a função listTop10BigestFiles()
			os.Exit(1)

		case 0:
			scrTerm.ClearScreen()
			fmt.Println("Saindo do programa.")
			os.Exit(0)

		default:
			scrTerm.ClearScreen()
			fmt.Println("Opção inválida. Por favor, escolha novamente.")
		}

		// Aguarda que o usuário pressione qualquer tecla antes de limpar a tela e exibir o menu novamente
		fmt.Print("Pressione qualquer tecla para voltar ao menu...")
		fmt.Scanln() // Aguarda a entrada do usuário
		scrTerm.ClearScreen()
	}
}
