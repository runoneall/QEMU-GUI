package qemu_manager

import (
	"os/exec"
	"strings"
)

type vmUUID string

type vmRunner struct {
	OSExec map[vmUUID]*exec.Cmd
}

func (vr *vmRunner) DeleteVMRunner(uuid vmUUID) {
	if cmd, ok := vr.OSExec[uuid]; ok {
		cmd.Process.Kill()
	}
	delete(vr.OSExec, uuid)
}

func (vr *vmRunner) StartVMRunner(uuid vmUUID, cmd *exec.Cmd) {
	vr.OSExec[uuid] = cmd
	cmd.Run()
	vr.DeleteVMRunner(uuid)
}

var RunnerPool = vmRunner{
	OSExec: map[vmUUID]*exec.Cmd{},
}

func StartVM(uuid string, command string) {
	cmds := strings.Fields(command)
	cmd := exec.Command(cmds[0], cmds[1:]...)
	go func() {
		RunnerPool.StartVMRunner(vmUUID(uuid), cmd)
	}()
}

func DeleteVM(uuid string) {
	RunnerPool.DeleteVMRunner(vmUUID(uuid))
}
