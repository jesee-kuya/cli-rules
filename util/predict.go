package util

import (
	"encoding/json"
	"fmt"

	"github.com/jesee-kuya/cli-rules/model"
	"github.com/jesee-kuya/cli-rules/repository"
)

// predictConflicts uses AI to predict conflicts before touching the code.
func PredictConflicts(dep model.Dependency) ([]string, error) {
	prompt := fmt.Sprintf(
		"Predict any code-level conflicts when upgrading %s from %s to %s in this Go codebase. "+
			"Return JSON array of strings (function names or file paths likely to break).",
		dep.Name, dep.CurrentVer, dep.LatestVer,
	)
	out, err := repository.RunQChat(prompt)
	if err != nil {
		return nil, err
	}
	var conflicts []string
	if err := json.Unmarshal(out, &conflicts); err != nil {
		return nil, fmt.Errorf("invalid JSON from Q conflict prediction: %w", err)
	}
	return conflicts, nil
}
