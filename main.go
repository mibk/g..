package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kisielk/gotool"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	var ipaths []string
	var i int
	for i = 0; i < len(args); i++ {
		// Is it a path?
		if !strings.ContainsAny(args[i], "/\\") {
			break
		}
	}
	if i == 0 {
		ipaths = []string{"./..."}
	} else {
		ipaths = args[:i]
		args = args[i:]
	}

	paths := gotool.ImportPaths(ipaths)
	var filtered []string
	for _, p := range paths {
		if !strings.Contains(p, "/vendor/") {
			filtered = append(filtered, p)
		}
	}

	if len(args) == 0 {
		// Only print and exit.
		for _, p := range filtered {
			fmt.Println(p)
		}
		return
	}

	for _, p := range filtered {
		args = append(args, p)
	}
	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: g.. [path ...] [go sub-command [arg ...]]

g.. executes a go sub-command (e.g. test, vet, fmt, ...) for the given paths
excluding vendored packages. If no paths are provided, g.. uses ./... as a
default path. If no sub-command is provided, g.. just lists the packages.`)
}
