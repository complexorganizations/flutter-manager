package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var (
	flutterPath     = fmt.Sprint(userDirectory() + "/flutter")
	flutterBin      = fmt.Sprint(flutterPath + "/bin")
	flutterTempPath = fmt.Sprint(os.TempDir() + "/flutter")
	unixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
)

func main() {
	selectOperatingSystem()
}

// Choose OS to install flutter
func selectOperatingSystem() {
	switch runtime.GOOS {
	case "windows":
		commandsRequirementsCheck()
		uninstallFlutterOnWindows()
		gitCloneFlutter()
		installFlutterOnWindows()
	case "darwin", "linux":
		commandsRequirementsCheck()
		uninstallFlutterOnUnix()
		gitCloneFlutter()
		installFlutterOnUnix()
	default:
		log.Fatalf("Warning: %s is not supported (yet).\n", runtime.GOOS)
	}
}

// System Requirements Check
func commandsRequirementsCheck() {
	if !folderExists(flutterPath) {
		if commandExists("flutter") {
			log.Fatal("Error: The application flutter was not found in the system.")
		}
	}
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
}

// git clone flutter
func gitCloneFlutter() {
	if !folderExists(os.TempDir()) {
		os.Mkdir(os.TempDir(), 0755)
	}
	if !folderExists(userDirectory()) {
		os.Mkdir(userDirectory(), 0755)
	}
	if folderExists(flutterTempPath) {
		os.RemoveAll(flutterTempPath)
	}
	if !folderExists(flutterTempPath) {
		os.Chdir(os.TempDir())
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Rename(flutterTempPath, flutterPath)
	}
}

// Install Flutter On Windows
func installFlutterOnWindows() {
	if folderExists(flutterPath) {
		path, exists := os.LookupEnv("PATH")
		if exists {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Println(err)
			}
			if !strings.Contains(string(data), "flutter") {
				cmd := exec.Command("setx", "flutter", flutterBin)
				err = cmd.Run()
				if err != nil {
					os.RemoveAll(flutterPath)
					log.Fatal("Error: Failed to write flutter in system path.")
				}
			}
		}
	}
}

// Uninstall flutter on Windows
func uninstallFlutterOnWindows() {
	if folderExists(flutterPath) {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Uninstall Flutter")
		fmt.Println("2. Exit")
		var number int
		fmt.Scanln(&number)
		switch number {
		case 1:
			os.RemoveAll(flutterPath)
			cmd := exec.Command("REG", "delete", "HKCU", `\`, "Environment", "/F /V", "Flutter")
			cmd.Run()
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Warning: this is not a valid response.")
		}
	}
}

// Install Flutter On Linux, Unix
func installFlutterOnUnix() {
	if folderExists(flutterPath) {
		if runtime.GOOS == "darwin" {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
		}
		data, err := os.ReadFile(unixProfilePath)
		if err != nil {
			log.Println(err)
		}
		if !strings.Contains(string(data), "flutter") {
			path, err := os.OpenFile(unixProfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				os.RemoveAll(flutterPath)
				log.Fatal("Error: Failed to write flutter in system path.")
			}
			path.Write([]byte("export PATH=$PATH:" + flutterBin))
			path.Close()
			cmd := exec.Command("source", unixProfilePath)
			cmd.Run()
		}
	}
}

// Uninstall Flutter on Linux
func uninstallFlutterOnUnix() {
	if folderExists(flutterPath) {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Uninstall Flutter")
		fmt.Println("2. Exit")
		var number int
		fmt.Scanln(&number)
		switch number {
		case 1:
			os.RemoveAll(flutterPath)
			if runtime.GOOS == "darwin" {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
			}
			data, err := os.ReadFile(unixProfilePath)
			if err != nil {
				log.Println(err)
			}
			if strings.Contains(string(data), "flutter") {
				read, err := os.ReadFile(unixProfilePath)
				if err != nil {
					log.Println(err)
				}
				newContents := strings.Replace(string(read), ("export PATH=$PATH:" + flutterBin), (""), -1)
				os.WriteFile(unixProfilePath, []byte(newContents), 0)
			}
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Warning: this is not a valid response.")
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

// Check if there is an app installed
func commandExists(cmd string) bool {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = cmd // variable declared and not used
	return true
}

// Get the current user dir
func userDirectory() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

// cmd = exec.Command("flutter", "config", "--enable-web")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-linux-desktop")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-macos-desktop")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-windows-desktop")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-android")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-ios")
// cmd.Run()
// cmd = exec.Command("flutter", "config", "--enable-fuchsia")
// cmd.Run()
// cmd = exec.Command("flutter", "upgrade")
// cmd.Run()
