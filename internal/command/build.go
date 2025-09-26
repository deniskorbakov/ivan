package command

import (
	"fmt"

	createBuildForm "github.com/deniskorbakov/ivan/internal/component/build"

	"github.com/deniskorbakov/ivan/configs/envconst"
	"github.com/spf13/cobra"
)

// buildCmd Command for build pipeline
var buildCmd = &cobra.Command{
	Use:   envconst.UseBuildCmd,
	Short: envconst.ShortBuildCmd,
	Long:  envconst.LongRootCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		fields, err := createBuildForm.Run()
		if err != nil {
			return err
		}

		fmt.Println("Your get repository: ", fields.Repository)

		return nil
	},
}
