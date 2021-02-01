package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	selectOperatingSystem()
}

// Choose OS to install flutter
func selectOperatingSystem() {
	switch runtime.GOOS {
	case "windows":
		commandsRequirementsCheck()
		installFlutterOnWindows()
	case "darwin":
		commandsRequirementsCheck()
		installFlutterOnMac()
	case "linux":
		commandsRequirementsCheck()
		installFlutterOnLinux()
	default:
		fmt.Printf("Error: System %s Not Supported.\n", runtime.GOOS)
	}
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

// Install Flutter on Windows
func installFlutterOnWindows() {
	if isNotExist("flutter") {
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		if isNotExist("/src") {
			os.Mkdir("/src", 0755)
		}
		os.Rename("flutter", "/src/flutter")
	} else {
		log.Println("Error: Failed to build a project.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	if isNotExist("flutter") {
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		if isNotExist("/usr/local") {
			os.Mkdir("/usr/local", 0755)
		}
		os.Rename("flutter", "/usr/local/flutter")
	} else {
		log.Println("Error: Failed to build a project.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	if isNotExist("flutter") {
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		if isNotExist("/usr/local") {
			os.Mkdir("/usr/local", 0755)
		}
		os.Rename("flutter", "/usr/local/flutter")
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
