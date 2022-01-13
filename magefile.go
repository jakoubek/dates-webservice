// +build mage

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

const (
	BUILD_DIR     string = "./bin"
	BUILD_BINARY  string = "dates-webservice"
	DEPLOY_TARGET string = "oli@opal5.opalstack.com:apps/datesapi/dates-webservice.new"
	DEPLOY_DIR    string = "oli@opal5.opalstack.com:apps/datesapi/"
)

var (
	buildVersion string
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	os.Setenv("TZ", "Europe/Berlin")

	mg.Deps(InstallDeps)
	fmt.Println("Building...")

	versionCmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	versionCmdOutput, err := versionCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	buildVersion = string(versionCmdOutput)
	buildTime := time.Now().UTC().Format("2006-01-02_15:04:05")
	//buildTime = "Hallo"

	cmd := exec.Command("go", "build", "-ldflags", "-X main.version="+buildVersion+" -X main.buildTime="+buildTime, "-o", path.Join(BUILD_DIR, BUILD_BINARY), ".")
	return cmd.Run()
}

func Debugrun() error {

	fmt.Println("Building and running locally...")

	cmd := exec.Command("go", "build", "-o", "dates-webservice.exe", ".")
	cmd.Run()

	cmd = exec.Command("dates-webservice.exe")

	return cmd.Run()
}

// Copies binary, template and config file to Virtualbox Programme folder
func Deploy() error {
	mg.Deps(Build)
	fmt.Println("Deploying...")
	cmd := exec.Command("scp", path.Join(BUILD_DIR, BUILD_BINARY), DEPLOY_TARGET)
	cmd.Run()
	fmt.Println("- binary ok")

	//cmd  = exec.Command("scp", "config.json", DEPLOY_TARGET)
	//err := sh.Copy(path.Join(DEPLOY_DIR, BUILD_BINARY), path.Join(BUILD_DIR, BUILD_BINARY))
	//_ = sh.Copy(path.Join(BUILD_DIR, "auflagenmeldung.html"), "auflagenmeldung.html")
	//cmd = exec.Command("scp", "-r", "views/", DEPLOY_DIR)
	//cmd.Run()
	//fmt.Println("- views ok")

	return nil
}

func Restart() error {
	fmt.Println("Restarting server...")
	cmd := exec.Command("ssh", "mopo@opal5.opalstack.com", "views/", DEPLOY_TARGET)
	cmd.Run()
	return nil
}

// A custom install step if you need your bin someplace other than go/bin
//func Install() error {
//	mg.Deps(Build)
//	fmt.Println("Installing...")
//	return os.Rename("./MyApp", "/usr/bin/MyApp")
//}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	//return exec.Command("dep", "ensure").Run()
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("bin")
}
