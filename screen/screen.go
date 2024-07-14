package screen

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
	"golang.org/x/image/draw"
	terminal "golang.org/x/term"
)

var scale = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,^`'. "

type Screen struct {
	Window *Block
	Blocks []PrintInterface
	UnimplementedPrintInterface
	Running bool
	Input   bool
}

var Scr *Screen

const (
	HORIZONTAL int = 0
	VERTICAL       = 1
	FLEX           = 2
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257)}
}

func (s *Screen) GetDim() (int, int) {
	return s.Window.GetDim()
}

func (s *Screen) PrintImage(file_name string, ext string) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	dest := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/5, img.Bounds().Max.Y/5))
	draw.NearestNeighbor.Scale(dest, dest.Rect, img, img.Bounds(), draw.Over, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	bounds := dest.Bounds()

	// fmt.Println(img.Bounds())

	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel

	pixels = make([][]Pixel, 0)

	for y := 0; y < height; y++ {
		var row []Pixel
		row = make([]Pixel, 0)
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(dest.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	for _, row := range pixels {
		for _, p := range row {
			avg := (float64(p.R) + float64(p.G) + float64(p.B)) / 3
			charIndex := int(float64(avg) / 255 * float64(len(scale)-1))
			fmt.Printf("%c", scale[charIndex])
		}
		fmt.Printf("\n")
	}
}

func (s *Screen) Reload() {
	width, height, err := terminal.GetSize(0)

	if err != nil {
		log.Fatal(err)
		return
	}

	s.Window.Width = width
	s.Window.Height = height

	for _, bl := range s.Window.Data {
		bl.Reload(s.Window)
	}
}

func (s *Screen) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (s *Screen) Refresh() {
	s.Reload()
	s.Clear()
	fmt.Printf("The size of the elements on screen are %d.\n", len(s.Window.Data))
	fmt.Printf("%s", s.Print())
}

func (s *Screen) Print() string {
	return s.Window.Print(s.Window)
}

func (s *Screen) AddBlock(b PrintInterface) {
	s.Window.AddBlock(b)
}

func (s *Screen) RemoveBlock(b PrintInterface) {
	for k, i := range s.Window.Data {
		if i == b {
			s.Window.Data = append(s.Window.Data[:k], s.Window.Data[k+1:]...)
		}
	}
}

func NewScreen() *Screen {

	w, h, e := terminal.GetSize(0)

	if e != nil {
		log.Fatal(e)
		return nil
	}

	Screen := &Screen{
		Window:  NewBlock(w, h, make([]PrintInterface, 0)),
		Running: true,
	}

	return Screen
}

func (s *Screen) Keys() {
	keysEvent, err := keyboard.GetKeys(10)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for s.Running {

		event := <-keysEvent

		if event.Err != nil {
			log.Fatal(event.Err)
			continue
		}

		switch event.Key {
		case keyboard.KeyEsc:
			s.Running = false
			break
		case keyboard.KeyEnter:
			fmt.Printf("Making choice")
			s.Window.MakeChoice()
			break
		default:
			if s.Input {
				s.Window.AddChar(event.Rune)
			} else {
				switch event.Rune {
				case int32('w'):
					fmt.Printf("It's a W\n")
					s.Window.DownChoice()
					break
				case int32('s'):
					fmt.Printf("It's an S\n")
					s.Window.UpChoice()
					break
				}
			}

		}

		s.Refresh()
	}
}
