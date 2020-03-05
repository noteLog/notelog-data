package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-github/v29/github"
)

// CreateUpdateBlogTable retrieves the list of my blog posts from https://tansawit.me/posts
// And feed each post data into the main MySQL database

func esGetClient(esURL, esPWD string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},

		Username: "elastic",
		Password: esPWD,
	}
	es, err := elasticsearch.NewClient(cfg)
	return es, err
}

func esIndexGitHub(esURL, esPWD, ghToken string) {
	es, err := esGetClient(esURL, esPWD)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	}

	reposUser, err := githubGetUserRepos(ghToken)
	if err != nil {
		log.Printf("Error fetching user repos: %v", err)
	}

	reposStarred, err := githubGetStarredRepos(ghToken)
	if err != nil {
		log.Printf("Error fetching starred repos: %v", err)
	}

	indexGitHub(es, "repo-personal", reposUser)
	indexGitHub(es, "repo-starred", reposStarred)

	log.Println(strings.Repeat("-", 37))
}

func indexGitHub(es *elasticsearch.Client, index string, repos []*github.Repository) {
	var wg sync.WaitGroup

	log.Printf(`%s Indexing Starred Repositories to Index "%s" %s`, strings.Repeat("-", 10), strings.Repeat("-", 10), index)
	for i, repo := range repos {
		wg.Add(1)

		if i > 0 && i%50 == 0 {
			time.Sleep(500 * time.Millisecond)
		}

		go func(i int, repo *github.Repository) {
			defer wg.Done()

			var b strings.Builder
			b.WriteString(`{"username" : "`)
			b.WriteString(*repo.Owner.Login)
			b.WriteString(`","reponame" : "`)
			b.WriteString(*repo.Name)
			b.WriteString(`","url" : "`)
			b.WriteString(nilableString(repo.HTMLURL))
			b.WriteString(`","description" : "`)
			b.WriteString(nilableString(repo.HTMLURL))
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      index,
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d item='%v'", res.Status(), i+1, b.String())
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, repo)
	}
	wg.Wait()
}

func nilableString(s *string) string {
	if s == nil {
		return ""
	}
	return strings.Replace(*s, `"`, `'`, -1)
}

//func esGetInfo(esURL, esPWD string) (string, string) {

//	log.SetFlags(0)

//	var r map[string]interface{}

//	//Get cluster info
//	es, err := esGetClient(esURL, esPWD)
//	if err != nil {
//		log.Fatalf("Error creating the client: %s", err)
//	}
//	res, err := es.Info()
//	if err != nil {
//		log.Fatalf("Error getting response: %s", err)
//	}
//	defer res.Body.Close()

//	// Check response status
//	if res.IsError() {
//		log.Fatalf("Error: %s", res.String())
//	}
//	// Deserialize the response into a map.
//	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//		log.Fatalf("Error parsing the response body: %s", err)
//	}
//	return elasticsearch.Version, r["version"].(map[string]interface{})["number"].(string)
//}
