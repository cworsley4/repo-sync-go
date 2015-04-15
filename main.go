package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "os/exec"
	"sync"
)

type Repo struct {
	RepoID         string
	RepoName       string
	RepoDesc       string
	SourceURL      string
	LocalPath      string
	Active         string
	IntranetUserID string
}

func main() {
	signature, _ := ioutil.ReadFile("./signature.txt")

	var repos []Repo
	newRepos := []string{}
	updatedRepos := []string{}

	me := Whoami()
	url := "http://intranet.redventures.net/admin/dev/repo_config/repo_ajax.php?action=get-user-repos&username="
	url += me

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &repos)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	var tasks = make(chan Repo)

	maxWorkers := len(repos)

	if maxWorkers > 10 {
		maxWorkers = 10
	}

	fmt.Println("Spinning up", maxWorkers, "workers")

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		fmt.Println("Adding Worker", i)
		go func(wg *sync.WaitGroup) {
			for task := range tasks {
				_, err := os.Stat(task.LocalPath)

				if err == nil {
					updatedRepos = append(updatedRepos, task.RepoName)
					Pull(&task)
					fmt.Println("Pull complete for", task.RepoName)
				} else {
					newRepos = append(newRepos, task.RepoName)
					Clone(&task)
					fmt.Println("Cloning complete for", task.RepoName)
				}

			}

			wg.Done()
		}(&wg)
	}

	for i := 0; i < len(repos); i++ {
		if len(repos[i].IntranetUserID) > 0 {
			tasks <- repos[i]
		}
	}

	fmt.Println("New Repos: %s", len(newRepos))
	fmt.Println("Updated Repos: %s", len(updatedRepos))

	close(tasks)

	wg.Wait()
	fmt.Print(string(signature))
}
