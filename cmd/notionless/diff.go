package main

import (
	"errors"
	"fmt"

	"github.com/smlx/notionless/internal/markdown"
	"github.com/smlx/notionless/internal/notion"
	"znkr.io/diff/textdiff"
)

// DiffCmd represents the `diff` command.
type DiffCmd struct {
	File string `kong:"arg,help='Path to the markdown file'"`
}

var ErrDiffFound = errors.New("differences found")

// Run the diff command.
func (cmd *DiffCmd) Run(cli *CLI) error {
	doc, err := markdown.ParseFile(cmd.File)
	if err != nil {
		return fmt.Errorf("couldn't parse file: %v", err)
	}
	client := notion.NewClient(cli.Token)
	remoteContent, err := client.GetPageMarkdown(doc.PageID)
	if err != nil {
		return fmt.Errorf("couldn't get page markdown: %v", err)
	}
	diffString := textdiff.Unified(
		remoteContent,
		doc.Content,
		textdiff.TerminalColors())
	if diffString != "" {
		fmt.Print(diffString)
		return ErrDiffFound
	}
	return nil
}
