package storage

import (
	"log"
	"os"
)

func TempDir() (string, error) {
	repDir, err := os.MkdirTemp("", "dir")
	if err != nil {
		return "", err
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Printf("Failed to remove temporary directory: %v", err)
		}
	}(repDir)

	return repDir, nil
}
