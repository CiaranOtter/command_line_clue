package screen

import (
	"fmt"
	"strings"
)

type Heading struct {
	UnimplementedPrintInterface
	Title     string
	BorderChr string
}

func (h *Heading) Print(b PrintInterface) string {

	heading := b.PrintPadding("Left")

	w, _ := b.GetDim()

	if strings.Compare(h.Title, "") != 0 {
		titLen := len(h.Title)

		side := (w - titLen - 2) / 2

		i := 0

		for i < side {
			heading = fmt.Sprintf("%s%s", heading, h.BorderChr)
			i++
		}

		heading = fmt.Sprintf("%s %s ", heading, h.Title)
		i += 2 + len(h.Title)

		for i < w {
			heading = fmt.Sprintf("%s%s", heading, h.BorderChr)
			i++
		}
	} else {
		for i := 0; i < w; i++ {
			heading = fmt.Sprintf("%s%s", heading, h.BorderChr)
		}
	}

	heading = fmt.Sprintf("%s%s\n", heading, b.PrintPadding("Right"))

	return heading
}

func NewHeading(title string, bc string) *Heading {
	return &Heading{
		Title:     title,
		BorderChr: bc,
	}
}
