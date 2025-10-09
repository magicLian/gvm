package main

import (
	"fmt"
	"gvm/pkg/commands"
	"gvm/pkg/utils"
	"os"
)

func main() {
	// check command line arguments
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("The Go version is required, for example: gvm install 1.25.1")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Install(version)
	case "use":
		version := ""
		if len(os.Args) == 3 {
			version = os.Args[2]
		}

		if version == "" {
			versions := commands.GetInstalledVersions()
			if len(versions) == 0 {
				fmt.Printf("no installed versions found\n")
				os.Exit(1)
			}
			currentVersion := commands.GetCurrentVersion()

			selected, err := utils.FuzzySelect(versions, currentVersion)
			if err != nil {
				fmt.Printf("Failed to select version: %v\n", err)
				os.Exit(1)
			}
			version = selected
		}

		commands.Use(version)
	case "ls":
		commands.ListInstalled()
	case "ls-remote":
		commands.ListRemote()
	case "uninstall":
		if len(os.Args) < 3 {
			fmt.Println("The Go version is required, for example: gvm uninstall 1.25.1")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Uninstall(version)
	case "--help":
		help()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		help()
		os.Exit(1)
	}
}

// help shows the help message
func help() {
	fmt.Println("gvm - Go Version Manager")
	fmt.Println("\nUsage:")
	fmt.Println("  gvm install <version>    Install the specified Go version")
	fmt.Println("  gvm use <version>        Switch to the specified Go version")
	fmt.Println("  gvm ls                   List installed Go versions")
	fmt.Println("  gvm ls-remote            List remote available Go versions")
	fmt.Println("  gvm uninstall <version>  Uninstall the specified Go version")
	fmt.Println("  gvm --help               Show the help message")
}
