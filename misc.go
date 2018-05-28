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
	if !check_code(prog) {
		return "", errors.New("Invalid code")
	}
	_, err := os.Stat("./problemes/" + file)
	if err != nil {
		return "", errors.New("Program not in database")
	}
	cmd := exec.Command("./problemes/" + file)
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New("Entrada Invàlida")
	}
	return string(out), nil
}

func check_code(code string) bool {
	if len(code) != 9 {
		return false
	}
	return code_matcher.MatchString(code)
}

func get_name(code string) string {

	name, ok := problem_names[code[:6]].(string)
	if !ok {
		return ""
	}
	return name
}
