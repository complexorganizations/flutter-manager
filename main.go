package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	flutterInstalledCheck()
}

func flutterInstalledCheck() {
	if commandExists("flutter") {
		log.Println("Error: Flutter found on the system")
        } else {
		gitInstalledCheck()
        }
}

func gitInstalledCheck() {
	if commandExists("git") {
		selectOperatingSystem()
        } else {
		log.Println("Error: Git not found on the system")
		os.Exit(1)
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
	windowsDir := "/src"
	if isNotExist(windowsDir) {
		//
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	macDir := "/usr/local/flutter"
	if isNotExist(macDir) {
		//
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	linuxDir := "/usr/local/flutter"
	if isNotExist(linuxDir) {
		//
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// GLOBAL CHECKS

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
