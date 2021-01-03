package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func wordOfTheDay() string {
	url := "https://dle.rae.es"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`<a data-cat='WOTD' data-acc='WOTD'[ A-z-="'%;À-ú/?0-9]+>`)

	res := regexp.MustCompile(`\/[ A-z-="'%;À-ú/0-9]+`)
	route := string(res.Find(re.Find(body)))

	return route
}

func getDefinition(wordRoute string) string {
	url := "https://dle.rae.es"
	resp, err := http.Get(url + wordRoute)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`<article id=[\s\S]+<\/article>`)
	defs := regexp.MustCompile(`<p class="j" id="[ A-z-="'%;À-ú/?0-9]+">[ A-z-="'%;À-ú/?0-9.<>&#|,]+<\/p>`)

	article := re.Find(body)

	t := defs.Find(article)

	doc, err := html.Parse(strings.NewReader(string(t)))
	if err != nil {
		fmt.Println(err)
	}

	def := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "mark" {
			def += n.FirstChild.Data + " "
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return def
}
