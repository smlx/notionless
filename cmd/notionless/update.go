package main

import (
	"fmt"

	"github.com/smlx/notionless/internal/markdown"
	"github.com/smlx/notionless/internal/notion"
)

// UpdateCmd represents the `update` command.
type UpdateCmd struct {
	File string `kong:"arg,help='Path to the markdown file'"`
}

// Run the update command.
func (cmd *UpdateCmd) Run(cli *CLI) error {
	doc, err := markdown.ParseFile(cmd.File)
	if err != nil {
		return fmt.Errorf("couldn't parse file: %v", err)
	}
	client := notion.NewClient(cli.Token)
	if err := client.ReplacePageMarkdown(doc.PageID, doc.Content); err != nil {
		return fmt.Errorf("couldn't replace remote markdown: %v", err)
	}
	return nil
}
