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
	windowsDir := "/src"
	if isNotExist(windowsDir) {
		os.Mkdir(windowsDir, 0755)
		//os.Chdir("/src")
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	macDir := "/src"
	if isNotExist(macDir) {
		os.Mkdir(macDir, 0755)
		//os.Chdir("/src")
	} else {
		log.Println("Error: Couldn't create project.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	linuxDir := "/src"
	if isNotExist(linuxDir) {
		os.Mkdir(linuxDir, 0755)
		//os.Chdir("/src")
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
