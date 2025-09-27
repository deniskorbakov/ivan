package info

import (
	"fmt"

	"github.com/deniskorbakov/ivan/internal/component/output"
)

func Show(info map[string][]string) {
	fmt.Println(output.Green("---System Info---"))

	fmt.Println(output.Green(info["Project"][0]))

	fmt.Println("\n", output.Green("---System Info---"))

	fmt.Println(output.Pink(info["Language"][0]))

	fmt.Println("\n", output.Green("---Dockerfiles Info---"))

	if len(info["DockerFiles"]) != 0 {
		for _, file := range info["DockerFiles"] {
			fmt.Println(output.Blue(file))
		}
	} else {
		fmt.Println(output.Red("Add Dockerfiles to your project for assembly"))
	}

	fmt.Println("\n", output.Green("---Docker Compose Info---"))

	if len(info["DockerComposeFiles"]) != 0 {
		for _, file := range info["DockerComposeFiles"] {
			fmt.Println(output.BlueExtra(file))
		}
	} else {
		fmt.Println(output.Red("To control docker add docker compose"))
	}

	fmt.Println("\n", output.Green("---Entry Points Info---"))

	if len(info["EntryPoints"]) != 0 {
		for _, file := range info["EntryPoints"] {
			fmt.Println(output.Orange(file))
		}
	} else {
		fmt.Println(output.Red("Make entry points to your app"))
	}

	fmt.Println("\n", output.Green("---Package Manager Info---"))

	if len(info["PackageManagers"]) != 0 {
		for _, file := range info["PackageManagers"] {
			fmt.Println(output.WhiteRed(file))
		}
	} else {
		fmt.Println(output.Red("Add a package manager"))
	}

	fmt.Println("\n", output.Green("---Frameworks Info---"))

	if len(info["Frameworks"]) != 0 {
		for _, name := range info["Frameworks"] {
			fmt.Println(output.Purple(name))
		}
	} else {
		fmt.Println(output.Red("Try using frameworks"))
	}

	fmt.Println("\n", output.Green("---Environments Info---"))

	if len(info["Environments"]) != 0 {
		for _, file := range info["Environments"] {
			fmt.Println(output.WhiteOrange(file))
		}
	} else {
		fmt.Println(output.Red("For dependencies in the project add an env file"))
	}
}
