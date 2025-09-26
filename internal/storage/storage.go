package storage

import "os"

func TempDir() (string, error) {
	repDir, err := os.MkdirTemp("", "dir")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(repDir)

	return repDir, nil
}
