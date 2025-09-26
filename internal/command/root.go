package command

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/deniskorbakov/ivan/configs/envconst"
	"github.com/deniskorbakov/ivan/internal/version"
	"github.com/spf13/cobra"
)

// Run Start app with cobra cmd
func Run() {
	cmd := &cobra.Command{
		Use:     envconst.UseRootCmd,
		Long:    envconst.LongRootCmd,
		Example: envconst.ExampleRootCmd,
	}

	// Disable default options cmd
	cmd.Root().CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(
		buildCmd,
	)

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithVersion(version.Get()),
	); err != nil {
		os.Exit(1)
	}
}
