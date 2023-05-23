package main

import (
	"os"
	"strconv"
	"tbox/box"
)

var term box.Terminal = *box.NewTerminal()

func main() {
	box.SaveCursor()

	args := os.Args[1:]

	if len(args) < 2 {
		box.Usage()
		os.Exit(0)
	}

	h, _ := strconv.ParseInt(args[0], 0, 0)
	w, _ := strconv.ParseInt(args[1], 0, 0)

	b := box.NewBox(int(h), int(w))

	switch len(args[2:]) {
	case 0:
		b.OffsetX = 1
		b.OffsetY = 1

		b.Stylize("thin")
	case 1:
		x, err := strconv.ParseInt(args[2], 0, 0)

		if err != nil {
			x := 1
			b.OffsetX = x
		}

		b.OffsetX = int(x)
		b.OffsetY = 1

		b.Stylize("thin")
	case 2:
		x, err := strconv.ParseInt(args[2], 0, 0)

		if err != nil {
			x := 1
			b.OffsetX = x
		}

		b.OffsetX = int(x)

		y, err := strconv.ParseInt(args[3], 0, 0)
		if err != nil {
			y := 1
			b.OffsetY = y
		}

		b.OffsetY = int(y)

		b.Stylize("thin")
	case 3:
		x, err := strconv.ParseInt(args[2], 0, 0)

		if err != nil {
			x := 1
			b.OffsetX = x
		}

		if int(x) <= 0 {
			x = 1
		}

		b.OffsetX = int(x)

		y, err := strconv.ParseInt(args[3], 0, 0)
		if err != nil {
			y := 1
			b.OffsetY = y
		}

		if int(y) <= 0 {
			y = 1
		}

		b.OffsetY = int(y)

		b.Stylize(args[4])
	}

	b.Show(term)

	box.RestoreCursor()
}
