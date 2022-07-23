package main

import "fmt"

func main() {
	fmt.Printf("Build Time: %s\n", GetBuildTime())
	fmt.Printf("Commit Hash: %s\n", GetCommitHash())
	fmt.Printf("Version: %s\n", GetVersion())
	fmt.Printf("Formatted: %s\t SHA: %s\t Build date: %s \n", GetVersion(), GetCommitHash(), GetBuildTime())

}
