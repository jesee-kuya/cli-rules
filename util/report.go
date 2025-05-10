package util

import (
	"encoding/json"
	"fmt"

	"github.com/jesee-kuya/cli-rules/model"
	"github.com/jesee-kuya/cli-rules/repository"
)

// generateCompatibilityReport asks AI to detail what needs refactoring.
func GenerateCompatibilityReport(dep model.Dependency, predicted []string) (model.CompatibilityReport, error) {
	payload := map[string]interface{}{
		"dependency": dep,
		"conflicts":  predicted,
	}
	j, _ := json.Marshal(payload)
	prompt := fmt.Sprintf(
		"Based on these predicted conflicts %s, generate a compatibility report for upgrading %s.",
		string(j), dep.Name,
	)
	out, err := repository.RunQChat(prompt)
	if err != nil {
		return model.CompatibilityReport{}, err
	}
	var rpt model.CompatibilityReport
	if err := json.Unmarshal(out, &rpt); err != nil {
		return model.CompatibilityReport{}, fmt.Errorf("invalid JSON from Q compat report: %w", err)
	}
	return rpt, nil
}
