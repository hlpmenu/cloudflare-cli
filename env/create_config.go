package env

import (
	_ "embed"
	"go-debug/output"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

//go:embed config.conf
var configContent []byte // or string if you prefer

// Example function to use the embedded content
func UseConfig() {
	// Use configContent directly as []byte or string
	// For example, writing it to a file:
	err := os.WriteFile("config.conf", configContent, fs.ModePerm)
	if err != nil {
		if !hasPermission(".") {
			output.Error("Could not create config config file")
			output.Exit("Exiting...")
		}
		output.Error("Could not create config file")
		output.Exit("Exiting...")
	}
}
func hasPermission(path string) bool {
	d, err := os.Stat(path)
	if err != nil {
		return false
	}

	stat, ok := d.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	currentUser, err := user.Current()
	if err != nil {
		return false
	}

	// Convert currentUser.Uid to uint32
	currentUserID, err := strconv.ParseUint(currentUser.Uid, 10, 32)
	if err != nil {
		return false
	}

	// Check if the current user is the owner
	if stat.Uid == uint32(currentUserID) {
		// Check write permission for the owner
		return d.Mode().Perm()&0200 != 0
	}

	return false
}
