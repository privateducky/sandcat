package shellcode

import (
	"fmt"
)

// Runner runner
func Runner(shellcode []byte) bool {
	fmt.Println("[!] Shellcode executor for linux not available")
	return false
}

// IsAvailable does a shellcode runner exist
func IsAvailable() bool {
	return false
}
