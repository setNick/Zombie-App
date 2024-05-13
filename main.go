package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "5")

	// Запуск процесса
	if err := cmd.Start(); err != nil {
		fmt.Printf("Ошибка при запуске команды: %s\n", err)
		os.Exit(1)
	}

	// Вывод ID дочернего процесса
	fmt.Printf("Дочерний процесс запущен, PID: %d.\n", cmd.Process.Pid)

	go func() {
		cmd.Wait()
		fmt.Printf("Дочерний процесс завершился, PID: %d.\n", cmd.Process.Pid)
	}()
	for {
		fmt.Println("Родительский процесс продолжает работать...")
		time.Sleep(1 * time.Second)
	}
}
