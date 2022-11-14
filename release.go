package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

// https://api.releases.hashicorp.com/v1/releases/terraform/latest

// CheckResponse is the response for a check request.
type CheckResponse struct {
	Builds                     []*Build `json:"builds"`
	LicenseClass               string   `json:"license_class"`
	Name                       string   `json:"name"`
	TimestampCreated           string   `json:"timestamp_created"`
	TimestampUpdated           string   `json:"timestamp_updated"`
	UrlChangelog               string   `json:"url_changelog"`
	UrlDockerRegistryDockerhub string   `json:"url_docker_registry_dockerhub"`
	UrlDockerRegistryEcr       string   `json:"url_docker_registry_ecr"`
	UrlLicense                 string   `json:"url_license"`
	UrlProjectWebsite          string   `json:"url_project_website"`
	UrlShasums                 string   `json:"url_shasums"`
	UrlSourceRepository        string   `json:"url_source_repository"`
	Version                    string   `json:"version"`
}

// Build is the different OS Arch builds.
type Build struct {
	Arch string
	OS   string
	Url  string
}

// FetchLatestTerraform grabs the latest verions of Terraform from Hashicorp
func FetchLatestTerraform() (string, error) {
	var u url.URL
	v := u.Query()

	u.Scheme = "https"
	u.Host = "api.releases.hashicorp.com"
	u.Path = "/v1/releases/terraform/latest"
	u.RawQuery = v.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Servian/Packer-Plugin-Terraform")

	client := cleanhttp.DefaultClient()

	// We use a short timeout since checking for new versions is not critical
	// enough to block on if the release api is broken/slow.
	client.Timeout = time.Duration(3000) * time.Millisecond

	log.Println(fmt.Sprintf("About to fetch from URL: %s", u.String()))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Unknown status: %d", resp.StatusCode)
	}

	var r io.Reader = resp.Body

	var result CheckResponse
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return "", err
	}
	log.Println(fmt.Sprintf("Got version response: %s", result.Version))
	return result.Version, nil
}
