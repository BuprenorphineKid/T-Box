package box

import ()

func (b *Box) thin() {
	b.borders["top"] = "–"
	b.borders["left"] = "|"
	b.borders["right"] = "|"
	b.borders["bottom"] = "–"
}

func (b *Box) thick() {
	b.borders["top"] = "Ξ"
	b.borders["left"] = "│|"
	b.borders["right"] = "|│"
	b.borders["bottom"] = "Ξ"
}

func (b *Box) chain() {
	b.borders["top"] = "┼"
	b.borders["left"] = "├│┤"
	b.borders["right"] = "├│┤"
	b.borders["bottom"] = "┼"
}

func (b *Box) sharp() {
	b.borders["top"] = "⇌"
	b.borders["left"] = "⁺┼₊"
	b.borders["right"] = "⁺┼₊"
	b.borders["bottom"] = "⇌"
}
