package box

import "fmt"

type Box struct {
	OffsetX int
	OffsetY int
	Height  int
	Width   int
	borders map[string]string
}

func NewBox(h int, w int) *Box {
	var b Box
	b.Height = h
	b.Width = w - 2

	b.borders = make(map[string]string)

	return &b
}

func (b *Box) Stylize(style string) {
	switch style {
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
}

func (b Box) Show(t Terminal) {
	if b.OffsetX >= t.Cols {
		println("\nError :\n\n\t Impossible X point of current terminal size")
	}

	if b.OffsetY >= t.Rows {
		println("\nError :\n\n\t Impossible X point of current terminal size")
	}

	MoveCursor(b.OffsetX, b.OffsetY)

	for i := 0; i <= b.Width; i++ {
		fmt.Printf("%s", b.borders["top"])
	}

	MoveCursor(b.OffsetX+1, b.OffsetY-1+b.Height)

	for i := 0; i < b.Width; i++ {
		fmt.Printf("%s", b.borders["bottom"])
	}

	for i := 0; i < b.Height; i++ {
		MoveCursor(b.OffsetX, b.OffsetY+i)

		fmt.Printf("%s", b.borders["left"])
	}

	for i := 0; i < b.Height; i++ {
		MoveCursor(b.OffsetX+b.Width, b.OffsetY+i)

		fmt.Printf("%s", b.borders["right"])
	}
}
