package main

import (
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

}