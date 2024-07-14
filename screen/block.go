package screen

import "fmt"

type Block struct {
	Width   int
	Height  int
	Padding map[string]int
	Data    []PrintInterface
	Parent  *Block
	Hidden  bool
	*UnimplementedPrintInterface
}

func (b *Block) RemoveFirst() {
	b.Data = b.Data[1:]
}

func (b *Block) RemoveLast() {
	b.Data = b.Data[:len(b.Data)]
}

func (b *Block) PutFirst(p PrintInterface) {
	b.Data = append([]PrintInterface{p}, b.Data...)
}

func (b *Block) Show() {
	b.Hidden = false
}

func (b *Block) Hide() {
	b.Hidden = true
}

func NewBlock(w int, h int, d []PrintInterface) *Block {
	b := &Block{
		Width:  w,
		Height: h,
		Padding: map[string]int{
			"Left":   0,
			"Right":  0,
			"Top":    0,
			"Bottom": 0,
		},
		Data:   d,
		Hidden: false,
	}

	return b
}

func (b *Block) RemoveBlock(f PrintInterface) {
	for i, k := range b.GetData() {
		if f == k {
			b.Data = append(b.Data[:i], b.Data[i+1:]...)
			return
		}
	}
}

func (b *Block) AddChar(letter int32) {
	for _, i := range b.GetData() {
		i.AddChar(letter)
	}
}

func (b *Block) UpChoice() {
	for _, i := range b.GetData() {
		i.UpChoice()
	}
}

func (b *Block) DownChoice() {
	for _, i := range b.GetData() {
		i.DownChoice()
	}
}

func (b *Block) MakeChoice() {
	fmt.Printf("Making choice\n")
	for _, i := range b.GetData() {
		fmt.Printf("Making the choice deeper\n")
		i.MakeChoice()
	}
}

func (b *Block) PrintPadding(s string) string {
	i, e := b.Padding[s]

	if !e {
		return ""
	}

	q := ""

	for j := 0; j < i; j++ {
		q = fmt.Sprintf("%s ", q)
	}

	return q
}

func (b *Block) GetPadding() map[string]int {
	return b.Padding
}

func (b *Block) GetDim() (int, int) {
	return b.Width, b.Height
}

func (b *Block) SetPadding(p map[string]int) {
	for key, pad := range b.Padding {
		s, e := p[key]

		if e {
			b.Padding[key] = s
		} else {
			b.Padding[key] = pad
		}
	}
}

func (b *Block) Print(p PrintInterface) string {

	if b.Hidden {
		return ""
	}
	s := ""

	for _, bl := range b.Data {
		s = fmt.Sprintf("%s%s%s%s", s, p.PrintPadding("Left"), bl.Print(b), p.PrintPadding("Right"))
	}

	return s
}

func (b *Block) GetData() []PrintInterface {
	return b.Data
}

func (b *Block) AddBlock(p PrintInterface) {
	b.Data = append(b.Data, p)
}

func (b *Block) Reload(c PrintInterface) {
	p := c.GetPadding()
	w, h := c.GetDim()
	b.Width = w - p["Left"] - p["Right"]
	b.Height = h - p["Top"] - p["Bottom"]

	for _, bl := range c.GetData() {
		bl.Reload(b)
	}

}

func (b *Block) Clear() {
	b.Data = make([]PrintInterface, 0)
}
