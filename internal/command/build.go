package command

import (
	"fmt"

	createBuildForm "github.com/deniskorbakov/ivan/internal/component/build"
	infoComponent "github.com/deniskorbakov/ivan/internal/component/info"

	"github.com/deniskorbakov/ivan/configs/envconst"
	"github.com/deniskorbakov/ivan/internal/detect"
	"github.com/deniskorbakov/ivan/internal/repository"
	"github.com/deniskorbakov/ivan/internal/storage"
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

		fmt.Println("Repository: ", fields.Repository)

		repDir, err := storage.TempDir()
		if err != nil {
			return err
		}

		err = repository.Save(repDir, fields.Repository)
		if err != nil {
			return err
		}

		files, err := repository.Files(repDir)
		if err != nil {
			return err
		}

		info, err := detect.Run(files, repDir)
		if err != nil {
			return err
		}

		infoComponent.Show(info)

		return nil
	},
}
