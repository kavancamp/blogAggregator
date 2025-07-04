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
type RSSFeed struct {
	Channel struct {
	Title       string `xml:"title"`
	Link        string  `xml:"link"`
	Description string `xml:"description"`
	Item       []RSSItem `xml:"item"`
	}`xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := httpClient.Do(req)
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

	parsed.Channel.Title = html.UnescapeString(parsed.Channel.Title)
	parsed.Channel.Description = html.UnescapeString(parsed.Channel.Description)
	for i, item := range parsed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		parsed.Channel.Item[i] = item
	}

	return &parsed, nil
}
