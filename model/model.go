package model

// Dependency represents one module dependency and its versions.
type Dependency struct {
	Name         string `json:"name"`
	CurrentVer   string `json:"currentVersion"`
	LatestVer    string `json:"latestVersion"`
	ShouldUpdate bool
}


