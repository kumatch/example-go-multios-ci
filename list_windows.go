package main

import (
	"os/exec"
	"strings"
)

func outputFiles(dir string) (string, error) {
	out, err := exec.Command("cmd", "/C", "dir", "/B", dir).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
