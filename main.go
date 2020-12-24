package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://dle.rae.es/")
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

	fmt.Printf(string(res.Find(re.Find(body))))
}
