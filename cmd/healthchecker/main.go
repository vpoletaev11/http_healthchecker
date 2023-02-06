package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"http_healthchecker/internal/check"
	"http_healthchecker/internal/configuration"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	mainCtx, cancel := context.WithCancel(context.Background())
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		<-s
		cancel()
	}()

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

	for _, urlCfg := range config {
		failedChecks := []string{}
		for _, urlCheck := range urlCfg.Checks {
			checker, ok := checkMap[urlCheck.Name]
			if !ok {
				log.Printf("incorrect check %q type for url %q\n", urlCheck, urlCfg.URL)
				failedChecks = append(failedChecks, urlCheck.Name)
				continue
			}

			err := checker.Check(mainCtx, urlCfg.URL, urlCheck.Params)
			if err != nil {
				if !errors.Is(err, check.IncorrectStatusCode) {
					log.Printf("check %q url %q error: %s\n", urlCheck.Name, urlCfg.URL, err)
				}
				failedChecks = append(failedChecks, urlCheck.Name)
			}
		}

		successfulChecksCnt := len(urlCfg.Checks) - len(failedChecks)
		var status string
		if successfulChecksCnt >= urlCfg.MinChecksCnt {
			status = "ok"
		} else {
			status = "fail"
		}

		if len(failedChecks) == 0 {
			fmt.Printf("%s %s\n", urlCfg.URL, status)
			return
		}
		fmt.Printf("%s %s (%s)\n", urlCfg.URL, status, strings.Join(failedChecks, ", "))
	}
}
