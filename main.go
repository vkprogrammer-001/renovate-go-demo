package main

import (
    "fmt"
	"io"
    "os"
    "strings"

    "golang.org/x/text/encoding/unicode"
    "golang.org/x/text/transform"
)

func main() {
    fmt.Println("Simple Text Transformer")
    
    // Create a vulnerable string with potential UTF-8 security issues
    input := "Hello, 世界!" // "Hello, World!" with some non-ASCII characters
    
    // Use the vulnerable dependency
    encoder := unicode.UTF8.NewEncoder()
    transformer := transform.NewReader(strings.NewReader(input), encoder)
    
    // Read and print the transformed content
    result, err := io.ReadAll(transformer)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    
    fmt.Println("Transformed text:", string(result))
}