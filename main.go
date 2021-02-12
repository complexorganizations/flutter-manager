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

// git clone flutter
func gitCloneFlutter() {
	if !folderExists(systemTempFolder) {
		os.Mkdir(systemTempFolder, 0755)
		os.Chdir(systemTempFolder)
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Mkdir(flutterSource, 0755)
		os.Rename(tempFlutterPath, flutterPath)
	} else {
		os.Chdir(systemTempFolder)
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Mkdir(flutterSource, 0755)
		os.Rename(tempFlutterPath, flutterPath)
	}
}

// Install Flutter On Windows
func installFlutterOnWindows() {
	if folderExists(flutterPath) {
		cmd := exec.Command("setx", "path", flutterBin)
		err := cmd.Run()
		if err != nil {
			log.Println("Error: Failed to write system path.")
			os.RemoveAll(flutterPath)
			os.Exit(0)
		}
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	if folderExists(flutterPath) {
		path, err := os.OpenFile("/etc/profile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		path.Write([]byte("export PATH=$PATH:/src/flutter/bin\n"))
		path.Close()
		if err != nil {
			log.Println("Error: Failed to write system path.")
			os.RemoveAll(flutterPath)
			os.Exit(0)
		}
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	if folderExists(flutterPath) {
		path, err := os.OpenFile("/etc/profile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		path.Write([]byte("export PATH=$PATH:/src/flutter/bin\n"))
		path.Close()
		if err != nil {
			log.Println("Error: Failed to write system path.")
			os.RemoveAll(flutterPath)
			os.Exit(0)
		}
	}
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Check if a command exists
func commandExists(cmd string) bool {
        cmd, err := exec.LookPath(cmd)
        if err != nil {
                return false
        }
        return true
}
