package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// GetBlogPosts scrapes my blog at https://tansawit.me for the list of published blog posts
// It then vists each posts and retrieves the relevant information and returns it in the form of
// a slice of Post structs
func GetBlogPosts() []Post {
	var allPosts []Post
	c := colly.NewCollector(
		// Allow colly to only scrape my domain
		colly.AllowedDomains("tansawit.me"),
	)

	// On every a HTML element which has href attribute call callback
	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {

		// Get the URL of the post
		postURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(postURL, "tansawit.me/posts/") != -1 {
			res, err := http.Get(postURL)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}
			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Scrape info from article content
			doc.Find(".wrapper main .container article").Each(func(i int, s *goquery.Selection) {
				var (
					categories []string
					tags       []string
				)

				// Post Title
				postTitle := s.Find("h1").Text()

				// Posting Date
				postDate := s.Find(".post-meta .post-meta-other time").Text()

				// Category List
				s.Find(".post-meta .post-meta-main span a").Each(func(i int, s *goquery.Selection) {
					categories = append(categories, s.Text())
				})

				// Tag List
				s.Find(".post-footer .post-info-more section span[class=tag] a").Each(func(i int, s *goquery.Selection) {
					tags = append(tags, strings.TrimSpace(s.Text()))
				})

				// Post Body
				var postContent string
				s.Find(".post-content").Each(func(j int, t *goquery.Selection) {
					postContent = postContent + " " + t.Find("p").Text()
				})

				// Add Post to list of all Posts
				allPosts = append(allPosts, Post{Title: postTitle, Date: postDate, Categories: categories, Tags: tags, Content: postContent, URL: postURL})
			})
		}
	})

	//Sort into chronological order
	for i, j := 0, len(allPosts)-1; i < j; i, j = i+1, j-1 {
		allPosts[i], allPosts[j] = allPosts[j], allPosts[i]
	}

	c.Visit("http://tansawit.me/posts/")

	return allPosts
}

// Post stores the relevant information for each blog post
type Post struct {
	Title      string   `json:"title"`
	Date       string   `json:"date"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
	Content    string   `json:"content"`
	URL        string   `json:"url"`
}
