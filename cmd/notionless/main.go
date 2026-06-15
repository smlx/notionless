// Package main implements the command-line interface of the tool.
package main

import (
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
	kctx.FatalIfErrorf(kctx.Run(&cli))
}
