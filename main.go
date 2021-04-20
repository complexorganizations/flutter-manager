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
	flutterPath = fmt.Sprint(userDirectory() + "/flutter")
	flutterBin  = fmt.Sprint(flutterPath + "/bin")
)

var tempUnixProfilePath string
var unixProfilePath string
var err error

func init() {
	// System Requirements Check
	if !folderExists(flutterPath) {
		if commandExists("flutter") {
			log.Fatal("Error: The application flutter was not found in the system.")
		}
	}
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
}

func main() {
	selectOperatingSystem()
}

// Choose OS to install flutter
func selectOperatingSystem() {
	switch runtime.GOOS {
	case "windows":
		uninstallFlutterOnDOS()
		gitCloneFlutter()
		installFlutterOnDOS()
	case "darwin", "linux":
		uninstallFlutterOnUnix()
		gitCloneFlutter()
		installFlutterOnUnix()
	default:
		log.Fatalf("Warning: %s is not supported (yet).\n", runtime.GOOS)
	}
}

// git clone flutter
func gitCloneFlutter() {
	if !folderExists(userDirectory()) {
		os.Mkdir(userDirectory(), 0755)
	}
	if !folderExists(flutterPath) {
		os.Chdir(flutterPath)
		cmd := exec.Command("git", "clone", "https://github.com/flutter/flutter.git", "-b", "stable")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Install Flutter On Windows
func installFlutterOnDOS() {
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
					log.Fatal("Error: Failed to write flutter in system path.")
				}
			}
		}
	}
}

// Uninstall flutter on Windows
func uninstallFlutterOnDOS() {
	if folderExists(flutterPath) {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Uninstall Flutter")
		fmt.Println("2. Exit")
		var number int
		fmt.Scanln(&number)
		switch number {
		case 1:
			err = os.RemoveAll(flutterPath)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("REG", "delete", "HKCU", `\`, "Environment", "/F /V", "Flutter")
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
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
		tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
		if fileExists(tempUnixProfilePath) {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
		}
		tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.bash_profile")
		if fileExists(tempUnixProfilePath) {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.bash_profile")
		}
		tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.bashrc")
		if fileExists(tempUnixProfilePath) {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.bashrc")
		}
		tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
		if fileExists(tempUnixProfilePath) {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
		}
		if !fileExists(tempUnixProfilePath) {
			unixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
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
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
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
			err = os.RemoveAll(flutterPath)
			if err != nil {
				log.Fatal(err)
			}
			tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
			if fileExists(tempUnixProfilePath) {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.zprofile")
			}
			tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.bash_profile")
			if fileExists(tempUnixProfilePath) {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.bash_profile")
			}
			tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.bashrc")
			if fileExists(tempUnixProfilePath) {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.bashrc")
			}
			tempUnixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
			if fileExists(tempUnixProfilePath) {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
			}
			if !fileExists(tempUnixProfilePath) {
				unixProfilePath = fmt.Sprint(userDirectory() + "/.profile")
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

// check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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
