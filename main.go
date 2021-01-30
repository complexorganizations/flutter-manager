package main

import (
    "fmt"
    "runtime"
)

func installFlutterOnWindows() {
    os.Mkdir("/src/", 0755)
    //os.Chdir("/src")
}

func selectOperatingSystem() {
    os := runtime.GOOS
    switch os {
    case "windows":
        fmt.Println("Windows")
    case "darwin":
        fmt.Println("MAC operating system")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("Error: %s Not Supported.\n", os)
    }
}

func main() {
    selectOperatingSystem()
}
