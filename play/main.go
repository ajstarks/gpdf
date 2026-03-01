package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/ajstarks/gpdf"
)

const (
	mm2pt        = 2.83464 // mm to pt conversion
	letterWidth  = 792
	letterHeight = 612
)

// PageDimen describes page dimensions
// the unit field is used to convert to pt.
type PageDimen struct {
	width  float64
	height float64
	unit   float64
}

// command line options
type options struct {
	sansfont   string
	serifont   string
	monofont   string
	symbolfont string
	pagesize   string
	output     string
}

// convert tabs to spaces
var codemap = strings.NewReplacer("\t", "    ")

// fontmap maps generic font names to specific implementation names
var fontmap = map[string]gpdf.CF{}

// pagemap defines page dimensions
var pagemap = map[string]PageDimen{
	"Letter":     {792, 612, 1},
	"Legal":      {1008, 612, 1},
	"Tabloid":    {1224, 792, 1},
	"ArchA":      {864, 648, 1},
	"Widescreen": {1152, 648, 1},
	"4R":         {432, 288, 1},
	"Index":      {360, 216, 1},
	"A2":         {420, 594, mm2pt},
	"A3":         {420, 297, mm2pt},
	"A4":         {297, 210, mm2pt},
	"A5":         {210, 148, mm2pt},
}

func pagedim(s string) (float64, float64) {
	// try lookup first...
	v, ok := pagemap[s]
	if ok {
		return v.width, v.height
	}
	// lookup fails, try WxH
	fields := strings.Split(s, "x")
	if len(fields) != 2 {
		return letterWidth, letterHeight
	}
	w, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return letterWidth, letterHeight
	}
	h, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return letterWidth, letterHeight
	}
	return w, h
}

func lf(c *gpdf.Canvas, name, file string) {
	f, err := c.LoadFontFile(file)
	if err != nil {
		fontmap[name] = nil
	} else {
		fontmap[name] = f
	}
}

// TextCode shows a text file, upper left at (x,y), dimensions (w, h) at size
// the border of the block background and textcolors are also specified.
func TextCode(c *gpdf.Canvas, name string, x, y, w, h, size, ls, border float64, bgcolor, txcolor string) {
	cf := c.CustomFont
	if c.CustomFont != nil {
		cf = c.CustomFont
	}
	c.CornerRect(x-border, y+border, w+(border*2), h+(border*2), txcolor)
	c.CornerRect(x, y, w, h, bgcolor)
	r, err := os.Open(name)
	if err != nil {
		c.Text(x+(w/2), y-(h/2), w/20, "File not found", "red")
		return
	}
	scanner := bufio.NewScanner(r)
	y -= size
	for scanner.Scan() {
		c.Text(x+(size/2), y-(size/2), size, codemap.Replace(scanner.Text()), txcolor)
		y -= (size * ls)
	}
	c.CustomFont = cf
}

func main() {
	var opts options
	flag.StringVar(&opts.pagesize, "pagesize", "Letter", "page size (name or WxH")
	flag.StringVar(&opts.sansfont, "sans", "sans.ttf", "sans font")
	flag.StringVar(&opts.serifont, "serif", "serif.ttf", "serif font")
	flag.StringVar(&opts.monofont, "mono", "mono.ttf", "monofont font")
	flag.StringVar(&opts.symbolfont, "symbol", "symbol.ttf", "default font")
	flag.StringVar(&opts.output, "o", "f.pdf", "output file")

	flag.Parse()

	output := opts.output
	args := flag.Args()
	if len(args) > 0 {
		output = args[0]
	}
	cw, ch := pagedim(opts.pagesize)
	canvas, err := gpdf.SetupCanvas(cw, ch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	lf(canvas, "sans", opts.sansfont)
	lf(canvas, "serif", opts.serifont)
	lf(canvas, "mono", opts.monofont)
	lf(canvas, "symbol", opts.symbolfont)

	canvas.SetFont(fontmap["sans"])

	y := 85.0
	ts := 1.8
	bgcolor := "black"
	canvas.Background(bgcolor)

	// usage
	canvas.Background(bgcolor)
	for range 500 {
		xr := rand.Float64() * 100
		yr := rand.Float64() * 100
		canvas.Circle(xr, yr, 0.25, "white")
	}
	canvas.Circle(50, 0, 50, "blue")
	canvas.CText(50, 20, 10, "hello, world", "white")
	codesize := 1.0
	border := codesize / 8
	canvas.SetFont(fontmap["mono"])
	TextCode(canvas, "excode", 2, 98, 40, 45, codesize, 1.5, border, "black", "white")

	canvas.NewPage(cw, ch)
	canvas.GradRect(0, 0, 100, 100, "white", "blue", 10)
	canvas.SetFont(fontmap["mono"])
	letters := "ABCEDEFGHIJKLMNOPQRSTUVWXYZ"
	x := 2.5
	y = 95.0
	ts = 2.0
	ls := ts * 1.6
	for _, r := range letters {
		canvas.Text(x, y, ts, string(r), "maroon")
		y -= ls
	}
	canvas.SetFont(fontmap["symbol"])
	x += 2.5
	y = 95.0
	for _, r := range letters {
		canvas.Text(x, y, ts, string(r), "black")
		y -= ls
	}
	palette := []string{"black", "red", "green", "blue", "orange"}
	a := 0.0
	n := 10
	op := 100.0
	canvas.Circle(50, 50, 2, "red")
	for i := range n {
		color := palette[i%(len(palette))]
		canvas.RText(50, 50, 40, a, "A", color, op)
		a += 360 / float64(n)
		op -= 10
	}

	canvas.NewPage(cw, ch)
	asize := 0.2
	aw := 10.0
	ah := 5.0
	canvas.Circle(50, 50, 0.5, "maroon")
	canvas.Arc(50, 50, aw, ah, 0, 90, asize, "red")
	canvas.Arc(50, 50, aw, ah, 90, 180, asize, "green")
	canvas.Arc(50, 50, aw, ah, 180, 270, asize, "blue")
	canvas.Arc(50, 50, aw, ah, 270, 360, asize, "orange")
	op = 100
	for r := 15.0; r <= 50; r += 5 {
		canvas.Arc(50, 50, r, r, 0, 90, asize*2, "red", op)
		canvas.Arc(50, 50, r, r, 90, 180, asize*2, "green", op)
		canvas.Arc(50, 50, r, r, 180, 270, asize*2, "blue", op)
		canvas.Arc(50, 50, r, r, 270, 360, asize*2, "orange", op)
		op -= 5

	}
	canvas.LGrid(0, 100, 0, 100, 0.1, 5, "gray", 70)
	err = canvas.Creator.WriteToFile(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
