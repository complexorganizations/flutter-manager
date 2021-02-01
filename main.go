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

// Install Flutter On Windows
func installFlutterOnWindows() {
	// make sure flutter directory is not there
	if isNotExist("/src/flutter") {
		// make sure flutter isnt there and clone
		if isNotExist("flutter") {
			cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
			cmd.Run()
			// make sure /src is there and if its not make the folder
			if isNotExist("/src") {
				os.Mkdir("/src", 0755)
			} else {
				log.Println("Error: Failed to build a project.")
				os.Exit(0)
			}
			// move the flutter folder to the correct path
			os.Rename("flutter", "/src/flutter")
			os.Setenv("PATH", "/src/flutter/bin")
		} else {
			log.Println("Error: Failed to build a project.")
			os.Exit(0)
		}
	} else {
		log.Println("Error: Flutter discovered in the system.")
		os.Exit(0)
	}
}

// Install Flutter On Mac
func installFlutterOnMac() {
	// make sure flutter directory is not there
	if isNotExist("/usr/local/flutter") {
		// make sure flutter isnt there and clone
		if isNotExist("flutter") {
			cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
			cmd.Run()
			// make sure /usr/local is there and if its not make the folder
			if isNotExist("/usr/local") {
				os.MkdirAll("/usr/local", 0755)
			} else {
				log.Println("Error: Failed to build a project.")
				os.Exit(0)
			}
			// move the flutter folder to the correct path
			os.Rename("flutter", "/usr/local/flutter")
			os.Setenv("PATH", "export PATH=$PATH:/usr/local/flutter/bin")
		} else {
			log.Println("Error: Failed to build a project.")
			os.Exit(0)
		}
	} else {
		log.Println("Error: Flutter discovered in the system.")
		os.Exit(0)
	}
}

// Install Flutter On Linux
func installFlutterOnLinux() {
	// make sure flutter directory is not there
	if isNotExist("/usr/local/flutter") {
		// make sure flutter isnt there and clone
		if isNotExist("flutter") {
			cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
			cmd.Run()
			// make sure /usr/local is there and if its not make the folder
			if isNotExist("/usr/local") {
				os.MkdirAll("/usr/local", 0755)
			} else {
				log.Println("Error: Failed to build a project.")
				os.Exit(0)
			}
			// move the flutter folder to the correct path
			os.Rename("flutter", "/usr/local/flutter")
			os.Setenv("PATH", "export PATH=$PATH:/usr/local/flutter/bin")
		} else {
			log.Println("Error: Failed to build a project.")
			os.Exit(0)
		}
	} else {
		log.Println("Error: Flutter discovered in the system.")
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
