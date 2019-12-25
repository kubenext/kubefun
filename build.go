package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var (
	VERSION    = "v0.0.1"
	GOPATH     = os.Getenv("GOPATH")
	GIT_COMMIT = gitCommit()
	BUILD_TIME = time.Now().UTC().Format(time.RFC3339)
	LD_FLAGS   = fmt.Sprintf("-X main.version=%s -X main.buildTime=%s -X main.gitCommit=%s", VERSION, BUILD_TIME, GIT_COMMIT)
	GO_FLAGS   = fmt.Sprintf("-ldflags=%s", LD_FLAGS)
)

// TODO web build

func main() {
	flag.Parse()
	for _, cmd := range flag.Args() {
		switch cmd {
		case "go-install":
			goInstall()
		case "clean":
			clean()
		case "generate":
			generate()
		case "build":
			build()
		case "run-dev":
			runDev()
		case "test":
			test()
		case "vet":
			vet()
		case "version":
			version()
		case "release":
			release()
		case "docker":
			docker()
		default:
			log.Fatalf("Unknown command %q", cmd)
		}
	}
}

func goInstall() {
	pkgs := []string{
		"github.com/GeertJohan/go.rice",
		"github.com/GeertJohan/go.rice/rice",
		"github.com/golang/mock/gomock",
		"github.com/golang/mock/mockgen",
		"github.com/golang/protobuf/protoc-gen-go",
	}
	for _, pkg := range pkgs {
		runCmd("go", map[string]string{"GO11MODULE": "on"}, "install", pkg)
	}
}

func clean() {
	if err := os.Remove("pkg/icon/rice-box.go"); err != nil {
		log.Fatalf("clean: %s", err)
	}
}

func generate() {
	removeFakes()
	runCmd("go", nil, "generate", "-v", "./pkg/...", "./internal/...")
}

func build() {
	newpath := filepath.Join(".", "build")
	os.MkdirAll(newpath, 0755)

	artifact := "kf"
	if runtime.GOOS == "windows" {
		artifact = "kf.exe"
	}
	runCmd("go", nil, "build", "-o", "build/"+artifact, GO_FLAGS, "-v", "./cmd/main")
}

func runDev() {
	runCmd("build/kf", nil)
}

func test() {
	runCmd("go", nil, "test", "-v", "./internal/...", "./pkg/...")
}

func vet() {
	runCmd("go", nil, "vet", "./internal/...", "./pkg/...")
}

func version() {
	fmt.Println(VERSION)
}

func release() {
	runCmd("git", nil, "tag", "-a", VERSION, "-m", fmt.Sprintf("\"Release %s\"", VERSION))
	runCmd("git", nil, "push", "--follow-tags")
}

func docker() {
	dockerVars := map[string]string{
		"CGO_ENABLED": "0",
		"GOOS":        "linux",
		"GOARCH":      "amd64",
	}
	runCmd("go", dockerVars, "build", "-o", "kf", GO_FLAGS, "-v", "./cmd/main/main.go")
}

func removeFakes() {
	checkDirs := []string{"pkg", "internal"}
	fakePaths := []string{}

	for _, dir := range checkDirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}
			if info.Name() == "fake" {
				fakePaths = append(fakePaths, filepath.Join(path, info.Name()))
			}
			return nil
		})
		if err != nil {
			log.Fatalf("generate (%s): %s", dir, err)
		}
	}

	log.Print("Removing fakes from pkg/ and internal/")
	for _, p := range fakePaths {
		os.RemoveAll(p)
	}
}

func gitCommit() string {
	cmd := newCmd("git", nil, "rev-parse", "--short", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("gitCommit: %s", err)
		return ""
	}
	return fmt.Sprintf("%s", out)
}

func runCmd(command string, env map[string]string, args ...string) {
	cmd := newCmd(command, env, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Running: %s\n", cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func newCmd(command string, env map[string]string, args ...string) *exec.Cmd {
	realCommand, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("unable to find command '%s'", command)
	}
	cmd := exec.Command(realCommand, args...)
	cmd.Env = os.Environ()

	cmd.Stderr = os.Stderr
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	return cmd
}
