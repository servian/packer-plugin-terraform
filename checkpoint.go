package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

// https://checkpoint-api.hashicorp.com/v1/check/terraform
// {"product":"terraform","current_version":"0.12.23","current_release":1583443574,"current_download_url":"https://releases.hashicorp.com/terraform/0.12.23/","current_changelog_url":"https://github.com/hashicorp/terraform/blob/v0.12.23/CHANGELOG.md","project_website":"https://www.terraform.io","alerts":[]}

// CheckResponse is the response for a check request.
type CheckResponse struct {
	Product             string        `json:"product"`
	CurrentVersion      string        `json:"current_version"`
	CurrentReleaseDate  int           `json:"current_release_date"`
	CurrentDownloadURL  string        `json:"current_download_url"`
	CurrentChangelogURL string        `json:"current_changelog_url"`
	ProjectWebsite      string        `json:"project_website"`
	Outdated            bool          `json:"outdated"`
	Alerts              []*CheckAlert `json:"alerts"`
}

// CheckAlert is used in the error message.
type CheckAlert struct {
	ID      int
	Date    int
	Message string
	URL     string
	Level   string
}

// FetchLatestTerraform grabs the latest verions of Terraform from Hashicorp
func FetchLatestTerraform() (string, error) {
	var u url.URL
	v := u.Query()

	v.Set("arch", runtime.GOARCH)
	v.Set("os", runtime.GOOS)
	u.Scheme = "https"
	u.Host = "checkpoint-api.hashicorp.com"
	u.Path = "/v1/check/terraform"
	u.RawQuery = v.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "HashiCorp/go-checkpoint")

	client := cleanhttp.DefaultClient()

	// We use a short timeout since checking for new versions is not critical
	// enough to block on if checkpoint is broken/slow.
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
	log.Println(fmt.Sprintf("Got version response: %s", result.CurrentVersion))
	return result.CurrentVersion, nil
}
