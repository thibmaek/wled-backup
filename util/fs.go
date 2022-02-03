package util

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	executable := os.Args[1]

	if executable != "wled-backup" {
		cwd, _ := os.Getwd()
		return cwd
	} else {
		cwd, _ := os.Executable()
		return cwd
	}
}

func GetDownloadDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Downloads")
}
