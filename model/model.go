package model

// Dependency represents one module dependency and its versions.
type Dependency struct {
	Name         string `json:"name"`
	CurrentVer   string `json:"currentVersion"`
	LatestVer    string `json:"latestVersion"`
	ShouldUpdate bool
}

// CompatibilityReport holds AI‑generated notes about an upgrade.
type CompatibilityReport struct {
	Dependency string   `json:"dependency"`
	Notes      []string `json:"notes"`
	Conflicts  []string `json:"predictedConflicts"`
}
