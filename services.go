package windows

import (
	"os/exec"
	"strings"
)

// Verifica se o serviço está em execução
func isServiceRunning(serviceName string) (bool, error) {
	// Executa o comando "sc query" para obter o status do serviço
	cmd := exec.Command("sc", "query", serviceName)
	stdout, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// Verifica se o serviço está no estado "running"
	return strings.Contains(string(stdout), "RUNNING"), nil
}

// Inicia o serviço
func startService(serviceName string) error {
	// Executa o comando "sc start" para iniciar o serviço
	cmd := exec.Command("sc", "start", serviceName)
	return cmd.Run()
}

// Interrompe o serviço
func stopService(serviceName string) error {
	// Executa o comando "sc stop" para interromper o serviço
	cmd := exec.Command("sc", "stop", serviceName)
	return cmd.Run()
}

// Pausa o serviço
func pauseService(serviceName string) error {
	// Executa o comando "sc pause" para pausar o serviço
	cmd := exec.Command("sc", "pause", serviceName)
	return cmd.Run()
}

// Reinicia o serviço
func restartService(serviceName string) error {
	// Executa o comando "sc restart" para reiniciar o serviço
	cmd := exec.Command("sc", "restart", serviceName)
	return cmd.Run()
}
