package ansible

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	createBuildForm "github.com/deniskorbakov/ivan/internal/component/build"
	"github.com/deniskorbakov/ivan/internal/component/output"
)

func Exec(info map[string][]string, fields *createBuildForm.Fields) error {
	args := []string{
		"ansible/playbook.yml",
		"-i", fmt.Sprintf("%s,", "62.109.2.242"),
		"-u", "ansible",
	}

	extraVars := map[string]string{
		"language":        getFirstValue(info, "Language", "null"),
		"package_manager": getFirstValue(info, "PackageManagers", "null"),
		"build_method": func() string {
			dockerfile := getFirstValue(info, "DockerComposeFiles", "null")
			if dockerfile != "null" {
				return "docker-compose"
			}

			return "null"
		}(),
		"repo_url":  fields.RepositoryUrl,
		"framework": getFirstValue(info, "Frameworks", "null"),
		"entry_point_path": func() string {
			frameWork := getFirstValue(info, "Frameworks", "null")
			language := getFirstValue(info, "Language", "null")

			if frameWork == "django" {
				managePy := getFirstValue(info, "ManagePy", "null")
				if managePy != "null" {
					return managePy
				}
			}

			if language == "go" {
				if len(info["EntryPoints"]) != 0 {
					for _, entryPoint := range info["EntryPoints"] {
						ext := filepath.Ext(entryPoint)

						if ext == ".go" {
							return entryPoint
						}
					}
				}
			}

			return getFirstValue(info, "EntryPoints", "null")
		}(),
		"project_name": "null",
		"env_path": func() string {
			envs := info["Environments"]
			if len(envs) == 0 {
				return "null"
			}

			for _, file := range envs {
				filename := strings.ToLower(filepath.Base(file))
				if filename == ".env" {
					return file
				}
			}

			return getFirstValue(info, "Environments", "null")
		}(),
	}

	for key, value := range extraVars {
		if value != "" {
			args = append(args, "-e", fmt.Sprintf("%s=%s", key, value))
		}
	}

	fmt.Println("\n\n", output.Blue("---Run Ansible--"))

	ansibleCmd := exec.Command("ansible-playbook", args...)

	fmt.Println(args)

	ansibleCmd.Stdout = os.Stdout
	ansibleCmd.Stderr = os.Stderr
	ansibleCmd.Stdin = os.Stdin

	err := ansibleCmd.Run()
	if err != nil {
		return fmt.Errorf("ansible-playbook execution error: %v", err)
	}

	return nil
}

func getFirstValue(info map[string][]string, key string, defaultValue string) string {
	if values, exists := info[key]; exists && len(values) > 0 {
		return strings.ToLower(values[0])
	}

	return strings.ToLower(defaultValue)
}
