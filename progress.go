package animation

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
)

const (
	CursorOff = "\x1B[?25l"
	CursorOn  = "\x1B[?25h"
)

type Animation struct {
	Frame    []string
	Width    int
	Interval time.Duration
}

var Dots = Animation{
	Frame: []string{
		"⠋", //	U+280B	Braille Pattern Dots-124
		"⠙", //	U+2819	Braille Pattern Dots-145
		"⠹", //	U+2839	Braille Pattern Dots-1456
		"⠸", //	U+2838	Braille Pattern Dots-456
		"⠼", //	U+283C	Braille Pattern Dots-3456
		"⠴", //	U+2834	Braille Pattern Dots-356
		"⠦", //	U+2826	Braille Pattern Dots-236
		"⠧", //	U+2827	Braille Pattern Dots-1236
		"⠇", //	U+2807	Braille Pattern Dots-123
		"⠏", //	U+280F	Braille Pattern Dots-1234
	},
}

var Bars = Animation{
	Frame:    []string{" /", " -", " \u2216", " |"},
	Interval: time.Second / 2,
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
	return Dots.Progress(os.Stdout)
}
