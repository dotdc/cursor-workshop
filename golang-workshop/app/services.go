package app

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/russross/blackfriday/v2"
)

var (
	workshopOnce sync.Once
	workshopHTML template.HTML
)

// GetBooks returns the seeded catalogue.
func GetBooks() []Book {
	return CloneBooks()
}

// GetHealthStatus exposes the current health indicator.
func GetHealthStatus() string {
	return "OK"
}

// FormatRarity renders the rating with a trailing star symbol.
func FormatRarity(rating float64) string {
	return fmt.Sprintf("%.1f★", rating)
}

// LoadWorkshopGoalsHTML parses the Workshop Quests section from README.md and converts it to HTML.
func LoadWorkshopGoalsHTML(basePath string) template.HTML {
	workshopOnce.Do(func() {
		readmePath := filepath.Join(basePath, "README.md")
		raw, err := os.ReadFile(readmePath)
		if err != nil {
			workshopHTML = template.HTML(`<p>Workshop goals are unavailable.</p>`)
			return
		}

		section := extractWorkshopSection(string(raw))
		if section == "" {
			workshopHTML = template.HTML(`<p>Workshop goals are unavailable.</p>`)
			return
		}

		rendered := blackfriday.Run([]byte(section))
		rendered = decorateHeadings(rendered)
		workshopHTML = template.HTML(rendered)
	})

	return workshopHTML
}

func extractWorkshopSection(content string) string {
	heading := regexp.MustCompile(`(?m)^##\s+Workshop[^\n]*`)
	loc := heading.FindStringIndex(content)
	if loc == nil {
		return ""
	}

	section := content[loc[1]:]

	if idx := strings.Index(section, "\n## "); idx != -1 {
		section = section[:idx]
	}

	if idx := strings.Index(section, "\n---"); idx != -1 {
		section = section[:idx]
	}

	lines := strings.Split(section, "\n")
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

var headingDecorator = regexp.MustCompile(`(?s)(<h3[^>]*>)(.*?)(</h3>)`)

func decorateHeadings(html []byte) []byte {
	decorated := headingDecorator.ReplaceAllFunc(html, func(match []byte) []byte {
		submatches := headingDecorator.FindSubmatch(match)
		if len(submatches) != 4 {
			return match
		}

		var buf bytes.Buffer
		buf.Write(submatches[1])
		buf.WriteString("✶ ")
		buf.Write(submatches[2])
		buf.WriteString(" ✶")
		buf.Write(submatches[3])
		buf.WriteString("\n<hr style=\"color:red\"><br>")
		return buf.Bytes()
	})
	return decorated
}
