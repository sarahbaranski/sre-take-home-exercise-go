package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"

	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Name    string            `yaml:"name"`
	URL     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

type DomainStats struct {
	Success int
	Total   int
}

var stats = make(map[string]*DomainStats)

func checkHealth(endpoint Endpoint) {
	var client = &http.Client{
		Timeout: 500 * time.Millisecond,
	}

	bodyBytes, err := json.Marshal(endpoint)
	if err != nil {
		return
	}
	reqBody := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, reqBody)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	for key, value := range endpoint.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	domain := extractDomain(endpoint.URL)

	stats[domain].Total++
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		stats[domain].Success++
	}

}

func extractDomain(urL string) string {
	u, err := url.Parse(urL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			fmt.Println("Error splitting host and port:", err)
		}
	}

	urlSplit := strings.Split(host, "//")
	domain := strings.Split(urlSplit[len(urlSplit)-1], "/")[0]

	return domain
}

func monitorEndpoints(endpoints []Endpoint) {
	for _, endpoint := range endpoints {
		domain := extractDomain(endpoint.URL)
		if stats[domain] == nil {
			stats[domain] = &DomainStats{}
		}
	}

	for {
		for _, endpoint := range endpoints {
			checkHealth(endpoint)
		}
		logResults()
		time.Sleep(15 * time.Second)
	}
}

func logResults() {
	for domain, stat := range stats {
		percentage := math.Round(100 * float64(stat.Success) / float64(stat.Total))
		fmt.Printf("%s has %f%% availability\n", domain, percentage)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <config_file>")
	}

	filePath := os.Args[1]
	fileExtension := filepath.Ext(filePath)
	if fileExtension != ".yaml" {
		log.Fatal("Error config is not yaml file")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	var endpoints []Endpoint
	if err := yaml.Unmarshal(data, &endpoints); err != nil {
		log.Fatal("Error parsing YAML:", err)
	}

	monitorEndpoints(endpoints)
}
