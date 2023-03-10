package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	switch os.Args[1] {
	case "init":
		initCmd()
	case "add":
		addCmd()
	case "del":
		delCmd()
	case "ls":
		lsCmd()
	case "clone":
		cloneCmd()
	case "update":
		updateCmd()
	case "help":
		helpCmd()
	default:
		help()
	}
}

func initCmd() {
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(gitList), 0755)
		ioutil.WriteFile(gitList, []byte("{}"), 0644)
	} else {
		fmt.Println("git-list.json already exists")
	}
}

func addCmd() {
	if len(os.Args) < 4 {
		fmt.Println("ait add repo url")
		return
	}
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		fmt.Println("git-list.json not exists")
		return
	}
	jsonData, err := ioutil.ReadFile(gitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonObj map[string]string
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonObj[os.Args[2]] = os.Args[3]
	jsonData, err = json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(gitList, jsonData, 0644)
}

func delCmd() {
	if len(os.Args) < 3 {
		fmt.Println("ait del repo")
		return
	}
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		fmt.Println("git-list.json not exists")
		return
	}
	jsonData, err := ioutil.ReadFile(gitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonObj map[string]string
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	delete(jsonObj, os.Args[2])
	jsonData, err = json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(gitList, jsonData, 0644)
}

func lsCmd() {
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		fmt.Println("git-list.json not exists")
		return
	}
	jsonData, err := ioutil.ReadFile(gitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonObj map[string]string
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	for repo, url := range jsonObj {
		fmt.Println(repo, url)
	}
}

func cloneCmd() {
	if len(os.Args) < 3 {
		fmt.Println("ait clone repo")
		return
	}
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		fmt.Println("git-list.json not exists")
		return
	}
	jsonData, err := ioutil.ReadFile(gitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonObj map[string]string
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	if os.Args[2] == "all" {
		for _, url := range jsonObj {
			clone(url)
		}
	} else {
		clone(jsonObj[os.Args[2]])
	}
}

func clone(url string) {
	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func updateCmd() {
	if len(os.Args) < 4 {
		fmt.Println("ait update repo url")
		return
	}
	var home, _ = os.UserHomeDir()
	gitList := filepath.Join(home, ".ait", "git-list.json")
	if _, err := os.Stat(gitList); os.IsNotExist(err) {
		fmt.Println("git-list.json not exists")
		return
	}
	jsonData, err := ioutil.ReadFile(gitList)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonObj map[string]string
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonObj[os.Args[2]] = os.Args[3]
	jsonData, err = json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(gitList, jsonData, 0644)
}

func help() {
	fmt.Println("ait init")
	fmt.Println("ait add repo url")
	fmt.Println("ait del repo")
	fmt.Println("ait ls")
	fmt.Println("ait clone repo")
	fmt.Println("ait clone all")
	fmt.Println("ait update repo url")
	fmt.Println("ait help")
}

func helpCmd() {
	if len(os.Args) < 3 {
		help()
		return
	}
	switch os.Args[2] {
	case "add":
		fmt.Println("ait add repo url")
	case "del":
		fmt.Println("ait del repo")
	case "ls":
		fmt.Println("ait ls")
	case "clone":
		fmt.Println("ait clone repo")
		fmt.Println("ait clone all")
	case "update":
		fmt.Println("ait update repo url")
	case "help":
		fmt.Println("ait help")
		fmt.Println("ait help add")
		fmt.Println("ait help del")
		fmt.Println("ait help ls")
		fmt.Println("ait help clone")
		fmt.Println("ait help update")
	default:
		help()
	}
}
