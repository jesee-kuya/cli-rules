package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// runTests executes `go test ./...`.
func RunTests(path string) error {
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

