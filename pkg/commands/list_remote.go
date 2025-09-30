package commands

import (
	"fmt"
	"gvm/pkg/utils"
	"runtime"
	"slices"
)

func ListRemote() {
	fmt.Println("Listing remote Go versions...")
	fmt.Println("")

	goRemoteVersions, err := utils.FetchGoVersions(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		fmt.Printf("Fetch Go remote versions failed: %v\n", err)
		return
	}

	// Collect version list.
	installedVersions := GetInstalledVersions()
	currentVersion := GetCurrentVersion()

	for _, remoteVersion := range goRemoteVersions {
		isCurrent := false
		if remoteVersion.Version == currentVersion {
			isCurrent = true
		}
		if slices.Contains(installedVersions, remoteVersion.Version) {
			if isCurrent {
				fmt.Printf("  => %s (installed) ✅ \n", remoteVersion.Version)
			} else {
				fmt.Printf("    %s (installed) ✅ \n", remoteVersion.Version)
			}
		} else {
			fmt.Printf("    %s\n", remoteVersion.Version)
		}
	}
}
