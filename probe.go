package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	generalRequestsTotal       float64
	generalFailedRequestsTotal float64
	virusScanRequestsTotal     float64
	virusScanVirusFoundTotal   float64
	virusScanFailureTotal      float64
}

func getStats(config *Config) (*Stats, error) {
	result, err := execCmd(config)
	if err != nil {
		return nil, fmt.Errorf("execCmd error: %w", err)
	}
	return parseResult(result)

}

func execCmd(config *Config) (string, error) {
	cmd := exec.Command(
		config.ICAPClientPath,
		"-s", "info?view=text",
		"-i", config.ICAPAddress,
		"-p", config.ICAPPort,
		"-req", "use-any-url",
	)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("cmd.Run error: %w %s %s", err, stderr.String(), stdout.String())
	}
	fmt.Printf("result: %q\n", stdout.String())
	return stdout.String(), nil
}

func parseResult(result string) (*Stats, error) {
	requests := strings.Split(strings.SplitAfterN(result, "REQUESTS : ", 2)[1], "\n")[0]
	requestsConverted, err := strconv.ParseFloat(requests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	failedRequests := strings.Split(strings.SplitAfterN(result, "FAILED REQUESTS : ", 2)[1], "\n")[0]
	failedRequestsConverted, err := strconv.ParseFloat(failedRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	scannedRequests := strings.Split(strings.SplitAfterN(result, "Requests scanned : ", 2)[1], "\n")[0]
	scannedRequestsConverted, err := strconv.ParseFloat(scannedRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	virusFoundRequests := strings.Split(strings.SplitAfterN(result, "Viruses found : ", 2)[1], "\n")[0]
	virusFoundRequestsConverted, err := strconv.ParseFloat(virusFoundRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	scanFailureRequests := strings.Split(strings.SplitAfterN(result, "Scan failures : ", 2)[1], "\n")[0]
	scanFailureRequestsConverted, err := strconv.ParseFloat(scanFailureRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	return &Stats{
		generalRequestsTotal:       requestsConverted,
		generalFailedRequestsTotal: failedRequestsConverted,
		virusScanRequestsTotal:     scannedRequestsConverted,
		virusScanVirusFoundTotal:   virusFoundRequestsConverted,
		virusScanFailureTotal:      scanFailureRequestsConverted,
	}, nil
}
