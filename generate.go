// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/leoh0/binctr/container"
)

// Pulls an image and saves the binary data in the container package bindata.go.
func main() {
	if err := container.EmbedImage("docker.io/leoh0/kakaotalk-wine-root"); err != nil {
		fmt.Fprintf(os.Stderr, "embed image failed: %v\n", err)
		os.Exit(1)
	}
}
