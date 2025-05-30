package helper

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

func ExcutableCommand(cmd string) (bool, string) {
	// system shell
	var shell string
	var shellArg string

	if runtime.GOOS == "windows" {
		shell = "cmd"
		shellArg = "/C"
	} else {
		shell = "sh"
		shellArg = "-c"
	}

	// make command
	command := exec.Command(shell, shellArg, cmd)

	// get output
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	// execute command
	err := command.Run()

	// get result
	if err != nil {
		// execute failed
		errorMsg := stderr.String()
		if errorMsg == "" {
			errorMsg = err.Error()
		}
		return false, strings.TrimSpace(errorMsg)
	}

	// execute success
	return true, strings.TrimSpace(stdout.String())
}
