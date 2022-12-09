package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Stats struct {
	generalRequestsTotal           float64
	generalFailedRequestsTotal     float64
	virusScanRequestsRespModsTotal float64
	virusScanRequestsScannedTotal  float64
	virusScanVirusFoundTotal       float64
	virusScanFailureTotal          float64
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
	// fmt.Printf("result: %q\n", stdout.String())
	return stdout.String(), nil
}

func parseResult(result string) (*Stats, error) {
	generalRequests := strings.Split(strings.SplitAfterN(result, "REQUESTS : ", 2)[1], "\n")[0]
	generalRequestsConverted, err := strconv.ParseFloat(generalRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	generalFailedRequests := strings.Split(strings.SplitAfterN(result, "FAILED REQUESTS : ", 2)[1], "\n")[0]
	generalFailedRequestsConverted, err := strconv.ParseFloat(generalFailedRequests, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	virusScanRequestsRespMods := strings.Split(strings.SplitAfterN(result, "Service virus_scan RESPMODS : ", 2)[1], "\n")[0]
	virusScanRequestsRespModsConverted, err := strconv.ParseFloat(virusScanRequestsRespMods, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	virusScanRequestsScanned := strings.Split(strings.SplitAfterN(result, "Requests scanned : ", 2)[1], "\n")[0]
	virusScanRequestsScannedConverted, err := strconv.ParseFloat(virusScanRequestsScanned, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	virusScanVirusFound := strings.Split(strings.SplitAfterN(result, "Viruses found : ", 2)[1], "\n")[0]
	virusScanVirusFoundConverted, err := strconv.ParseFloat(virusScanVirusFound, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	virusScanFailure := strings.Split(strings.SplitAfterN(result, "Scan failures : ", 2)[1], "\n")[0]
	virusScanFailureConverted, err := strconv.ParseFloat(virusScanFailure, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	return &Stats{
		generalRequestsTotal:           generalRequestsConverted,
		generalFailedRequestsTotal:     generalFailedRequestsConverted,
		virusScanRequestsRespModsTotal: virusScanRequestsRespModsConverted,
		virusScanRequestsScannedTotal:  virusScanRequestsScannedConverted,
		virusScanVirusFoundTotal:       virusScanVirusFoundConverted,
		virusScanFailureTotal:          virusScanFailureConverted,
	}, nil
}
