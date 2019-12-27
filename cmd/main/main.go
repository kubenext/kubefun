package main

import (
	"github.com/kubenext/kubefun/internal/commands"
	"math/rand"
	"time"
)

var (
	version   = "dev-version"
	gitCommit = "dev-commit"
	buildTime = "dev-build-time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	commands.Execute(version, gitCommit, buildTime)
}
