package main

import (
	"fmt"
	"runtime"

	"github.com/TChi91/coretabs/cmd"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "linux":
		cmd.Execute()
	default:
		fmt.Print("We support only linux OS\nTry later")
	}

}
