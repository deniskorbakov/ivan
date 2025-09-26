package detect

import (
	"os"
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
	"Frameworks":         {},
	"Environments":       {},
}

const (
	DockerPattern        = "dockerfile"
	DockerComposePattern = "docker-compose"
	EntryPointPattern    = "main"
	EnvPattern           = "env"
)

func Run(files []string, dir string) (map[string][]string, error) {
	languageCount := make(map[string]int)

	for _, file := range files {
		relativePath := strings.ReplaceAll(file, dir, "")
		filename := filepath.Base(file)
		ext := filepath.Ext(file)

		if fileByPattern(filename, DockerPattern) {
			info["DockerFiles"] = append(info["DockerFiles"], relativePath)
		}
		if fileByPattern(filename, DockerComposePattern) {
			info["DockerComposeFiles"] = append(info["DockerComposeFiles"], relativePath)
		}
		if fileByPattern(filename, EntryPointPattern) {
			info["EntryPoints"] = append(info["EntryPoints"], relativePath)
		}
		if fileByPattern(filename, EnvPattern) {
			info["Environments"] = append(info["Environments"], relativePath)
		}

		fileLanguage := language(ext)
		if fileLanguage != "" {
			languageCount[fileLanguage]++
		}

		pm := packageManager(filename, file)
		if pm != "" {
			info["PackageManagers"] = append(info["PackageManagers"], pm)
		}
	}

	setMainLanguage(languageCount)

	return info, nil
}

func packageManager(filename string, fullPath string) string {
	if fileByPattern(filename, "uv.lock") {
		if searchInFile(fullPath, "django") {
			info["Frameworks"] = append(info["Frameworks"], "Django")
		}

		if searchInFile(fullPath, "fastapi") {
			info["Frameworks"] = append(info["Frameworks"], "FastAPI")
		}

		return "uv"
	}

	if fileByPattern(filename, "requirements.txt") {
		if searchInFile(fullPath, "django") {
			info["Frameworks"] = append(info["Frameworks"], "Django")
		}

		if searchInFile(fullPath, "fastapi") {
			info["Frameworks"] = append(info["Frameworks"], "FastAPI")
		}

		return "pip"
	}

	if fileByPattern(filename, "composer.json") {
		if searchInFile(fullPath, "laravel") {
			info["Frameworks"] = append(info["Frameworks"], "Laravel")
		}

		if searchInFile(fullPath, "symfony") {
			info["Frameworks"] = append(info["Frameworks"], "Symfony")
		}

		if searchInFile(fullPath, "yii2") {
			info["Frameworks"] = append(info["Frameworks"], "Yii2")
		}

		return "composer"
	}

	if fileByPattern(filename, "package.json") {
		if searchInFile(fullPath, "vue") {
			info["Frameworks"] = append(info["Frameworks"], "Vue")
		}

		if searchInFile(fullPath, "react") {
			info["Frameworks"] = append(info["Frameworks"], "React")
		}

		if searchInFile(fullPath, "angular") {
			info["Frameworks"] = append(info["Frameworks"], "Angular")
		}

		return "npm"
	}

	return ""

}

func searchInFile(file string, search string) bool {
	content, err := os.ReadFile(file)
	if err == nil {
		fileContent := string(content)

		return strings.Contains(fileContent, search)
	}

	return false
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
