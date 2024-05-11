package main

import (
	"fmt"
	"runtime"

	"github.com/hubschluft/hubschluft.github.io/cmd"
)

func main() {
	os_slice := []string{"linux", "openbsd", "freebsd", "netbsd", "dragonfly", "android", "darwin"}

	os_type := false

	for _, str := range os_slice {
		if str == runtime.GOOS {
			os_type = true
			break
		}
	}

	if os_type == true {
		cmd.Arguments()
	} else {
		fmt.Printf("You are not using: %s\n", os_slice)
	}
}
