package utils

import (
	"testing"
)

func TestFetchGoVersionsLinuxAmd64(t *testing.T) {
	versions, err := FetchGoVersions("linux", "amd64")
	if err != nil {
		t.Fatal(err)
	}
	if len(versions) == 0 {
		t.Fatal("no versions found")
	}
	for _, v := range versions {
		t.Logf("%s %s %s %s %s %s", v.Version, v.File, v.OS, v.Arch, v.URL, v.Size)
	}
}

func TestFetchGoVersionsWindowsAmd64(t *testing.T) {
	versions, err := FetchGoVersions("windows", "amd64")
	if err != nil {
		t.Fatal(err)
	}
	if len(versions) == 0 {
		t.Fatal("no versions found")
	}
	for _, v := range versions {
		t.Logf("%s %s %s %s %s %s", v.Version, v.File, v.OS, v.Arch, v.URL, v.Size)
	}
}

func TestFetchGoVersionsDarwinArm64(t *testing.T) {
	versions, err := FetchGoVersions("darwin", "arm64")
	if err != nil {
		t.Fatal(err)
	}
	if len(versions) == 0 {
		t.Fatal("no versions found")
	}
	for _, v := range versions {
		t.Logf("%s %s %s %s %s %s", v.Version, v.File, v.OS, v.Arch, v.URL, v.Size)
	}
}
