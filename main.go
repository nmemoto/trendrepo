package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

const githubURL = "https://github.com"

func main() {

	var language, period string
	flag.StringVar(&language, "l", "", "Programming Language")
	flag.StringVar(&period, "p", "today", "Date Range")

	flag.Parse()

	baseURL := githubURL + "/trending/"
	url := baseURL + language + "?since=" + period
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
	doc, err := htmlquery.Parse(r)
	if err != nil {
		fmt.Println(err)
	}

	var repo Repository
	list := htmlquery.Find(doc, "//article")
	for _, n := range list {
		// Href, Author, Name
		hrefPath := htmlquery.SelectAttr(htmlquery.FindOne(n, "/h1/a"), "href")
		repo.Href = githubURL + hrefPath
		repoName := strings.Split(hrefPath, "/")
		repo.Author = repoName[1]
		repo.Name = repoName[2]

		// Description
		descNode, err := htmlquery.QueryAll(n, "/p")
		if err != nil {
			repo.Description = ""
		} else {
			for _, n := range descNode {
				repo.Description = strings.TrimSpace(htmlquery.InnerText(n))
			}
		}

		// Language
		langNode, err := htmlquery.QueryAll(n, "/div[2]/span[1]/span[2]")
		if err != nil {
			repo.Language = ""
		} else {
			for _, n := range langNode {
				repo.Language = htmlquery.InnerText(n)
			}
		}

		// Stars
		starsNode, err := htmlquery.QueryAll(n, "/div[2]/a[1]")
		if err != nil {
			repo.Stars = 0
		} else {
			for _, n := range starsNode {
				repo.Stars, err = strconv.Atoi(strings.TrimSpace(strings.Replace(htmlquery.InnerText(n), ",", "", -1)))
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// Forks
		forksNode, err := htmlquery.QueryAll(n, "/div[2]/a[2]")
		if err != nil {
			repo.Forks = 0
		} else {
			for _, n := range forksNode {
				repo.Forks, err = strconv.Atoi(strings.TrimSpace(strings.Replace(htmlquery.InnerText(n), ",", "", -1)))
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// StarsInPeriod
		starsInPeriodNode, err := htmlquery.QueryAll(n, "/div[2]/span[3]")
		if err != nil {
			repo.StarsInPeriod = 0
		} else {
			for _, n := range starsInPeriodNode {
				r := regexp.MustCompile(`\d+`)
				starsInPeriod := r.FindAllString(htmlquery.InnerText(n), -1)
				repo.StarsInPeriod, err = strconv.Atoi(starsInPeriod[0])
				if err != nil {
					log.Fatal(err)
				}

			}
		}

		repos = append(repos, repo)
	}
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
