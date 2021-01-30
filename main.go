package main

import (
    "fmt"
    "runtime"
)

func main() {
    os := runtime.GOOS
    switch os {
    case "windows":
        fmt.Println("Windows")
    case "darwin":
        fmt.Println("MAC operating system")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s.\n", os)
    }
}
