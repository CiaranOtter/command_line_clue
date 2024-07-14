package screen

import "fmt"

type List struct {
	UnimplementedPrintInterface
	Items []string
}

func (cl *List) AddChoice(s string) {
	cl.Items = append(cl.Items, s)
	fmt.Print("Added choice\n")
}

func (l *List) Print(b PrintInterface) string {
	s := ""

	for _, key := range l.Items {
		s = fmt.Sprintf("%s%s\n", s, key)
	}

	return s
}

func NewList(items []string) *List {

	if items == nil {
		items = make([]string, 0)
	}
	return &List{
		Items: items,
	}
}
