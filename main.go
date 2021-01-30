package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	requirementsCheck()
}

// System Requirements Check
func requirementsCheck() {
	if commandExists("flutter") {
		log.Println("Error: Flutter found on the system")
		os.Exit(0)
	}
	if commandExists("git") {
		selectOperatingSystem()
	}

}

// Choose OS to install flutter
func selectOperatingSystem() {
	os := runtime.GOOS
	switch os {
	case "windows":
		installFlutterOnWindows()
	case "darwin":
		installFlutterOnMac()
	case "linux":
		installFlutterOnLinux()
	default:
		fmt.Printf("Error: %s Not Supported.\n", os)
	}
}

// Install Flutter on Windows
func installFlutterOnWindows() {
	windowsDir := "/src/flutter"
	if isNotExist(windowsDir) {
		exec.Command("git", "clone", "git@github.com:flutter/flutter.git", "/src/flutter", "-b stable")
		exec.Command("setx", "path", "/src/flutter/bin")
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	macDir := "/usr/local/flutter"
	if isNotExist(macDir) {
		exec.Command("git", "clone", "git@github.com:flutter/flutter.git", "/usr/local/flutter", "-b stable")
		ioutil.WriteFile("~/.profile", []byte("export PATH=$PATH:/usr/local/flutter/bin"), 0644)
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	linuxDir := "/usr/local/flutter"
	if isNotExist(linuxDir) {
		exec.Command("git", "clone", "git@github.com:flutter/flutter.git", "/usr/local/flutter", "-b stable")
		ioutil.WriteFile("~/.profile", []byte("export PATH=$PATH:/usr/local/flutter/bin"), 0644)
	} else {
		log.Println("Error: Couldn't create project.")
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
