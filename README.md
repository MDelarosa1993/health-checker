# Health Monitor

A simple Go-based health-checking application that checks whether a URL is up or down and prints the result as formatted JSON.

This project is the first step toward building a cloud health monitor. Right now, it runs locally from the command line. Later, this same logic can be reused inside AWS Lambda, connected to CloudWatch, SNS, DynamoDB, and eventually deployed with Terraform.

## Project Purpose

The goal of this project is to understand the core logic behind a cloud health monitoring system before adding AWS services.

This version teaches:

* How to structure a Go project using packages
* How to read configuration from environment variables and command-line flags
* How to make HTTP requests in Go
* How to handle request errors and timeouts
* How to return clean JSON output
* How to separate application concerns into different folders

## Project Structure

```txt
health-monitor/
  go.mod
  cmd/
    healthcheck/
      main.go
  internal/
    checker/
      checker.go
    config/
      config.go
    models/
      health_result.go
```

## Folder Breakdown

### `cmd/healthcheck`

Contains the main entry point for running the application locally.

The `main.go` file:

* Loads configuration
* Creates a new URL checker
* Runs the health check
* Converts the result into formatted JSON
* Prints the result to the terminal

### `internal/config`

Contains configuration logic for the app.

The `config` package checks for a URL in this order:

1. Command-line flag using `-url`
2. Environment variable called `URL_TO_CHECK`
3. Default value: `https://example.com`

Example:

```bash
go run ./cmd/healthcheck -url https://google.com
```

Or with an environment variable:

```bash
URL_TO_CHECK=https://google.com go run ./cmd/healthcheck
```

### `internal/checker`

Contains the core health-checking logic.

The `checker` package:

* Creates an HTTP client with a 5-second timeout
* Sends a GET request to the target URL
* Marks the URL as `UP` if the response status code is below 400
* Marks the URL as `DOWN` if the request fails or returns a status code of 400 or higher

### `internal/models`

Contains shared data structures used across the app.

The `HealthResult` model represents the response from a health check.

```go
type HealthResult struct {
    URL        string `json:"url"`
    Status     string `json:"status"`
    StatusCode int    `json:"statusCode"`
    Error      string `json:"error,omitempty"`
    CheckedAt  string `json:"checkedAt"`
}
```

## How to Run the App

From the root of the project, run:

```bash
go run ./cmd/healthcheck
```

By default, the app checks:

```txt
https://example.com
```

## Check a Custom URL

Use the `-url` flag:

```bash
go run ./cmd/healthcheck -url https://google.com
```

## Use an Environment Variable

You can also set the URL with `URL_TO_CHECK`:

```bash
URL_TO_CHECK=https://github.com go run ./cmd/healthcheck
```

## Example Successful Output

```json
{
  "url": "https://example.com",
  "status": "UP",
  "statusCode": 200,
  "checkedAt": "2026-06-16T12:00:00Z"
}
```

## Example Failed Output

```json
{
  "url": "https://not-a-real-site-12345.com",
  "status": "DOWN",
  "error": "Get \"https://not-a-real-site-12345.com\": dial tcp: lookup not-a-real-site-12345.com: no such host",
  "checkedAt": "2026-06-16T12:00:00Z"
}
```

## Current Features

* Check any URL from the command line
* Use an environment variable as a default URL
* Use a command-line flag to override the URL
* Return structured JSON output
* Mark URLs as `UP` or `DOWN`
* Handle HTTP errors
* Handle network errors
* Add a timeout to prevent long-running requests

## What This Project Teaches

This project introduces the foundation of a production-style health monitor.

The current version focuses on local application logic. Future versions can add cloud services such as:

* AWS Lambda to run the health check without managing servers
* Amazon EventBridge to run the check on a schedule
* Amazon CloudWatch for logs and alerts
* Amazon SNS for email notifications
* Amazon DynamoDB to store health-check history
* Terraform to create the AWS infrastructure

## Future Improvements

Possible next steps:

* Add unit tests for the checker package
* Add support for multiple URLs
* Save health check results to a file
* Add retry logic
* Add AWS Lambda support
* Send alerts when a URL is down
* Store results in DynamoDB
* Deploy infrastructure with Terraform

## Summary

This project is a simple but practical Go application that checks the health of a URL and returns a structured result.

It is designed to start small locally, then grow into a cloud-based monitoring tool using AWS services.
