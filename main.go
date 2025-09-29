package main

import (
	"fmt"
	"gvm/pkg/commands"
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
		if len(os.Args) < 3 {
			fmt.Println("The Go version is required, for example: gvm use 1.25.1")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Use(version)
	case "ls":
		commands.ListInstalled()
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
	fmt.Println("  gvm ls                 	List installed Go versions")
	fmt.Println("  gvm uninstall <version>  Uninstall the specified Go version")
	fmt.Println("  gvm --help               Show the help message")
}
