package main

import (
  "fmt"
  "os/exec"
  "io/ioutil"
  "encoding/json"
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

func clone(repo Repo) {
  cmd := exec.Command("git", "clone", repo.SourceURL)
  cmd.Start()
  cmd.Wait()
  fmt.Print(cmd)
}

func main() {
  var repos []Repo
  r, _ := ioutil.ReadFile("./repos.json")
  err := json.Unmarshal(r, &repos)

  if err != nil {
    panic(err)
  }

  for key, _ := range repos {
    fmt.Print("\n")
    repo := repos[key]
    go func(repo Repo) {
      cmd := exec.Command("git", "clone", repo.SourceURL)
      cmd.Start()
      cmd.Wait()
      fmt.Print(cmd)
    }(&repo)
  }

  fmt.Print(ending)

  select{}
}
