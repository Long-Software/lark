package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Long-Software/Bex/apps/cmd/everia/internal/logger"
	"github.com/Long-Software/Bex/packages/file"
	"github.com/Long-Software/Bex/packages/log"
	"github.com/PuerkitoBio/goquery"
)

const (
	FetcherDirName = "fetcher"
)

type Controller struct {
	baseDir string `json:"-"`
}

func NewController() *Controller {
	rootDir, err := file.GetExecDir()
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}
	dir := filepath.Join(rootDir, FetcherDirName)
	err = file.MkdirAll(dir)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}

	return &Controller{baseDir: dir}
}

// TODO: make this function only scan for html/text content and write that to the markdown file
func (c *Controller) FetchAndConvertToMarkdown(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		return "", err
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		return "", err
	}

	var builder strings.Builder

	filename := "Output.md"
	if strings.Contains(contentType, "application/json") {
		builder.WriteString("# JSON Response\n\n")

		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, body, "", "  ")
		if err != nil {
			logger.NewLog(log.FATAL, err.Error())
		}
		builder.WriteString("```json\n")
		builder.WriteString(prettyJSON.String())
		builder.WriteString("\n```\n")
	} else if strings.Contains(contentType, "text/html") {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
		if err != nil {
			logger.NewLog(log.FATAL, err.Error())
			return "", err
		}

		// Extract page title
		title := strings.TrimSpace(doc.Find("title").First().Text())
		builder.WriteString(fmt.Sprintf("# %s\n\n", title))

		// Convert HTML content to markdown
		doc.Find("h1, h2, h3, p, li, code, pre").Each(func(i int, s *goquery.Selection) {
			builder.WriteString(contentFromTag(goquery.NodeName(s), s.Text()))
		})
	} else {
		builder.WriteString("Unsupported content type: " + contentType + "\n")
	}
	filename = c.FilePath(filename)
	err = file.Write(filename, builder.String())
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		return "", err
	}
	return filename, nil
}
func (c *Controller) FilePath(filename string) string {
	return filepath.Join(c.baseDir, filename)
}

func contentFromTag(tag, content string) string {
	switch tag {
	case "h1":
		return "# " + content + "\n\n"
	case "h2":
		return "## " + content + "\n\n"
	case "h3":
		return "### " + content + "\n\n"
	case "p":
		return content + "\n\n"
	case "li":
		return "- " + content + "\n"
	case "pre", "code":
		return "```\n" + content + "\n```\n"
	default:
		return content
	}
}

func SanitizeFilename(name string) string {
	// Step 1: Trim whitespace
	name = strings.TrimSpace(name)

	// Step 2: Replace invalid characters with underscore
	// This covers: \ / : * ? " < > |
	re := regexp.MustCompile(`[<>:"/\\|?*]`)
	name = re.ReplaceAllString(name, "_")

	// Step 3: Fallback to default if empty
	if name == "" {
		name = "untitled"
	}

	// Step 4: Truncate to a safe length (255 characters max)
	if len(name) > 255 {
		name = name[:255]
	}

	// Optionally lower case and trim again
	return strings.TrimSpace(name)
}
