//usr/bin/env go run $0 $@; exit

// # <xbar.title>My Azure Devops Pull Requests</xbar.title>
// # <xbar.version>v1.0</xbar.version>
// # <xbar.author>Evan Anger</xbar.author>
// # <xbar.author.github>devandanger</xbar.author.github>
// # <xbar.desc>Access Pull Requests of Interests</xbar.desc>
// # <xbar.dependencies>swift</xbar.dependencies>

package main

import (
	"fmt"
	"encoding/json"
	"os"
	"time"
	"net/http"
	"net/url"
	"io/ioutil"
)

type Repository struct {
	Id string
	Name string
	Url string
}

type PullRequest struct {
	Repository Repository
	Url string
	Title string
	Status string
	MergeStatus string
	IsDraft bool
}

func main() {
	printOutput()
}

func printOutput() {
	fmt.Println("")
	org, orgValid := checkEnv("AZDO_ORG")
	if !orgValid {
		fmt.Println("Unable to load value from", "AZDO_ORG")
		return
	}
	project, projectValid := checkEnv("AZDO_PROJ")
	if !projectValid {
		fmt.Println("Unable to load value from", "AZDO_PROJ")
		return
	}
	user, userValid := checkEnv("AZDO_USER")
	if !userValid {
		fmt.Println("Unable to load value from", "AZDO_USER")
		return
	}
	PAT, PATValid := checkEnv("AZDO_PAT")
	if !PATValid {
		fmt.Println("Unable to load value from", "AZDO_PAT")
		return
	}

	pullRequests := getPR(org, project, user, PAT)

	fmt.Println("")
	fmt.Println("My Pull Requests")
	fmt.Println("---")
	for _, pr := range pullRequests {
		fmt.Println(pr.Repository.Name, "/", pr.Title)
	}
	fmt.Println("Assigned Pull Requests")
	fmt.Println("---")
	val := DoneAsync()
	fmt.Println("Done is running ...")
	fmt.Println(<-val)
}

func checkEnv(key string) (string, bool) {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value, true
	} else {
		return value, false
	}
}

func getPR(org string, project string, user string, PAT string) []PullRequest {
	urlString := "https://dev.azure.com/" + org + "/" + project + "/_apis/git/pullrequests?api-version=6.1-preview.1"
	reqUrl, _ := url.Parse(urlString)
	request := &http.Request {
		Method: "GET",
		URL: reqUrl,
		Header: map[string][]string {
			
		},
	}

	request.SetBasicAuth(user, PAT)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("Error making requests")
		return []PullRequest{}
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return getPRData(data)
}

func getPRData(data []byte) [] PullRequest {
	type PullRequestData struct {
		Value []PullRequest
		Count json.Number
	}

	// var jsonBlob = []byte(`
	// {
	// 	"value": [],
	// 	"count": 3
	// }`)

	var result PullRequestData
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		return []PullRequest{}
	}
	fmt.Println("Result:", result)
	return result.Value
}



func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		r <- 1
		fmt.Println("Done ...")
	}()
	return r
}
