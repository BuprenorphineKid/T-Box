package box

import (
	"fmt"
	"os"

	term "golang.org/x/crypto/ssh/terminal"
)

const ESC = "\033"

type Terminal struct {
	Cols, Rows int
}

func NewTerminal() *Terminal {
	var t Terminal

	var err error

	t.Cols, t.Rows, err = term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Printf("Error: Couldnt read size of terminal")
		os.Exit(1)
	}

	return &t
}

func SaveCursor() {
	print(ESC + "7")
}

func RestoreCursor() {
	print(ESC + "8")
}

func MoveCursor(x int, y int) {
	fmt.Printf("%s[%d;%dH", ESC, y, x)
}