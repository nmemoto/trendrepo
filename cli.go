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

const (
	ExitCodeOK int = 0

	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeInvalidURL
	ExitCodeParseHtmlError
	ExitCodeJsonMarshalError
)

const githubURL = "https://github.com"

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {

	var language, period string

	flags := flag.NewFlagSet("trendrepo", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&language, "l", "", "Programming Language")
	flags.StringVar(&language, "lang", "", "Programming Language")
	flags.StringVar(&period, "p", "today", "Date Range")
	flags.StringVar(&period, "period", "today", "Date Range")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	baseURL := githubURL + "/trending/"
	url := baseURL + language + "?since=" + period
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(cli.errStream, "HTTP Request/Response error: %v", err)
		return ExitCodeInvalidURL
	}
	defer resp.Body.Close()
	var r io.Reader = resp.Body
	repos, err := ParseRepos(r)
	if err != nil {
		fmt.Fprintf(cli.errStream, "HTML Parse error: %v", err)
		return ExitCodeParseHtmlError
	}

	jsonBytes, err := json.Marshal(repos)
	if err != nil {
		fmt.Fprintf(cli.errStream, "JSON Marshal error: %v", err)
		return ExitCodeJsonMarshalError
	}

	var buf bytes.Buffer
	json.Indent(&buf, jsonBytes, "", "    ")
	fmt.Fprintln(cli.outStream, buf.String())
	return ExitCodeOK
}

func ParseRepos(r io.Reader) ([]Repository, error) {
	repos := []Repository{}
	doc, err := htmlquery.Parse(r)
	if err != nil {
		return nil, err
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
	return repos, nil
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
