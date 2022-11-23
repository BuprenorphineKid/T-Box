package main

import (
	"boxpkg/box"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	h, _ := strconv.ParseInt(args[0], 0, 0)
	w, _ := strconv.ParseInt(args[1], 0, 0)

	b := box.NewBox(int(h), int(w), args[2])

	x, err := strconv.ParseInt(args[3], 0, 0)
	if err != nil {
		x := 0
		b.OffsetX(x)
	}

	b.OffsetX(int(x))

	y, err := strconv.ParseInt(args[4], 0, 0)
	if err != nil {
		y := 0
		b.OffsetY(y)
	}

	b.OffsetY(int(y))

	b.Show()
}
