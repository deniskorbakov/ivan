package build

import (
	"errors"
	"net/url"
	"os"
	"strings"
)

func validateRepository(repository string) error {
	_, err := url.ParseRequestURI(repository)
	if err != nil {
		return errors.New("url is invalid")
	}

	return nil
}

func validateLocalRepository(repository string) error {
	repository = strings.ReplaceAll(repository, " ", "")

	errorExists := errors.New("repository does not exist")
	errorNotDir := errors.New("repository is not a directory")

	info, err := os.Stat(repository)
	if errors.Is(err, os.ErrNotExist) {
		return errorExists
	}

	if info.IsDir() {
		return nil
	}

	return errorNotDir
}
