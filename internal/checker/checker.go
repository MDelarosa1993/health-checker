package checker

import (
	"context"
	"net/http"
	"time"

	"health-monitor/internal/models"
)

type Checker struct {
	Client *http.Client
}

func NewChecker() Checker {
	return Checker{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c Checker) CheckURL(ctx context.Context, url string) models.HealthResult {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.HealthResult{
			URL:       url,
			Status:    "DOWN",
			Error:     err.Error(),
			CheckedAt: time.Now().UTC().Format(time.RFC3339),
		}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return models.HealthResult{
			URL:       url,
			Status:    "DOWN",
			Error:     err.Error(),
			CheckedAt: time.Now().UTC().Format(time.RFC3339),
		}
	}
	defer resp.Body.Close()

	status := "UP"

	if resp.StatusCode >= 400 {
		status = "DOWN"
	}

	return models.HealthResult{
		URL:        url,
		Status:     status,
		StatusCode: resp.StatusCode,
		CheckedAt:  time.Now().UTC().Format(time.RFC3339),
	}
}
