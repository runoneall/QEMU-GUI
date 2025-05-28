package qemu_manager

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"
)

var cmdManager = CommandManager{
	processes: make(map[string]*exec.Cmd),
}

func RunCommand(commandID string, commandStr string) error {
	cmdParts := strings.Fields(commandStr)
	if len(cmdParts) == 0 {
		return errors.New("empty command")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	cmdManager.mu.Lock()
	cmdManager.processes[commandID] = cmd
	cmdManager.mu.Unlock()

	go func() {
		if err := cmd.Start(); err != nil {
			cmdManager.removeCommand(commandID)
			return
		}

		cmd.Wait()
		cmdManager.removeCommand(commandID)
	}()

	return nil
}

func StopCommand(commandID string) error {
	cmdManager.mu.Lock()
	defer cmdManager.mu.Unlock()

	if cmd, exists := cmdManager.processes[commandID]; exists && cmd.Process != nil {
		return cmd.Process.Signal(syscall.SIGINT)
	}
	return errors.New("command not found")
}

func (cm *CommandManager) removeCommand(commandID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.processes, commandID)
}
