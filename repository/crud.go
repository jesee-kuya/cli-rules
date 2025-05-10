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

// runGitBranch creates or checks out the feature branch.
func RunGitBranch(name string) error {
	return exec.Command("git", "checkout", "-B", name).Run()
}

// runGitAddAll stages all changes.
func RunGitAddAll() error {
	return exec.Command("git", "add", "-A").Run()
}

// runGitCommit makes a commit.
func RunGitCommit(msg string) error {
	return exec.Command("git", "commit", "-m", msg).Run()
}

// runGitPush pushes branch to origin.
func RunGitPush(branch string) error {
	return exec.Command("git", "push", "--set-upstream", "origin", branch).Run()
}

/