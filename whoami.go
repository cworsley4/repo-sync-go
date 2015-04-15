package main

import (
	"os/exec"
	"strings"
)

func Whoami() string {
	whoami := exec.Command("whoami")
	output, err := whoami.Output()

	if err != nil {
		panic(err)
	}

	return strings.Trim(string(output), " \n")
}
