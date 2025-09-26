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

		dockerFiles := detect.DockerFiles(files)
		if len(dockerFiles) == 0 {
			fmt.Println("Dockerfiles: No docker files found")
		} else {
			for _, file := range dockerFiles {
				fmt.Println("Dockerfile: ", file)
			}
		}

		dockerComposeFiles := detect.DockerComposeFiles(files)
		if len(dockerComposeFiles) == 0 {
			fmt.Println("Docker Compose files: No docker files found")
		} else {
			for _, file := range dockerComposeFiles {
				fmt.Println("Docker Compose file: ", file)
			}
		}

		entryPoints := detect.EntryPoints(files)
		if len(entryPoints) == 0 {
			fmt.Println("Entry points: No entry points found")
		} else {
			for _, file := range entryPoints {
				fmt.Println("Entry point: ", file)
			}
		}

		return nil
	},
}
