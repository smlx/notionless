package main

import (
	"fmt"
	"os"

	"github.com/smlx/notionless/internal/notion"
)

// FetchCmd represents the `fetch` command.
type FetchCmd struct {
	PageID string `kong:"arg,help='Notion page ID to fetch'"`
	File   string `kong:"arg,help='Path to the file to write'"`
}

// Run the fetch command.
func (cmd *FetchCmd) Run(cli *CLI) error {
	client := notion.NewClient(cli.Token)
	remoteContent, err := client.GetPageMarkdown(cmd.PageID)
	if err != nil {
		return fmt.Errorf("couldn't get page markdown: %v", err)
	}
	content := fmt.Sprintf("---\npage-id: %s\n---\n%s", cmd.PageID, remoteContent)
	if err := os.WriteFile(cmd.File, []byte(content), 0644); err != nil {
		return fmt.Errorf("couldn't write file: %v", err)
	}
	return nil
}
