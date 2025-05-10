package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jesee-kuya/cli-rules/model"
)

// upgradeDependency edits go.mod and runs `go get name@version`.
func UpgradeDependency(path string, dep model.Dependency) error {
	cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", dep.Name, dep.LatestVer))
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
