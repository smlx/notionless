package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURL       = "https://api.notion.com/v1"
	notionVersion = "2026-03-11"
)

// Client is a Notion API client.
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient creates a new Notion API client.
func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

// PageMarkdownResponse represents the response from the get markdown endpoint.
type PageMarkdownResponse struct {
	Object          string   `json:"object"`
	ID              string   `json:"id"`
	Markdown        string   `json:"markdown"`
	Truncated       bool     `json:"truncated"`
	UnknownBlockIDs []string `json:"unknown_block_ids"`
}

// ReplaceContentRequest represents the request body for replacing page markdown.
type ReplaceContentRequest struct {
	Type           string         `json:"type"`
	ReplaceContent ReplaceContent `json:"replace_content"`
}

type ReplaceContent struct {
	NewStr string `json:"new_str"`
}

// GetPageMarkdown retrieves the markdown content for a given page ID.
func (c *Client) GetPageMarkdown(pageID string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse base URL: %v", err)
	}
	u.Path = path.Join(u.Path, "pages", pageID, "markdown")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("couldn't create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Notion-Version", notionVersion)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't do request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf(`couldn't get page: status=%d body=%s`, resp.StatusCode, string(body))
	}
	var parsedResp PageMarkdownResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsedResp); err != nil {
		return "", fmt.Errorf("couldn't decode response: %v", err)
	}
	return parsedResp.Markdown, nil
}

// ReplacePageMarkdown replaces the markdown content for a given page ID.
func (c *Client) ReplacePageMarkdown(pageID string, markdownContent string) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("couldn't parse base URL: %v", err)
	}
	u.Path = path.Join(u.Path, "pages", pageID, "markdown")
	
	reqBody := ReplaceContentRequest{
		Type: "replace_content",
		ReplaceContent: ReplaceContent{
			NewStr: markdownContent,
		},
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("couldn't marshal request: %v", err)
	}
	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("couldn't create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Notion-Version", notionVersion)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("couldn't do request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(`couldn't replace page: status=%d body=%s`, resp.StatusCode, string(body))
	}
	return nil
}
