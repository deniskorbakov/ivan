package build

import (
	"github.com/charmbracelet/huh"
)

// Run Main function that runs an interactive form to collect SSH connection details
func Run() (*Fields, error) {
	fields := &Fields{}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Repository").
				Description("Your remote git repository").
				Validate(validateRepository).
				Value(&fields.Repository),
		),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, err
	}

	return fields, nil
}
