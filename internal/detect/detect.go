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

var info = map[string][]string{
	"Language":           {},
	"DockerFiles":        {},
	"DockerComposeFiles": {},
	"EntryPoints":        {},
	"PackageManagers":    {},
}

const (
	DockerPattern        = "dockerfile"
	DockerComposePattern = "docker-compose"
	EntryPointPattern    = "main"
)

func Run(files []string) (map[string][]string, error) {
	languageCount := make(map[string]int)

	for _, file := range files {
		filename := filepath.Base(file)
		ext := filepath.Ext(file)

		if fileByPattern(filename, DockerPattern) {
			info["DockerFiles"] = append(info["DockerFiles"], file)
		}
		if fileByPattern(filename, DockerComposePattern) {
			info["DockerComposeFiles"] = append(info["DockerComposeFiles"], file)
		}
		if fileByPattern(filename, EntryPointPattern) {
			info["EntryPoints"] = append(info["EntryPoints"], file)
		}

		fileLanguage := language(ext)
		if fileLanguage != "" {
			languageCount[fileLanguage]++
		}

		pm := packageManager(filename)
		if pm != "" {
			info["PackageManagers"] = append(info["PackageManagers"], pm)
		}
	}

	setMainLanguage(languageCount)

	return info, nil
}

func packageManager(filename string) string {
	if fileByPattern(filename, "uv.lock") {
		return "uv"
	}

	if fileByPattern(filename, "requirements.txt") {
		return "pip"
	}

	if fileByPattern(filename, "composer.json") {
		return "composer"
	}

	if fileByPattern(filename, "package.json") {
		return "npm"
	}

	return ""

}

func language(ext string) string {
	ext = strings.ToLower(ext)
	for lang, exts := range languageExtensions {
		for _, langExt := range exts {
			if ext == langExt {
				return lang
			}
		}
	}

	return ""
}

func setMainLanguage(languageCount map[string]int) {
	if len(languageCount) > 0 {
		mainLanguage := ""
		maxCount := 0
		for lang, count := range languageCount {
			if count > maxCount {
				maxCount = count
				mainLanguage = lang
			}
		}
		info["Language"] = []string{mainLanguage}
	} else {
		info["Language"] = []string{"not found"}
	}
}

func fileByPattern(filename string, pattern string) bool {
	return strings.Contains(strings.ToLower(filename), pattern)
}
