package models

type HealthResult struct {
	URL        string `json:"url"`
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error,omitempty"`
	CheckedAt  string `json:"checkedAt"`
}
