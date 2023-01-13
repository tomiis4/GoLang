package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Repos struct {
    Items []Repo `json:"items"`
}

type Repo struct {
    Name string `json:"name"`
    Description string `json:"description"`
    URL string `json:"html_url"`
}

type File struct {
    Name string `json:"name"`
    ContentUrl string `json:"download_url"`
}

func randInt(max int) int {
	 randTime := rand.NewSource(time.Now().UnixNano())
	 newRand := rand.New(randTime)
	 randomNum := newRand.Intn(max)
    	
	 return randomNum
}

func getRepoUrl(lang string) string {
    link := "https://api.github.com/search/repositories?q=language:"+ lang +"&sort=stars"
    
    // send req.
    resp, err := http.Get(link)
    if err != nil { fmt.Println(err) }
    defer resp.Body.Close()
    
    // get repos from req. and convert it
    var repositories Repos
    json.NewDecoder(resp.Body).Decode(&repositories)
    
    // make sure there is at least 1 repo
    if len(repositories.Items) == 0 { return ""}
    
    // get repo
    randomRepo := repositories.Items[randInt(len(repositories.Items))]
    repoUrl := randomRepo.URL
    
    // log it
    fmt.Println("name:", randomRepo.Name)
    fmt.Println("url:", repoUrl)
    
    return repoUrl
}

func getFile(url string) {
    nameUrl := strings.Replace(url, "https://github.com/", "", 1)
    repoUrl := "https://api.github.com/repos/" + nameUrl + "/contents"
    
    // send req.
    resp, err := http.Get(repoUrl)
    if err != nil { fmt.Println(err) }
    defer resp.Body.Close()
    
    // get repos from req. and convert it
    var files []File
    json.NewDecoder(resp.Body).Decode(&files)
    
    // make sure there is at least 1 repo
    if len(files) == 0 { fmt.Println("No items were found") }
    
    fmt.Println(files)
}

func main() { 
    repoUrl := getRepoUrl("typescript")
    getFile(repoUrl)
}
