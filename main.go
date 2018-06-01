package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

type GithubIssue struct {
	State string `json:"state"`
}

func main() {
	// read the standard input for lines
	// and do someting for each
	var line string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line = scanner.Text()
		getIssueState(line)
	}
}

func getIssueState(url string) {
	fmt.Println(url)

	token := ""
	r, _ := regexp.Compile("github.com/([[:graph:]]+/[[:graph:]]+/(?:issues)?(?:pull)?/[[:digit:]]+)")
	m := r.FindStringSubmatch(url)

	if len(m) > 0 {
		// fmt.Println("extracted:", m[1])
		githubUrl := fmt.Sprintf(
			"https://api.github.com/repos/%s?access_token=%s",
			m[1], token,
		)
		// fmt.Println("gh url: ", githubUrl)

		// Build the request
		req, err := http.NewRequest("GET", githubUrl, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
		}

		defer resp.Body.Close()

		var githubIssue GithubIssue

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&githubIssue); err != nil {
			log.Println(err)
		}

		fmt.Println("  State:", githubIssue.State)
	} else {
		fmt.Println("  no reference found.")
	}
}
