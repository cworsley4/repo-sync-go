
package main

import (
  "fmt"
  "bytes"
  "os/exec"
)

func Pull(task *Repo) {
  var out bytes.Buffer
  var stderr bytes.Buffer

  fmt.Print("Getting the status", task.LocalPath)

  status := exec.Command("git", "-C", task.LocalPath, "status")

  status.Stdout = &out
  status.Stderr = &stderr

  err := status.Run()
  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
  }

  pull := exec.Command("git", "-C", task.LocalPath, "pull")
  fmt.Println("Trying to pull on repo: git pull ", task.LocalPath)
  pull.Stdout = &out
  pull.Stderr = &stderr
  err = pull.Run()

  fmt.Print("Pulled", task.LocalPath)

  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
  }

  return
}
