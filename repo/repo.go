package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	maxBodyBytes = 1 << 20 // 1 MB
	httpTimeout  = 10 * time.Second
)

var httpClient = &http.Client{Timeout: httpTimeout}

func getRepoInfo(owner, name string) (description string, updated time.Time, err error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", url.PathEscape(owner), url.PathEscape(name))
	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", time.Time{}, fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var info struct {
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	r := io.LimitReader(resp.Body, maxBodyBytes)
	if err := json.NewDecoder(r).Decode(&info); err != nil {
		return "", time.Time{}, err
	}

	return info.Description, info.UpdatedAt, nil
}

func main() {
	desc, updatedAt, err := getRepoInfo("golang", "go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Description:", desc)
	fmt.Println("Updated at:", updatedAt.Format(time.RFC1123))
}
