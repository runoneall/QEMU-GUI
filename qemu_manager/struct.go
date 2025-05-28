package qemu_manager

import (
	"os/exec"
	"sync"
)

type CommandManager struct {
	processes map[string]*exec.Cmd
	mu        sync.Mutex
}
