// +build darwin

package shells

import (
	"../execute"
	"os/exec"
)

type Osascript struct {
	shortName string
	path string
	execArgs []string
}

func init() {
	shell := &Osascript{
		shortName: "osa",
		path: "osascript",
		execArgs: []string{"-e"},
	}
	if shell.CheckIfAvailable() {
		execute.Executors[shell.path] = shell
	}
}

func (o *Osascript) Run(command string, timeout int) ([]byte, string, string) {
	return runShellExecutor(*exec.Command(o.path, append(o.execArgs, command)...), timeout)
}

func (o *Osascript) String() string {
	return o.shortName
}

func (o *Osascript) CheckIfAvailable() bool {
	return checkExecutorInPath(o.path)
} 