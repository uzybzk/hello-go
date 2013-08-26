package main

import "fmt"

func variables() {
    var name string = "Go"
    var version float32 = 1.1
    
    fmt.Printf("Language: %s\n", name)
    fmt.Printf("Version: %.1f\n", version)
}