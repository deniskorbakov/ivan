package command

import (
	"fmt"
	"log"

	createBuildForm "github.com/deniskorbakov/ivan/internal/component/build"
	infoComponent "github.com/deniskorbakov/ivan/internal/component/info"

	"github.com/deniskorbakov/ivan/internal/ansible"
	"github.com/deniskorbakov/ivan/internal/component/output"

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

		repDir := fields.RepositoryLocal
		project := repDir

		if fields.RepositoryUrl != "" {
			repDir, err = storage.TempDir()
			if err != nil {
				return err
			}
			defer func() {
				err := storage.DeleteDir(repDir)
				if err != nil {
					log.Fatal(err)
				}
			}()

			err = repository.Save(repDir, fields.RepositoryUrl)
			if err != nil {
				return err
			}

			project = fields.RepositoryUrl
		}

		files, err := repository.Files(repDir)
		if err != nil {
			return err
		}

		info, err := detect.Run(files, repDir)
		if err != nil {
			return err
		}

		info["Project"] = append(info["Project"], project)

		infoComponent.Show(info)

		err = ansible.Exec(info, fields)
		if err != nil {
			return fmt.Errorf("ansible execution failed: %v", output.Red(err.Error()))
		}

		return nil
	},
}
