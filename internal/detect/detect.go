package detect

import (
	"path/filepath"
	"strings"
)

var languageExtensions = map[string][]string{
	"Go":         {".go"},
	"Python":     {".py"},
	"JavaScript": {".js", ".jsx", ".ts", ".tsx"},
	"Java":       {".java"},
	"C++":        {".cpp", ".cxx", ".cc", ".c"},
	"C#":         {".cs"},
	"Ruby":       {".rb"},
	"PHP":        {".php"},
	"Rust":       {".rs"},
	"Swift":      {".swift"},
}

const (
	DockerFilename        = "dockerfile"
	DockerComposeFilename = "docker-compose.yml"
)

func Language(files []string) (string, error) {
	var mainLanguage string

	languageCount := make(map[string]int)

	for _, file := range files {
		ext := filepath.Ext(file)

		for lang, exts := range languageExtensions {
			for _, langExt := range exts {
				if ext == langExt {
					languageCount[lang]++
					break
				}
			}
		}
	}

	maxCount := 0
	for lang, count := range languageCount {
		if count > maxCount {
			maxCount = count
			mainLanguage = lang
		}
	}

	if mainLanguage == "" {
		return "unknown", nil
	}

	return mainLanguage, nil
}

func DockerFiles(files []string) ([]string, error) {
	var dockerFiles []string

	for _, file := range files {
		filename := filepath.Base(file)

		if strings.ToLower(filename) == DockerFilename {
			dockerFiles = append(dockerFiles, file)
		}
	}

	return dockerFiles, nil
}

func DockerComposeFiles(files []string) ([]string, error) {
	var dockerComposeFiles []string

	for _, file := range files {
		filename := filepath.Base(file)

		if strings.ToLower(filename) == DockerComposeFilename {
			dockerComposeFiles = append(dockerComposeFiles, file)
		}
	}

	return dockerComposeFiles, nil
}
