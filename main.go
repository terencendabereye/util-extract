/*
Copyright © 2023 Terence Ndabereye ndabereye@gmail.com
*/
package main

import (
	"extract/cmd"
)

var (
	version string
)
func init() {
	cmd.AppVersion = version
}
func main() {
	cmd.Execute()
}
