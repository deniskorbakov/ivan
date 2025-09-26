package repository

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

var skipDirs = []string{
	".git",
	".gitignore",
	".gitmodules",
	".gitattributes",
	".gitkeep",
	"vendor",
}

func Save(dir string, RepositoryUrl string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: RepositoryUrl,
	})

	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func Files(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, skipDir := range skipDirs {
				if info.Name() == skipDir {
					return filepath.SkipDir
				}
			}

			return nil
		}

		relativePath := strings.ReplaceAll(path, dir, "")
		files = append(files, relativePath)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
