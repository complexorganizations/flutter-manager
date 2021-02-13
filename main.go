package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"io/ioutil"
	"runtime"
)

var (
	flutterSource    = "/src/"
	flutterPath      = "/src/flutter/"
	flutterManager   = "/src/flutter/flutter-manager"
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
		uninstallFlutter()
		gitCloneFlutter()
		installFlutterOnWindows()
	case "darwin":
		commandsRequirementsCheck()
		uninstallFlutter()
		gitCloneFlutter()
		installFlutterOnMac()
	case "linux":
		commandsRequirementsCheck()
		uninstallFlutter()
		gitCloneFlutter()
		installFlutterOnLinux()
	default:
		fmt.Printf("Error: System %s Not Supported.\n", runtime.GOOS)
		os.Exit(0)
	}
}

// System Requirements Check
func commandsRequirementsCheck() {
	if commandExists("flutter") {
		log.Fatal("Error: Flutter command discovered in the system.")
	}
	if !commandExists("git") {
		log.Fatal("Error: Git was not discovered in the system.")
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
		ioutil.WriteFile(flutterManager, []byte("flutter-manager: true"), 0644)
	} else {
		os.Chdir(systemTempFolder)
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Mkdir(flutterSource, 0755)
		os.Rename(tempFlutterPath, flutterPath)
		ioutil.WriteFile(flutterManager, []byte("flutter-manager: true"), 0644)
	}
}

// Install Flutter On Windows
func installFlutterOnWindows() {
	if folderExists(flutterPath) {
		cmd := exec.Command("setx", "path", flutterBin)
		err := cmd.Run()
		if err != nil {
			os.RemoveAll(flutterPath)
			log.Fatal("Error: Failed to write system path.")
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
			os.RemoveAll(flutterPath)
			log.Fatal("Error: Failed to write system path.")
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
			os.RemoveAll(flutterPath)
			log.Fatal("Error: Failed to write system path.")
		}
	}
}

// Uninstall Flutter
func uninstallFlutter() {
	if fileExists(flutterManager) {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Uninstall Flutter")
		fmt.Println("2. Exit")
		var number int
		fmt.Scanln(&number)
		switch number {
		case 1:
			os.RemoveAll(flutterPath)
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Error: this is not a valid response.")
		}
	}
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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
