package main

import (
	"fmt"
	"runtime"

	"github.com/TChi91/coretabs/cmd"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		cmd.Execute()
	default:
		fmt.Print("We currently support only Windows OS\nTry later")
	}

}
