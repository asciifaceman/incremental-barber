package main

import (
	"fmt"

	"github.com/asciifaceman/incremental-barber/pkg/core"
)

var version = "localdev"

func main() {
	c := core.New(version)
	fmt.Printf("Launching Incremental Barber %s\n", c.Version)
	c.Init()
}
