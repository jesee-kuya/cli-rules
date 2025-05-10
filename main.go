// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jesee-kuya/cli-rules/repository"
	"github.com/jesee-kuya/cli-rules/util"
)

func main() {
	projectPath := flag.String("path", ".", "path to the Go project")
	githubRepo := flag.String("repo", "", "GitHub repo in owner/name format (required to open PR)")
	flag.Parse()

	if *githubRepo == "" {
		fmt.Fprintln(os.Stderr, "Error: -repo is required")
		os.Exit(1)
	}

	// 1. scan for outdated deps
	deps, err := util.ScanDependencies(*projectPath)
	repository.ExitOnErr("scanning dependencies", err)

	// 2. ask AI to predict conflicts before making changes
	for i := range deps {
		if deps[i].LatestVer != "" && deps[i].LatestVer != deps[i].CurrentVer {
			deps[i].ShouldUpdate = true
			report, err := util.PredictConflicts(deps[i])
			repository.ExitOnErr("predicting conflicts", err)
			// 3. generate full compatibility report
			compRpt, err := util.GenerateCompatibilityReport(deps[i], report)
			repository.ExitOnErr("generating compatibility report", err)
			// 4. apply upgrade if safe
			if len(compRpt.Conflicts) == 0 {
				err = util.UpgradeDependency(*projectPath, deps[i])
				repository.ExitOnErr("upgrading dependency", err)
			} else {
				fmt.Printf("Skipping %s: predicted conflicts: %v\n", deps[i].Name, compRpt.Conflicts)
			}
		}
	}

	// 5. run tests
	if err := repository.RunTests(*projectPath); err != nil {
		fmt.Fprintln(os.Stderr, "Tests failed—aborting PR creation")
		os.Exit(1)
	}

	// 6. commit, push, and open PR
	branch := "auto-dep-updates"
	repository.Must(repository.RunGitBranch(branch))
	repository.Must(repository.RunGitAddAll())
	repository.Must(repository.RunGitCommit("chore: automated dependency upgrades"))
	repository.Must(repository.RunGitPush(branch))
	prURL, err := repository.CreatePR(*githubRepo, branch, "Automated dependency upgrades", "This PR upgrades dependencies as per AI analysis.")
	repository.ExitOnErr("creating PR", err)

	fmt.Println("PR created:", prURL)
}
