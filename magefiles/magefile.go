package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	BUILD_DIR         string = "./dist"
	SOURCE_ENTRY      string = "./cmd/server/."
	LOCAL_DEPLOY_PATH string = "D:/VirtualboxOrdner/PSG_Serverprogramme/psg-setbildung-arvato"
)

var buildBinary string
var buildBinaryLinux string

func init() {
	buildBinaryLinux = "datesapi-server"
	buildBinary = buildBinaryLinux + ".exe"
}

var Default = Run

// Build builds the application for the current platform
func Build() error {
	switch runtime.GOOS {
	case "linux":
		return BuildLinux()
	case "windows":
		return BuildWindows()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func BuildLinux() error {
	mg.Deps(Clean)

	fmt.Println("Building for Linux...")

	os.MkdirAll(BUILD_DIR, 0755)

	gocmd := mg.GoCmd()

	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "amd64",
	}

	return sh.RunWithV(env, gocmd, "build", "-o", path.Join(BUILD_DIR, buildBinaryLinux), "-ldflags="+flags(), SOURCE_ENTRY)
}

// BuildWindows builds the application for a Windows target
func BuildWindows() error {
	mg.Deps(Clean)

	fmt.Println("Building for Windows...")

	os.MkdirAll(BUILD_DIR, 0755)

	gocmd := mg.GoCmd()

	env := map[string]string{
		"GOOS":   "windows",
		"GOARCH": "amd64",
	}

	return sh.RunWithV(env, gocmd, "build", "-o", path.Join(BUILD_DIR, buildBinary), "-ldflags="+flags(), SOURCE_ENTRY)
}

func CopyAdditional() error {

	fmt.Println("Copy config file and email template...")
	err := sh.Copy(path.Join(LOCAL_DEPLOY_PATH, "config.toml"), "config.toml")
	//err = sh.Copy(path.Join(LOCAL_DEPLOY_PATH, "emailtext.html"), "emailtext.html")

	return err
}

func Run() error {
	fmt.Println("Building and running locally...")

	cmd := exec.Command("go", "run", "./cmd/server/main.go")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

	return nil
}

func Clean() error {
	fmt.Println("Cleaning...")
	return os.RemoveAll(BUILD_DIR)
}

func flags() string {
	return fmt.Sprintf(`-w -s`)
}
