package main

import (
	"context"
	"flag"
	"fmt"
	"http_healthchecker/internal/check"
	"http_healthchecker/internal/configuration"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	mainCtx := context.Background()

	configPath := flag.String("config_path", "", "path to config file")
	reqTimeoutSec := flag.Int("req_timeout_sec", 10, "http request timeout in seconds")
	flag.Parse()
	if *configPath == "" {
		panic("config_path flag is undefined")
	}

	config, err := configuration.GetConfigFromFile(*configPath)
	if err != nil {
		panic(fmt.Errorf("get configuration: %w", err))
	}

	client := http.Client{
		Timeout: time.Duration(*reqTimeoutSec) * time.Second,
	}

	checkMap := check.NewCheckerMap(&client)

	for _, cfg := range config {
		failedChecks := []string{}
		for _, check := range cfg.Checks {
			checker, ok := checkMap[check.Name]
			if !ok {
				log.Printf("incorrect check %q type for url %q", check, cfg.URL)
			}

			err := checker.Check(mainCtx, cfg.URL, check.Params)
			if err != nil {
				failedChecks = append(failedChecks, check.Name)
			}
		}

		successfulChecksCnt := len(cfg.Checks) - len(failedChecks)
		var status string
		if successfulChecksCnt >= cfg.MinChecksCnt {
			status = "ok"
		} else {
			status = "fail"
		}

		if len(failedChecks) == 0 {
			fmt.Printf("%s %s", cfg.URL, status)
			return
		}
		fmt.Printf("%s %s (%s)", cfg.URL, status, strings.Join(failedChecks, ", "))
	}
}
