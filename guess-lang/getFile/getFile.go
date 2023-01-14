package getFile 

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
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
    Branch string `json:"default_branch"`
}

type Files struct {
    Items []File `json:"tree"`
}

type File struct {
    Name string `json:"name"`
    PathName string `json:"path"`
    Content string `json:"content"`
}

func randInt(max int) int {
	 randTime := rand.NewSource(time.Now().UnixNano())
	 newRand := rand.New(randTime)
	 randomNum := newRand.Intn(max)
    	
	 return randomNum
}

func getRepoUrl(lang string) [2]string {
    link := "https://api.github.com/search/repositories?q=language:"+ lang +"&sort=stars"
    
    // send req.
    resp, err := http.Get(link)
    if err != nil { fmt.Println(err) }
    defer resp.Body.Close()
    
    // get repos from req. and convert it
    var repositories Repos
    json.NewDecoder(resp.Body).Decode(&repositories)
    
    // make sure there is at least 1 repo
    if len(repositories.Items) == 0 { return [2]string{"", ""}}
    
    // get repo
    randomRepo := repositories.Items[randInt(len(repositories.Items))]
    repoUrl := randomRepo.URL
    repoBranch := randomRepo.Branch
    
    
    result := [2]string{repoUrl, repoBranch}
    
    return result
}

func getFile(url string, branch string) string {
    nameUrl := strings.Replace(url, "https://github.com/", "", 1)
    repoUrl := "https://api.github.com/repos/" + nameUrl + "/git/trees/" + branch + "?recursive=1"
    
    // send req.
    resp, err := http.Get(repoUrl)
    if err != nil { fmt.Println(err) }
    defer resp.Body.Close()
    
    // get repos from req. and convert it
    var files Files
    json.NewDecoder(resp.Body).Decode(&files)
    
    // make sure there is at least 1 repo
    if len(files.Items) == 0 { fmt.Println("No items were found") }
    
    // get files which have extension "ts"
    selectedFiles := []string{}
    for _, value := range files.Items {
        if strings.HasSuffix(value.PathName, "ts") {
            selectedFiles = append(selectedFiles, value.PathName)
        }
    }
    
    if len(selectedFiles) <=0 { fmt.Println("There is not any file") }
    
    return selectedFiles[randInt(len(selectedFiles))]
}

func getFileContent(nameUrl string, fileName string) string {
    fileUrl := strings.Replace(nameUrl, "https://github.com/", "", 1)
    repoUrl := "https://api.github.com/repos/" + fileUrl + "/contents/" + fileName
    
    // send req.
    resp, err := http.Get(repoUrl)
    if err != nil { fmt.Println(err) }
    defer resp.Body.Close()
    
    // get repos from req. and convert it
    var randomFile File
    json.NewDecoder(resp.Body).Decode(&randomFile)
    
    return randomFile.Content
}

func parseFile(content string) []string {
    // convert base64 to string
    data, _ := base64.URLEncoding.DecodeString(content)
    
    // split by lines
    dataStr := string(data)
    regexPattern := `\r?\n` // end of the line
    regexObj := regexp.MustCompile(regexPattern)
    
    // split string
    fileResult := regexObj.Split(dataStr, -1)
    
    // first 13 lines
    guessContent := []string{}
    
    if len(fileResult) > 13 {
        for i:=0; i < 13; i++ {
            guessContent = append(guessContent, fileResult[i])
        }
    } else {
        guessContent = fileResult
    }
    
    return guessContent 
}

func GetContent(lang string) []string { 
    repo := getRepoUrl(lang)
    
    repoUrl := repo[0]
    repoBranch := repo[1]
    fileName := getFile(repoUrl, repoBranch)
    
    contentBase64 := getFileContent(repoUrl, fileName)
    result := parseFile(contentBase64)
    
    // if len(result) < 13 {
    //     return GetContent(lang)
    // }
    
    return result
}
