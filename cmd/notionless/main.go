// Package main implements the command-line interface of the tool.
package main

import (
	"errors"
	"os"

	"github.com/alecthomas/kong"
)

// CLI represents the command-line interface.
type CLI struct {
	Token   string     `kong:"env='NOTION_TOKEN',help='Notion personal access token',required"`
	Version VersionCmd `kong:"cmd,help='Print version information'"`
	Diff    DiffCmd    `kong:"cmd,help='Diff local markdown file against Notion page'"`
	Update  UpdateCmd  `kong:"cmd,help='Update Notion page with local markdown file content'"`
	Fetch   FetchCmd   `kong:"cmd,help='Fetch Notion page markdown and write to a file'"`
}

func main() {
	cli := CLI{}
	kctx := kong.Parse(&cli,
		kong.UsageOnError(),
	)
	err := kctx.Run(&cli)
	if errors.Is(err, ErrDiffFound) {
		os.Exit(2)
	}
	kctx.FatalIfErrorf(err)
}
