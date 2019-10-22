package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const githubURL = "https://github.com"

func main() {
	url := githubURL + "/trending/go?since=today"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var r io.Reader = resp.Body
	repos := ParseRepos(r)

	jsonBytes, err := json.Marshal(repos)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	var buf bytes.Buffer
	json.Indent(&buf, jsonBytes, "", "    ")
	fmt.Println(buf.String())

}

func ParseRepos(r io.Reader) []Repository {
	repos := []Repository{}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}

	var repo Repository
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			h1 := n.FirstChild.NextSibling.NextSibling.NextSibling
			aName := h1.FirstChild.NextSibling
			for _, a := range aName.Attr {
				if a.Key == "href" {
					repo.Href = githubURL + a.Val
					repoName := strings.Split(a.Val, "/")
					repo.Author = repoName[1]
					repo.Name = repoName[2]
				}
			}
			pDesc := h1.NextSibling.NextSibling
			repo.Description = strings.TrimSpace(pDesc.FirstChild.Data)
			spanLang := pDesc.NextSibling.NextSibling.FirstChild.NextSibling
			repo.Language = spanLang.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
			aStars := spanLang.NextSibling.NextSibling
			repo.Stars, err = strconv.Atoi(strings.TrimSpace(strings.Replace(aStars.FirstChild.NextSibling.NextSibling.Data, ",", "", -1)))
			if err != nil {
				log.Fatal(err)
			}
			aForks := aStars.NextSibling.NextSibling
			repo.Forks, err = strconv.Atoi(strings.TrimSpace(strings.Replace(aForks.FirstChild.NextSibling.NextSibling.Data, ",", "", -1)))
			if err != nil {
				log.Fatal(err)
			}
			spanStarsInPeriod := aForks.NextSibling.NextSibling.NextSibling.NextSibling
			starsInPeriodStr := spanStarsInPeriod.FirstChild.NextSibling.NextSibling.Data
			r := regexp.MustCompile(`\d+`)
			starsInPeriod := r.FindAllString(starsInPeriodStr, -1)
			repo.StarsInPeriod, err = strconv.Atoi(starsInPeriod[0])
			if err != nil {
				log.Fatal(err)
			}
			repos = append(repos, repo)

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return repos
}

type Repository struct {
	Author        string `json:"author"`
	Name          string `json:"name"`
	Href          string `json:"href"`
	Description   string `json:"description"`
	Language      string `json:"language"`
	Stars         int    `json:"stars"`
	Forks         int    `json:"forks"`
	StarsInPeriod int    `json:"starsInPeriod"`
}
