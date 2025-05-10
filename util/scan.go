package util

import (
	"encoding/json"
	"fmt"

	"github.com/jesee-kuya/cli-rules/model"
	"github.com/jesee-kuya/cli-rules/repository"
)

// scanDependencies uses `q chat` to ask Amazon Q to list outdated deps as JSON.
func ScanDependencies(path string) ([]model.Dependency, error) {
	prompt := fmt.Sprintf(
		"Scan the Go project at %s for outdated module dependencies. "+
			"Return JSON array of {name, currentVersion, latestVersion}.",
		path,
	)
	out, err := repository.RunQChat(prompt)
	if err != nil {
		return nil, err
	}
	var deps []model.Dependency
	if err := json.Unmarshal(out, &deps); err != nil {
		return nil, fmt.Errorf("invalid JSON from Q scan: %w", err)
	}
	return deps, nil
}
