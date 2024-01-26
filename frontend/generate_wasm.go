//go:build ignore
// +build ignore

package main

import (
	"os"
	"os/exec"
)

func main() {
	// generate the app.wasm file.
	// execute the command: GOARCH=wasm GOOS=js go build . -o ./web/app.wasm
	// in the current directory.

	env := []string{
		"GOARCH=wasm",
		"GOOS=js",
	}

	args := []string{
		"go",
		"build",
		"-o",
		"web/app.wasm",
		".",
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
