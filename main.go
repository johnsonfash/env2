package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: env2 <action>")
        return
    }

    action := os.Args[1]
    switch action {
    case "say-hello":
        fmt.Println("Hello from env2!")
    default:
        fmt.Println("Unknown action:", action)
    }
}
