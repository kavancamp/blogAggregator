package utils

import (
	"time"
	"encoding/xml"
	"context"
	"fmt"
	"io"
	"net/http"
)
type RSSFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
		//fetch feed from given url 
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil{
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	
	//read and parse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var parsed RSSFeed
	if err := xml.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("error unmarshaling response %w", feedURL)
	}

	if len(parsed.Channel.Items) == 0 {
		return nil, fmt.Errorf("retreiving data at %s", feedURL)
	}

	//return filled-out rssfeed struct
	return &parsed, nil 
}