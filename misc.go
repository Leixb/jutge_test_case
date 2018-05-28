package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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

	// Wait for the process to finish or kill it after a timeout:

	var stdoutBuf bytes.Buffer

	stdoutIn, _ := cmd.StdoutPipe()

	var errStdout error
	//stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stdout := &stdoutBuf
	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	done := make(chan error, 1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(500 * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
		}
		log.Println("process killed as timeout reached")
		return "", errors.New("Timed out")
	case err := <-done:
		if err != nil {
			log.Fatalf("process finished with error = %v", err)
			return "", errors.New("Execution Error")
		}
		log.Print("process finished successfully")
	}

	if errStdout != nil {
		log.Fatal("failed to capture stdout\n")
	}
	outStr := string(stdoutBuf.Bytes())

	return outStr, nil
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
