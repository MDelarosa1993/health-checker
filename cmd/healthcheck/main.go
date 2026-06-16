package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"health-monitor/internal/checker"
	"health-monitor/internal/config"
)

func main() {
	cfg := config.Load()

	urlChecker := checker.NewChecker()

	result := urlChecker.CheckURL(context.Background(), cfg.URLToCheck)

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
