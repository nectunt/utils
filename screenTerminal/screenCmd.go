package utils

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") // Para Windows, use "cmd" e "/c" comandos para executar "cls"
	cmd.Stdout = os.Stdout
	cmd.Run()
}
