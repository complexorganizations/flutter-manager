package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	selectOperatingSystem()
}

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

func isNotExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return true
	}
	return !info.IsDir()
}
