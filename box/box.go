package box

import(
  "strings"
  "strconv"
)

type signal struct{}

type Box struct {
  Height int
  Width int
  chars []string
  lines []string
  shape string
  style string
  Area int
} 

func NewBox(h int, w int, s string) *Box {
  var b Box
  b.Height = h
  b.Width = w

  b.Area = (h * w)
    
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

func (b *Box) thin() {
  for i := 0; i <= b.Height; i++ {
    b.chars = append(b.chars, "|")
    
    if i == 0 {

      for j := 0; j <= b.Width; j++ {
        b.chars = append(b.chars,"⁻")
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
        b.chars = append(b.chars,"Ξ")
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
        b.chars = append(b.chars,"┼")
      } 
        b.chars = append(b.chars, "├│┤")
        b.chars = append(b.chars, "\n")
    } else if i == b.Height {
      for r := 0; r <= b.Width; r++  {
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

func (b *Box) Build() {
  switch {
  case b.style == "thin":
    b.thin()
  case b.style == "thick":
    b.thick()
  case b.style == "chain":
    b.chain()
  }
  b.shape = strings.Join(b.chars, "")
}

func (b Box) Show() {
  b.shape = strings.Join(b.lines, "\n")
  println(b.shape)
}
