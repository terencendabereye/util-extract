/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
