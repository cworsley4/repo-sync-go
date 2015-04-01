
package main

import (
  "fmt"
  "os/exec"
)

func Clone(task *Repo) {
  fmt.Println("Cloning... %s", task.RepoName, "to", task.LocalPath, "from", task.SourceURL)
  cmd := exec.Command("git", "clone", task.SourceURL, task.LocalPath)
  cmd.Run()
  return
}
