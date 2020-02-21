package main

import (
	"testing"
)

func TestGetBlogPosts(t *testing.T) {
	posts := GetBlogPosts()
	for _, post := range posts {
		if len(post.Title) == 0 || len(post.URL) == 0 || len(post.Date) == 0 {
			t.Errorf("Malformed post struct")
		}
	}
}
