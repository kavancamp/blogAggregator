package feeds

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Title       string `xml:"channel>title"`
	Description string `xml:"channel>description"`
	Items       []Item `xml:"channel>item"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var parsed RSSFeed
	if err := xml.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	parsed.Title = html.UnescapeString(parsed.Title)
	parsed.Description = html.UnescapeString(parsed.Description)
	for i := range parsed.Items {
		parsed.Items[i].Title = html.UnescapeString(parsed.Items[i].Title)
		parsed.Items[i].Description = html.UnescapeString(parsed.Items[i].Description)
	}

	return &parsed, nil
}
