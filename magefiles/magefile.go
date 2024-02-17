package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	BUILD_DIR         string = "./dist"
	BUILD_BINARY      string = "datesapi-server"
	SOURCE_ENTRY      string = "./cmd/server/."
	LOCAL_DEPLOY_PATH string = "D:/VirtualboxOrdner/PSG_Serverprogramme/psg-setbildung-arvato"
)

var Default = Run

func Build() error {
	mg.Deps(Clean)

	fmt.Println("Building...")

	gocmd := mg.GoCmd()

	return sh.RunV(gocmd, "build", "-o", path.Join(BUILD_DIR, BUILD_BINARY), "-ldflags="+flags(), SOURCE_ENTRY)
}

func Deploy() error {
	//mg.Deps(Build, CopyAdditional)
	mg.Deps(Build)

	fmt.Println("Deploy locally...")
	return sh.Copy(path.Join(LOCAL_DEPLOY_PATH, BUILD_BINARY), path.Join(BUILD_DIR, BUILD_BINARY))

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
