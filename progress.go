package animation

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
)

type Animation struct {
	Frame    []string
	Width    int
	Interval time.Duration
}

func (this Animation) Progress(out io.Writer) func() {
	done := make(chan struct{})
	go func() {
		w := this.Width
		if w <= 0 {
			w = runewidth.StringWidth(this.Frame[0])
		}
		backspace := strings.Repeat("\b", w)
		erase := strings.Repeat(" ", w)

		io.WriteString(out, this.Frame[0])

		interval := this.Interval
		if interval <= 0 {
			interval = time.Second / 10
		}
		ticker := time.NewTicker(interval)
		i := 1
		for {
			select {
			case <-done:
				ticker.Stop()
				close(done)
				io.WriteString(out, backspace)
				io.WriteString(out, erase)
				io.WriteString(out, backspace)
				return
			case <-ticker.C:
				if i >= len(this.Frame) {
					i = 0
				}
				io.WriteString(out, backspace)
				io.WriteString(out, this.Frame[i])
				i++
			}
		}
	}()

	return func() {
		done <- struct{}{}
	}
}

func Progress() func() {
	return Animation{
		Frame: []string{" /", " -", " \u2216", " |"},
		Width: 2,
	}.Progress(os.Stdout)
}
