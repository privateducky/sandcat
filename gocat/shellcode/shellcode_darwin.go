package shellcode

import (
	"fmt"
)

// Runner runner
func Runner(shellcode []byte) bool {
	fmt.Println("[!] Shellcode executor for darwin not available")
	return false
}

// IsAvailable does a shellocode runner exist
func IsAvailable() bool {
	return false
}
