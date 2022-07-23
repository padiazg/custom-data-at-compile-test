package main

import "fmt"

var (
	buildTime  string
	commitHash string
)

func main() {
	fmt.Printf("Build Time: %s\n", buildTime)
	fmt.Printf("Commit Hash: %s\n", commitHash)
}
