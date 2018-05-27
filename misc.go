package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func test(input, prog string) (string, error) {
	file := filepath.Base(prog)
	_, err := os.Stat("./problemes/" + file)
	if err != nil {
		return "", errors.New("Invalid code")
	}
	cmd := exec.Command("./problemes/" + file)
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New("Entrada Invàlida")
	}
	return string(out), nil
}
