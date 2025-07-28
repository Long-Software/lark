package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Long-Software/lark/pkg/log"
	"github.com/PuerkitoBio/goquery"
)


type Result struct {
	Rank  int
	URL   string
	Title string
	Desc  string
}

type Agent struct {
	names []string
}

func NewAgent() *Agent {
	return &Agent{
		names: []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
			"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
		},
	}
}
func (a *Agent) Random() string {
	idx := rand.Int() % len(a.names)
	return a.names[idx]
}

type Scraper struct {
	terms        string
	TLD          string
	pages        int
	count        int
	proxy        string
	timeout      int
	languageCode string
}

func (s *Scraper) Run() ([]Result, error) {
	results := []Result{}
	counter := 0
	urls, err := s.build()
	if err != nil {
		return nil, err
	}

	lg.NewLog(log.INFO, fmt.Sprintf("Starting scrape with %d URLs", len(urls)))

	for i, url := range urls {
		lg.NewLog(log.INFO, fmt.Sprintf("Requesting URL %d/%d: %s", i+1, len(urls), url))
		res, err := s.request(url)
		if err != nil {
			lg.NewLog(log.ERROR, "Request failed: "+err.Error())
			continue
		}

		data, err := s.parse(res, counter)
		if err != nil {
			lg.NewLog(log.ERROR, "Parse failed: "+err.Error())
			continue
		}

		counter += len(data)
		for _, item := range data {
			results = append(results, item)
		}

		if i < len(urls)-1 { // Don't sleep after the last request
			time.Sleep(time.Duration(s.timeout) * time.Second)
		}
	}

	lg.NewLog(log.INFO, fmt.Sprintf("Total results found: %d", len(results)))
	return results, nil
}

func (s *Scraper) request(searchURL string) (*http.Response, error) {
	a := NewAgent()

	// Handle proxy if provided
	var client *http.Client
	if s.proxy != "" {
		proxyUrl, err := url.Parse(s.proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
			Timeout: 30 * time.Second,
		}
	} else {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Set comprehensive headers to appear more like a real browser
	req.Header.Set("User-Agent", a.Random())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate") // Removed 'br' to avoid Brotli compression issues
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")
	// Removing Sec-Fetch headers as they might trigger bot detection
	// req.Header.Set("Sec-Fetch-Dest", "document")
	// req.Header.Set("Sec-Fetch-Mode", "navigate")
	// req.Header.Set("Sec-Fetch-Site", "none")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		res.Body.Close()
		return nil, fmt.Errorf("scraper received a non-200 status code (%d) suggesting a ban", res.StatusCode)
	}

	return res, nil
}

func (s *Scraper) parse(res *http.Response, rank int) ([]Result, error) {
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	results := []Result{}

	// Log the HTML title for debugging
	title := doc.Find("title").Text()

	// Check if Google is blocking us
	if strings.Contains(strings.ToLower(title), "blocked") ||
		strings.Contains(strings.ToLower(title), "unusual traffic") ||
		strings.Contains(strings.ToLower(title), "captcha") ||
		strings.Contains(strings.ToLower(title), "verify") {
		return nil, fmt.Errorf("Google has blocked this request: %s", title)
	}

	// Log some key page elements for debugging
	lg.NewLog(log.DEBUG, fmt.Sprintf("Page has search form: %t", doc.Find("form[action*='search']").Length() > 0))
	lg.NewLog(log.DEBUG, fmt.Sprintf("Page has Google branding: %t", doc.Find("[alt*='Google'], #logo").Length() > 0))
	lg.NewLog(log.DEBUG, fmt.Sprintf("Page total divs: %d", doc.Find("div").Length()))

	// Try multiple modern selectors as Google changes their structure frequently
	selectors := []string{
		"div#rso",             // Classic selector
		".tF2Cxc",           // Updated 2023+ selector
		"div[data-ved]",     // Modern selector with data attribute
		".g",                // Alternative
		"div.dURPMd",        // Another modern selector
		"div[jscontroller]", // JS controller divs
		".srg > div",        // Search results group children
		"#search div",       // Basic search content
	}

	var sel *goquery.Selection
	for _, selector := range selectors {
		sel = doc.Find(selector)
		if sel.Length() > 0 {
			lg.NewLog(log.INFO, fmt.Sprintf("Using selector '%s', found %d elements", selector, sel.Length()))
			break
		} else {
			lg.NewLog(log.DEBUG, fmt.Sprintf("Selector '%s' found 0 elements", selector))
		}
	}

	if sel.Length() == 0 {
		// Enhanced debugging - log more specific elements
		lg.NewLog(log.DEBUG, fmt.Sprintf("Found <a> tags: %d", doc.Find("a").Length()))
		lg.NewLog(log.DEBUG, fmt.Sprintf("Found <h3> tags: %d", doc.Find("h3").Length()))
		lg.NewLog(log.DEBUG, fmt.Sprintf("Found search-related divs: %d", doc.Find("div[id*='search'], div[class*='search']").Length()))

		bodyText := doc.Find("body").Text()
		if len(bodyText) > 1000 {
			bodyText = bodyText[:1000] + "..."
		}
		htmlContent, _ := doc.Html()
		if len(htmlContent) > 2000 {
			htmlContent = htmlContent[:2000] + "..."
		}

		return results, nil
	}

	rank++
	sel.Each(func(i int, item *goquery.Selection) {
		// Try multiple link selectors
		var link string
		var title string
		var desc string

		// Try different link selectors for modern Google
		linkSelectors := []string{
			"h3 a",                  // Most common
			"a[href^='http']",       // Direct HTTP links
			"a[href^='https']",      // Direct HTTPS links
			".yuRUbf a",             // Modern container
			"div[role='heading'] a", // Role-based
			".LC20lb a",             // Title link
			"a",                     // Any link as fallback
		}

		for _, linkSel := range linkSelectors {
			if linkEl := item.Find(linkSel).First(); linkEl.Length() > 0 {
				if href, exists := linkEl.Attr("href"); exists && href != "" {
					link = href
					break
				}
			}
		}

		// Try different title selectors
		titleSelectors := []string{
			"h3",                  // Standard heading
			"h3.LC20lb",           // Modern title class
			".DKV0Md",             // Alternative title
			"[role='heading'] h3", // Role-based heading
			".yuRUbf h3",          // Container-based
			"a h3",                // Link containing heading
		}

		for _, titleSel := range titleSelectors {
			if titleEl := item.Find(titleSel).First(); titleEl.Length() > 0 {
				title = strings.TrimSpace(titleEl.Text())
				if title != "" {
					break
				}
			}
		}

		// Try different description selectors
		descSelectors := []string{
			".VwiC3b",                           // Modern description
			".s",                                // Classic
			".st",                               // Old classic
			"[data-sncf='1']",                   // Data attribute
			".aCOpRe",                           // Alternative
			"span[style*='-webkit-line-clamp']", // CSS-based description
		}

		for _, descSel := range descSelectors {
			if descEl := item.Find(descSel).First(); descEl.Length() > 0 {
				desc = strings.TrimSpace(descEl.Text())
				if desc != "" {
					break
				}
			}
		}

		// Clean up the link
		link = strings.TrimSpace(link)

		// Handle Google's URL wrapping
		if strings.HasPrefix(link, "/url?q=") {
			// Extract the actual URL from Google's wrapper
			if parts := strings.Split(link, "&"); len(parts) > 0 {
				link = strings.TrimPrefix(parts[0], "/url?q=")
				// URL decode the link
				if decoded, err := url.QueryUnescape(link); err == nil {
					link = decoded
				}
			}
		}

		// Validate and add result
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") &&
			(strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://")) {
			result := Result{
				Rank:  rank,
				URL:   link,
				Title: title,
				Desc:  desc,
			}
			results = append(results, result)
			rank++
		} else if link != "" {
			lg.NewLog(log.DEBUG, fmt.Sprintf("Skipped invalid link: '%s'", link))
		}
	})

	return results, nil
}

func (s *Scraper) build() ([]string, error) {
	urls := []string{}
	s.terms = strings.Trim(s.terms, " ")
	s.terms = strings.Replace(s.terms, " ", "+", -1)

	if base, found := domains[s.TLD]; found {
		for i := 0; i < s.pages; i++ {
			start := i * s.count
			url := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", base, s.terms, s.count, s.languageCode, start)
			urls = append(urls, url)
		}
	} else {
		return nil, fmt.Errorf("country (%s) is not supported", s.TLD)
	}
	return urls, nil
}

func main() {
	s := Scraper{
		terms:        "akhil sharma",
		TLD:          "com",
		languageCode: "en",
		pages:        1,
		count:        10, // Reduced for testing
		timeout:      2,  // Reduced timeout for faster testing
	}

	lg.NewLog(log.INFO, "Starting Google scraper...")
	res, err := s.Run()
	if err != nil {
		lg.NewLog(log.ERROR, err.Error())
		return
	}

	lg.NewLog(log.INFO, fmt.Sprintf("Total results retrieved: %d", len(res)))
	for i, r := range res {
		lg.NewLog(log.INFO, fmt.Sprintf("Result %d: %s - %s", i+1, r.Title, r.URL))
	}
}

var lg = log.Logger{
	IsProduction: false, // Set to false for detailed debugging
	HasTimestamp: true,
	HasFilepath:  true,
	HasMethod:    true,
}
