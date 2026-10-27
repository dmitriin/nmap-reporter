package dispatch

import (
	"fmt"
	"os"
	"os/exec"
)

const prefix = "nmap-reporter-"

// TryDispatch ищет в PATH исполняемый файл с именем prefix+subcommand.
// Если находит — запускает его, передавая оставшиеся аргументы.
// Возвращает true, если подкоманда была найдена и выполнена.
func TryDispatch(subcommand string, args []string) (bool, error) {
	if subcommand == "" {
		return false, nil
	}

	binaryName := prefix + subcommand
	path, err := exec.LookPath(binaryName)
	if err != nil {
		return false, nil // не нашли — не ошибка, просто нет расширения
	}

	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		return true, fmt.Errorf("failed to run %s: %w", binaryName, err)
	}
	return true, nil
}