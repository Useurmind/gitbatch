package main

import (
	"os"

	"github.com/Useurmind/gitbatch/cmd"
)
  
func main() {
	cmd.Init()

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
