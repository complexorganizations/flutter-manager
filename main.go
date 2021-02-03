package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var (
	flutterSource    = "/src/"
	flutterPath      = "/src/flutter/"
	flutterBin       = "/src/flutter/bin/"
	systemTempFolder = "/tmp/"
	tempFlutterPath  = "/tmp/flutter/"
)

func main() {
	selectOperatingSystem()
}

// Choose OS to install flutter
func selectOperatingSystem() {
	switch runtime.GOOS {
	case "windows":
		commandsRequirementsCheck()
		gitCloneFlutter()
		installFlutterOnWindows()
	case "darwin":
		commandsRequirementsCheck()
		gitCloneFlutter()
		installFlutterOnMac()
	case "linux":
		commandsRequirementsCheck()
		gitCloneFlutter()
		installFlutterOnLinux()
	default:
		fmt.Printf("Error: System %s Not Supported.\n", runtime.GOOS)
	}
}

// System Requirements Check
func commandsRequirementsCheck() {
	if commandExists("flutter") {
		log.Println("Error: Flutter command discovered in the system.")
		os.Exit(0)
	}
	if !commandExists("git") {
		log.Println("Error: Git was not discovered in the system.")
		os.Exit(0)
	}
}

func gitCloneFlutter() {
	if isNotExist(systemTempFolder) {
		os.Mkdir(systemTempFolder, 0755)
		if isNotExist(tempFlutterPath) {
			os.Chdir(systemTempFolder)
			cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
			cmd.Run()
			os.Mkdir(flutterSource, 0755)
			os.Rename(tempFlutterPath, flutterPath)
		}
	} else {
		os.Chdir(systemTempFolder)
		if isNotExist(tempFlutterPath) {
			cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
			cmd.Run()
			os.Mkdir(flutterSource, 0755)
			os.Rename(tempFlutterPath, flutterPath)
		}
	}
}

// Install Flutter On Windows
func installFlutterOnWindows() {
	if !isNotExist(flutterPath) {
		cmd = exec.Command("setx", "path", flutterBin)
		cmd.Run()
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	if !isNotExist(flutterPath) {
		path, err := os.OpenFile("/etc/profile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		path.Write([]byte("export PATH=$PATH:/src/flutter/bin\n"))
		path.Close()
		if err != nil {
			log.Println("Error: Failed to write path /etc/profile.")
			os.Exit(0)
		}
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	if !isNotExist(flutterPath) {
		path, err := os.OpenFile("/etc/profile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		path.Write([]byte("export PATH=$PATH:/src/flutter/bin\n"))
		path.Close()
		if err != nil {
			log.Println("Error: Failed to write path /etc/profile.")
			os.Exit(0)
		}
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
