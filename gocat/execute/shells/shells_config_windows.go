package shells

import "syscall"

func getPlatformSysProcAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: true}
}