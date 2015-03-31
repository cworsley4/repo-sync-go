package main

import (
  "fmt"
	"encoding/json"
	"io/ioutil"
  "net/http"
	"os/exec"
)

type Repo struct {
	RepoID       string
	RepoName     string
	RepoDesc     string
	SourceURL    string
	LocalPath    string
	Active       string
	ServerRoleID string
	Required     string
	IsRVStandard string
	StashProject string
}

var ending = `
     _____                     ____                  ____
    / ___/__  ______  _____   / __ \___  _______  __/ / /______
    \__ \/ / / / __ \/ ___/  / /_/ / _ \/ ___/ / / / / __/ ___/
   ___/ / /_/ / / / / /__   / _, _/  __(__  ) /_/ / / /_(__  )
  /____/\__, /_/ /_/\___/  /_/ |_|\___/____/\____/_/\__/____/
       /____/
`

func main() {
	var repos []Repo
	var done = make(chan int)

  resp, _ := http.Get("http://intranet.redventures.net/admin/dev/repo_config/repo_ajax.php?action=get-user-repos&username=cworsley")
  body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &repos)

	if err != nil {
		panic(err)
	}

	for key, _ := range repos {
		fmt.Print("\n")
		repo := repos[key]

		go func(repo Repo, done chan int) {
      fmt.Println("Cloning... %s", repo.RepoName)
			// cmd := exec.Command("git", "clone", repo.SourceURL)
      cmd := exec.Command("sleep", "5")
      cmd.Run()

			done <- 0
		}(repo, done)

		<-done
	}

	fmt.Print(ending)
}
