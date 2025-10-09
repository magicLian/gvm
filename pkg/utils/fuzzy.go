package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func FuzzySelect(versions []string, currentVersion string) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "=> {{ . | cyan }}", // 当前光标所在项
		Inactive: "  {{ . }}",
	}

	startIndex := 0
	for i, v := range versions {
		if v == currentVersion {
			startIndex = i
			break
		}
	}

	prompt := promptui.Select{
		Label:     "Select installed Go version",
		Items:     versions,
		Size:      len(versions),
		CursorPos: startIndex,
		Templates: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return "", err
	}

	return result, nil
}
