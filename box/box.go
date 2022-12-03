package box

import (
	"strconv"
	"strings"
)

type signal struct{}

type Box struct {
	Height int
	Width  int
	chars  []string
	lines  []string
	shape  string
	style  string
	Area   int
}

func SaveCursor() {
  print("\0337")
}


func RestoreCursor() {
  print("\0338")
}

func NewBox(h int, w int, s string) *Box {
	var b Box
	b.Height = h
	b.Width = w - 2

	b.Area = (b.Height * b.Width)

	b.chars = make([]string, 0, b.Area)

	b.style = s
	b.Build()

	return &b
}

func (b *Box) OffsetX(x int) {
	b.lines = strings.Split(b.shape, "\n")
	for i, _ := range b.lines {
		for k := 0; k <= x; k++ {
			b.lines[i] = " " + b.lines[i]
		}
	}
}

func (b *Box) OffsetY(y int) {
	print("\033[" + strconv.FormatInt(int64(y), 10) + ";0H")
	
}

func (b *Box) Build() {
	switch b.style {
	case "thin":
		b.thin()
	case "thick":
		b.thick()
	case "chain":
		b.chain()
  case "sharp":
    b.sharp()
  default:
    b.thin()
	}
	b.shape = strings.Join(b.chars, "")
}

func (b Box) Show() {
	b.shape = strings.Join(b.lines, "\n")
	println(b.shape)
}
