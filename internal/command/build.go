package command

import (
	"fmt"

	createBuildForm "github.com/deniskorbakov/ivan/internal/component/build"
	"github.com/deniskorbakov/ivan/internal/storage"

	"github.com/deniskorbakov/ivan/configs/envconst"
	"github.com/deniskorbakov/ivan/internal/detect"
	"github.com/deniskorbakov/ivan/internal/repository"
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

		language, err := detect.Language(files)
		if err != nil {
			return err
		}

		fmt.Println("Language: ", language)

		dockerFiles, err := detect.DockerFiles(files)
		if err != nil {
			return err
		}

		if len(dockerFiles) == 0 {
			fmt.Println("Dockerfiles: No docker files found")
		} else {
			for _, file := range dockerFiles {
				fmt.Println("Dockerfile: ", file)
			}
		}

		dockerComposeFiles, err := detect.DockerComposeFiles(files)
		if err != nil {
			return err
		}

		if len(dockerComposeFiles) == 0 {
			fmt.Println("Docker Compose files: No docker files found")
		} else {
			for _, file := range dockerComposeFiles {
				fmt.Println("Docker Compose file: ", file)
			}
		}

		return nil
	},
}
