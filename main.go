package main

import (
	"extract/cmd"

	"os"
)

// import "cmd"
var (
	version string
)

func main() {
	os.Setenv("VERSION_NUMBER", version)
	cmd.Extract()
	// fmt.Printf("version=%s build=%s\n", version, date)
}