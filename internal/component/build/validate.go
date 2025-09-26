package build

import (
	"net/url"
)

func validateRepository(repository string) error {
	_, err := url.ParseRequestURI(repository)
	return err
}
