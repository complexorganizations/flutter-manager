package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	selectOperatingSystem()
}

// System Requirements Check
func commandsRequirementsCheck() {
	if commandExists("flutter") {
		log.Println("Error: Flutter discovered in the system.")
		os.Exit(0)
	}
	if !commandExists("git") {
		log.Println("Error: Git was not discovered in the system.")
		os.Exit(0)
	}
}

func pathCheck() {
	data, err := ioutil.ReadFile("~/.profile")
	fileData := string(data)
	if strings.Contains(fileData, "flutter") {
		log.Println("Error: Flutter discovered in your path.", err)
		os.Exit(0)
	} else {
		selectOperatingSystem()
	}
}

// Choose OS to install flutter
func selectOperatingSystem() {
	switch runtime.GOOS {
	case "windows":
		commandsRequirementsCheck()
		installFlutterOnWindows()
	case "darwin":
		commandsRequirementsCheck()
		pathCheck()
		installFlutterOnMac()
	case "linux":
		commandsRequirementsCheck()
		pathCheck()
		installFlutterOnLinux()
	default:
		fmt.Printf("Error: System %s Not Supported.\n", runtime.GOOS)
	}
}

// Install Flutter on Windows
func installFlutterOnWindows() {
	if isNotExist("/src/flutter") {
		exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "/src/flutter", "-b stable")
		exec.Command("setx", "path", "/src/flutter/bin")
	} else {
		log.Println("Error: Failed to build a project.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	if isNotExist("/usr/local/flutter") {
		exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "/usr/local/flutter", "-b stable")
		ioutil.WriteFile("~/.profile", []byte("export PATH=$PATH:/usr/local/flutter/bin"), 0644)
		exec.Command("source", "~/.profile")
	} else {
		log.Println("Error: Failed to build a project.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	if isNotExist("/usr/local/flutter") {
		exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "/usr/local/flutter", "-b stable")
		ioutil.WriteFile("~/.profile", []byte("export PATH=$PATH:/usr/local/flutter/bin"), 0644)
		exec.Command("source", "~/.profile")
	} else {
		log.Println("Error: Failed to build a project.")
		os.Exit(0)
	}
}

// Check if a directory exists
func isNotExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return true
	}
	return !info.IsDir()
}

// Check if a command exists
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
