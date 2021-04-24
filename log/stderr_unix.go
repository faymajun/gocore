// Log the panic under unix to the log file
// +build darwin unix linux

package log

import (
	Log "log"
	"os"
	"syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		Log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
