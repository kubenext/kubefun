package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	VERSION    = "v0.0.3"
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
		case "ci":
			test()
			vet()
			webDeps()
			webTest()
			webBuild()
			build()
		case "ci-quick":
			webDeps()
			webBuild()
			build()
		case "web-deps":
			webDeps()
		case "web-test":
			webDeps()
			webTest()
		case "web-build":
			webDeps()
			webBuild()
		case "web":
			webDeps()
			webTest()
			webBuild()
		case "serve":
			serve()
		case "install-test-plugin":
			installTestPlugin()
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

func webDeps() {
	cmd := newCmd("npm", nil, "i")
	cmd.Stdout = os.Stdout
	cmd.Dir = "./web"
	if err := cmd.Run(); err != nil {
		log.Fatalf("web-deps: %s", err)
	}
}

func webTest() {
	cmd := newCmd("npm", nil, "run", "test:headless")
	cmd.Stdout = os.Stdout
	cmd.Dir = "./web"
	if err := cmd.Run(); err != nil {
		log.Fatalf("web-test: %s", err)
	}
}

func webBuild() {
	cmd := newCmd("npm", nil, "run", "build")
	cmd.Stdout = os.Stdout
	cmd.Dir = "./web"
	if err := cmd.Run(); err != nil {
		log.Fatalf("web-build: %s", err)
	}
	runCmd("go", nil, "generate", "./web")
}

func serve() {
	var wg sync.WaitGroup

	uiVars := map[string]string{"API_BASE": "http://localhost:7777"}
	uiCmd := newCmd("npm", uiVars, "run", "start")
	uiCmd.Stdout = os.Stdout
	uiCmd.Dir = "./web"
	if err := uiCmd.Start(); err != nil {
		log.Fatalf("uiCmd: start: %s", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := uiCmd.Wait(); err != nil {
			log.Fatalf("serve: npm run: %s", err)
		}
	}()

	serverVars := map[string]string{
		"KUBEFUN_DISABLE_OPEN_BROWSER": "true",
		"KUBEFUN_LISTENER_ADDR":        "localhost:7777",
		"KUBEFUN_PROXY_FRONTEND":       "http://localhost:4200",
	}
	serverCmd := newCmd("go", serverVars, "run", "./cmd/main/main.go")
	serverCmd.Stdout = os.Stdout
	if err := serverCmd.Start(); err != nil {
		log.Fatalf("serveCmd: start: %s", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serverCmd.Wait(); err != nil {
			log.Fatalf("serve: go run: %s", err)
		}
	}()

	wg.Wait()
}

func installTestPlugin() {
	dir := pluginDir()
	log.Printf("Plugin path: %s", dir)
	os.MkdirAll(dir, 0755)
	pluginFile := fmt.Sprintf("%s/kubefun-sample-plugin", dir)
	runCmd("go", nil, "build", "-o", pluginFile, "github.com/kubenext/kubefun/cmd/kubefun-sample-plugin")
}

func pluginDir() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		return filepath.Join(xdgConfigHome, "kubefun", "plugins")
	} else if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "kubefun", "plugins")
	} else {
		return filepath.Join(os.Getenv("HOME"), ".config", "kubefun", "plugins")
	}
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
