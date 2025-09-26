package info

import "fmt"

func Show(info map[string][]string) {
	fmt.Println("Language: ", info["Language"][0])

	if len(info["DockerFiles"]) != 0 {
		for _, file := range info["DockerFiles"] {
			fmt.Println("Dockerfile: ", file)
		}
	} else {
		fmt.Println("Dockerfiles: not found")
	}

	if len(info["DockerComposeFiles"]) != 0 {
		for _, file := range info["DockerComposeFiles"] {
			fmt.Println("Docker compose: ", file)
		}
	} else {
		fmt.Println("Docker compose: not found")
	}

	if len(info["EntryPoints"]) != 0 {
		for _, file := range info["EntryPoints"] {
			fmt.Println("Entry point: ", file)
		}
	} else {
		fmt.Println("Entry points: not found")
	}

	if len(info["PackageManagers"]) != 0 {
		for _, file := range info["PackageManagers"] {
			fmt.Println("Package manager: ", file)
		}
	} else {
		fmt.Println("Package managers: not found")
	}
}
