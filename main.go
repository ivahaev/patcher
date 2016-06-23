package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	version  = "0.0.6"
	filename = "main.go"
	varName  = "version"
)

func main() {
	args := os.Args
	if len(args) > 1 && args[1] != "-p" {
		filename = args[1]
	}
	if len(args) > 2 && args[2] != "-p" {
		varName = args[2]
	}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Can't open file: %v\n", err)
		return
	}
	stat, err := f.Stat()
	if err != nil {
		fmt.Printf("Can't read file stat: %v\n", err)
		return
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Can't read file: %v\n", err)
		return
	}

	rgxVersion := regexp.MustCompile(`\s+` + varName + `\s+=\s+"\d+\.\d+\.\d+"`)
	matched := string(rgxVersion.Find(file))
	if matched == "" {
		fmt.Println("Version not found")
		return
	}
	rgxNum := regexp.MustCompile(`\d+\.\d+\.\d+`)
	oldVersion := rgxNum.FindString(matched)
	if oldVersion == "" {
		fmt.Println("Can't extract version")
		return
	}
	subVersions := strings.Split(oldVersion, ".")
	patchVersion, _ := strconv.Atoi(subVersions[2])
	patchVersion++
	subVersions[2] = strconv.Itoa(patchVersion)
	newVersion := strings.Join(subVersions, ".")
	matched = strings.Replace(matched, oldVersion, newVersion, 1)
	file = rgxVersion.ReplaceAll(file, []byte(matched))
	err = ioutil.WriteFile(filename, file, stat.Mode())
	if err != nil {
		fmt.Printf("Can't write file: %v\n", err)
		return
	}
	cmd := exec.Command("git", "add", filename)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Add add file %s to git index: %v\n", filename, err)
		return
	}
	cmd = exec.Command("git", "commit", "-m", "Version update: "+newVersion)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Add commit file %s : %v\n", filename, err)
		return
	}
	cmd = exec.Command("git", "tag", "-a", "v"+newVersion, "-m", "Version: "+newVersion)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Add git tag: %v\n", err)
		return
	}
	fmt.Printf("Version updated to: %s\n", newVersion)
	push := flag.Bool("-p", false, "Type -p flag to push changes immediatly")
	flag.Parse()
	if !*push {
		return
	}
	cmd = exec.Command("git", "push", "origin", "master", "--follow-tags")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Can't push changes: %v\n", err)
		return
	}
	fmt.Println("Pushed to repo")
}
