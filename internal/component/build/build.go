package build

import (
	"github.com/charmbracelet/huh"
)

// Run Main function that runs an interactive form to collect SSH connection details
func Run() (*Fields, error) {
	var urlRepConfirm bool

	fields := &Fields{}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Type clone").
				Description("Choice method clone repository").
				Affirmative("Url").
				Negative("Local path").
				Value(&urlRepConfirm),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Repository").
				Description("Url git repository").
				Validate(validateRepository).
				Value(&fields.RepositoryUrl),
		).WithHideFunc(func() bool {
			return !urlRepConfirm
		}),
		huh.NewGroup(
			huh.NewInput().
				Title("Repository").
				Description("Local git repository").
				Validate(validateLocalRepository).
				Value(&fields.RepositoryLocal),
		).WithHideFunc(func() bool {
			return urlRepConfirm
		}),
		huh.NewGroup(
			huh.NewInput().
				Title("Remote User").
				Description("ssh user for deploy ansible").
				Value(&fields.RemoteUser),
			huh.NewInput().
				Title("Remote Host").
				Description("remote host for deploy ansible").
				Value(&fields.RemoteHost),
		),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, err
	}

	return fields, nil
}
