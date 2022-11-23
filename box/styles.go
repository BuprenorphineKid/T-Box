package box

import ()

func (b *Box) thin() {
	for i := 0; i <= b.Height; i++ {
		b.chars = append(b.chars, "|")

		if i == 0 {

			for j := 0; j <= b.Width; j++ {
				b.chars = append(b.chars, "⁻")
			}
			b.chars = append(b.chars, "|")
			b.chars = append(b.chars, "\n")
		} else if i == b.Height {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, "_")
			}
			b.chars = append(b.chars, "|")
			b.chars = append(b.chars, "\n")
		} else {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, " ")
			}
			b.chars = append(b.chars, "|")
			b.chars = append(b.chars, "\n")
		}
	}
}

func (b *Box) thick() {
	for i := 0; i <= b.Height; i++ {
		b.chars = append(b.chars, "|│")

		if i == 0 {

			for j := 0; j <= b.Width; j++ {
				b.chars = append(b.chars, "Ξ")
			}
			b.chars = append(b.chars, "│|")
			b.chars = append(b.chars, "\n")
		} else if i == b.Height {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, "Ξ")
			}
			b.chars = append(b.chars, "│|")
			b.chars = append(b.chars, "\n")
		} else {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, " ")
			}
			b.chars = append(b.chars, "│|")
			b.chars = append(b.chars, "\n")
		}
	}
}

func (b *Box) chain() {
	for i := 0; i <= b.Height; i++ {
		b.chars = append(b.chars, "├│┤")

		if i == 0 {

			for j := 0; j <= b.Width; j++ {
				b.chars = append(b.chars, "┼")
			}
			b.chars = append(b.chars, "├│┤")
			b.chars = append(b.chars, "\n")
		} else if i == b.Height {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, "┼")
			}
			b.chars = append(b.chars, "├│┤")
			b.chars = append(b.chars, "\n")
		} else {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, " ")
			}
			b.chars = append(b.chars, "├│┤")
			b.chars = append(b.chars, "\n")
		}
	}
}

func (b *Box) sharp() {
	for i := 0; i <= b.Height; i++ {
		b.chars = append(b.chars, "⁺┼₊")

		if i == 0 {     

			for j := 0; j <= b.Width; j++ {
				b.chars = append(b.chars, "⇌")
			}
			b.chars = append(b.chars, "⁺┼₊")
			b.chars = append(b.chars, "\n")
		} else if i == b.Height {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, "⇌")
			}
			b.chars = append(b.chars, "⁺┼₊")
			b.chars = append(b.chars, "\n")
		} else {
			for r := 0; r <= b.Width; r++ {
				b.chars = append(b.chars, " ")
			}
			b.chars = append(b.chars, "⁺┼₊")
			b.chars = append(b.chars, "\n")
		}
	}
}
