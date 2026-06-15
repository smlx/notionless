package markdown

import (
	"errors"
	"os"
	"strings"
)

// Document represents a parsed markdown document with its frontmatter page ID and body content.
type Document struct {
	PageID  string
	Content string
}

// ParseFile reads a markdown file and extracts its frontmatter page-id and body content.
func ParseFile(path string) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(data)
	// expect the file to start with frontmatter delimiters
	if !strings.HasPrefix(content, "---\n") {
		return nil, errors.New("missing YAML frontmatter")
	}
	// find the end of the frontmatter
	endIndex := strings.Index(content[4:], "---\n")
	if endIndex == -1 {
		return nil, errors.New("missing end of YAML frontmatter")
	}
	endIndex += 4 // adjust for the offset
	frontmatter := content[4:endIndex]
	body := content[endIndex+4:] // skip "---\n"
	// parse page-id from frontmatter
	var pageID string
	for line := range strings.SplitSeq(frontmatter, "\n") {
		line = strings.TrimSpace(line)
		if after, ok := strings.CutPrefix(line, "page-id:"); ok {
			pageID = strings.TrimSpace(after)
			// remove possible quotes
			pageID = strings.Trim(pageID, `"'`)
			break
		}
	}
	if pageID == "" {
		return nil, errors.New("missing page-id in frontmatter")
	}
	return &Document{
		PageID:  pageID,
		Content: body,
	}, nil
}
