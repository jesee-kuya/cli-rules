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

// createPR uses GitHub CLI (gh) to open a PR.
func CreatePR(repo, branch, title, body string) (string, error) {
	cmd := exec.Command("gh", "pr", "create",
		"--repo", repo,
		"--head", branch,
		"--title", title,
		"--body", body,
		"--json", "url",
	)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	// gh returns JSON like { "url": "https://..." }
	var resp struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(buf.Bytes(), &resp); err != nil {
		return "", err
	}
	return resp.URL, nil
}

// runQChat invokes `q chat` and returns the raw AI response.
func RunQChat(prompt string) ([]byte, error) {
	// “q chat” is the Amazon Q Developer CLI chat entrypoint :contentReference[oaicite:0]{index=0}
	cmd := exec.Command("q", "chat", "--no-stream", "--message", prompt)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ExitOnErr(what string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s: %v\n", what, err)
		os.Exit(1)
	}
}

