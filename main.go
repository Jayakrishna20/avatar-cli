package main

import "avtr/cmd"

var version = "dev" // This gets seamlessly overwritten by GoReleaser during the build

func main() {
	cmd.Version = version
	cmd.Execute()
}
