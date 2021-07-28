package main

import (
	"os"

	"github.com/rajasoun/tdd-go/cmd"
)

func main() {
	//Set Output to Standard Output
	cmd.NewRootCmd().SetOut(os.Stdout)
	cmd.Execute()
}
