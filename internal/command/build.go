package command

import (
	"fmt"
	"os"
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

		repDir := fields.RepositoryLocal
		project := repDir

		if fields.RepositoryUrl != "" {
			repDir, err = storage.TempDir()
			if err != nil {
				return err
			}

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

		err = runAnsiblePlaybook(info, fields)
		if err != nil {
			return fmt.Errorf("ansible execution failed: %v", err)
		}

		return nil
	},
}
func runAnsiblePlaybook(info map[string][]string, fields *createBuildForm.BuildForm) error {
	// Базовые аргументы для ansible-playbook
	args := []string{
		"/ansible/playbook.yml",
		"-i", fmt.Sprintf("\"%s,\"", getFirstValue(info, "TargetIP", "62.109.2.242")), // inventory
		"-u", "ansible", // пользователь
	}

	// Добавляем переменные (-e параметры)
	extraVars := map[string]string{
		"target_ip":        getFirstValue(info, "TargetIP", "62.109.2.242"),
		"language":         getFirstValue(info, "Language", ""),
		"package_manager":  getFirstValue(info, "PackageManagers", ""),
		"build_method":     "default",
		"repo_url":         fields.RepositoryUrl,
		"framework":        getFirstValue(info, "Frameworks", ""),
		"entry_point_path": getFirstValue(info, "EntryPoints", ""),
		"project_name":     getFirstValue(info, "Project", ""),
	}

	// Добавляем только непустые переменные
	for key, value := range extraVars {
		if value != "" {
			args = append(args, "-e", fmt.Sprintf("%s=%s", key, value))
		}
	}

	// Создаем команду
	ansibleCmd := exec.Command("ansible-playbook", args...)
	
	// Настраиваем вывод
	ansibleCmd.Stdout = os.Stdout
	ansibleCmd.Stderr = os.Stderr
	ansibleCmd.Stdin = os.Stdin

	fmt.Printf("Executing: ansible-playbook %v\n", args)
	
	// Запускаем команду
	err := ansibleCmd.Run()
	if err != nil {
		return fmt.Errorf("ansible-playbook execution error: %v", err)
	}

	return nil
}

// getFirstValue возвращает первое значение из массива или значение по умолчанию
func getFirstValue(info map[string][]string, key string, defaultValue string) string {
	if values, exists := info[key]; exists && len(values) > 0 {
		return values[0]
	}
	return defaultValue
}

