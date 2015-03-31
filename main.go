package main

import (
  "fmt"
  "sync"
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

  resp, _ := http.Get("")
  body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &repos)

	if err != nil {
		panic(err)
	}

  var wg sync.WaitGroup
  var tasks = make(chan Repo)

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(wg *sync.WaitGroup) {
      defer wg.Done()
      for task := range tasks {
        fmt.Println("Cloning... %s", task.RepoName)
        cmd := exec.Command("git", "clone", task.SourceURL)
        cmd.Run()
      }
    }(&wg)
  }

  for i := 0; i < len(repos); i++ {
    tasks <- repos[i]
  }

  close(tasks)

  wg.Wait()
	fmt.Print(ending)
}
