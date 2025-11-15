package storage

import (
	"os"
)

func TempDir() (string, error) {
	repDir, err := os.MkdirTemp("", "dir")
	if err != nil {
		return "", err
	}

	return repDir, nil
}

func DeleteDir(dir string) error {
	return os.RemoveAll(dir)
}
