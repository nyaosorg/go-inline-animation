//go:build run

package main

import (
	"fmt"
	"time"

	"github.com/mattn/go-colorable"

	"github.com/nyaosorg/go-inline-animation"
)

func main() {
	term := colorable.NewColorableStdout()
	fmt.Fprint(term, animation.CursorOff)
	defer fmt.Fprintln(term, animation.CursorOn)

	end := animation.Dots.Progress(term)
	defer end()

	// You can write a time-consuming process instead of Sleep.
	time.Sleep(time.Second * 10)
}
