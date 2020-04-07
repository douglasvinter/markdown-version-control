package model

// Markdown - Expected request model for a Markdown update after GIT WebHook action
type Markdown struct {
	Commit   string `json:"commitId"`
	Build    string `json:"buildNumber"`
	Markdown string `json:"payload"`
}

// GitProject - List of saved markdown projects
type GitProject struct {
	ProjectID string   `json:"projectId"`
	Data      Markdown `json:"markdown"`
}
