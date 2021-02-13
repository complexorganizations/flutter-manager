package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var (
	flutterSource = userDir()
	flutterPath   = fmt.Sprint(userDir(), "/flutter")
	flutterBin    = fmt.Sprint(flutterPath, "/bin")
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
	case "darwin":
		commandsRequirementsCheck()
		uninstallFlutterOnUnix()
		gitCloneFlutter()
		installFlutterOnUnix()
	case "linux":
		commandsRequirementsCheck()
		uninstallFlutterOnUnix()
		gitCloneFlutter()
		installFlutterOnUnix()
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
	if !folderExists(os.TempDir()) {
		os.Mkdir(os.TempDir(), 0755)
		os.Chdir(os.TempDir())
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Mkdir(flutterSource, 0755)
		os.Rename("flutter", flutterPath)
	} else {
		os.Chdir(os.TempDir())
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		cmd.Run()
		os.Mkdir(flutterSource, 0755)
		os.Rename("flutter", flutterPath)
	}
}

// Install Flutter On Windows
func installFlutterOnWindows() {
	if folderExists(flutterPath) {
		cmd := exec.Command("setx", "Flutter", flutterBin)
		err := cmd.Run()
		if err != nil {
			os.RemoveAll(flutterPath)
			log.Fatal("Error: Failed to write system path.")
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
			cmd := exec.Command(`REG delete HKCU \ Environment / F / V Flutter`)
			err := cmd.Run()
			if err != nil {
				log.Fatal("Error: Failed to remove system path.")
			}
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Error: this is not a valid response.")
		}
	}
}

// Install Flutter On Linux, Mac
func installFlutterOnUnix() {
	if folderExists(flutterPath) {
		path, err := os.OpenFile("/etc/profile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		path.Write([]byte("export PATH=$PATH:", flutterBin))
		path.Close()
		if err != nil {
			os.RemoveAll(flutterPath)
			log.Fatal("Error: Failed to write system path.")
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
			read, err := ioutil.ReadFile("/etc/profile")
			if err != nil {
				log.Println(err)
			}
			newContents := strings.Replace(string(read), ("export PATH=$PATH:", flutterBin), (""), -1)
			ioutil.WriteFile("/etc/profile", []byte(newContents), 0)
		case 2:
			os.Exit(0)
		default:
			fmt.Println("Error: this is not a valid response.")
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

// Get the current user dir
func userDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
