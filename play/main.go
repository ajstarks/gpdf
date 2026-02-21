package main

import (
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

// fontmap maps generic font names to specific implementation names
var fontmap = map[string]string{}

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

func main() {
	var opts options
	flag.StringVar(&opts.pagesize, "pagesize", "Letter", "page size (name or WxH")
	flag.StringVar(&opts.sansfont, "sans", "sans.ttf", "sans font")
	flag.StringVar(&opts.serifont, "serif", "serif.ttf", "serif font")
	flag.StringVar(&opts.monofont, "mono", "mono.ttf", "monofont font")
	flag.StringVar(&opts.symbolfont, "symbol", "symbol.ttf", "default font")
	flag.StringVar(&opts.output, "o", "f.pdf", "output file")
	flag.Parse()

	fontmap["mono"] = opts.monofont
	fontmap["sans"] = opts.sansfont
	fontmap["serif"] = opts.serifont
	fontmap["symbol"] = opts.symbolfont
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
	err = canvas.LoadFont(fontmap["sans"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// API
	labels := []string{
		"Arc(x,y, r,a1,a2, size,color)",
		"Circle(x,y, r, color)",
		"Ellipse(x,y, w,h, color)",
		"Image(x,y, w,h, name)",
		"Line(x1,y1, x2,y2, size, color)",
		"Polygon(x,y, color)",
		"QuadCurve(bx,by, cx,cy, ex,ey, size, color)",
		"Rect(x,y, w,h, color)",
		"{B,C,E}Text(x,y, size, color),   RText(x,y, size,angle, color)",
		"TextCode(name, x,y,w,h, size,border, bcolor,tcolor)",
		"Grid(x1,x2, y1,y2, size, incr, color)",
		"Background(color)",
	}
	yspace := 7.5
	yspace2 := yspace / 2
	y := 85.0
	lw := 0.3
	c1 := 5.0
	c2 := c1 + 85
	c3 := c2 - 5
	c4 := c2 + 5
	ts := 2.5
	dotsize := 0.4
	shapecolor := "rgb(200,200,200)" //"lightgray"
	dotcolor := "white"
	tcolor := "gray"
	canvas.Background("black")
	canvas.Text(c1, 95, 3, "Generate PDF (gpdf) API", "white")
	canvas.StdFont = gpdf.Mono
	for i := range labels {
		canvas.Text(c1, y, ts, labels[i], tcolor)
		y -= yspace
	}
	y = 85.0
	// arc
	canvas.Circle(c2, y, dotsize, dotcolor)
	canvas.Arc(c2, y, 5, 0, 90, lw, shapecolor)
	y -= yspace
	// circle
	canvas.Circle(c2, y, 3, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// ellipse
	canvas.Ellipse(c2, y, 5, 2, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// image
	canvas.Image(c2, y, 64, 48, "follow.jpg")
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// line
	canvas.Line(c3, y, c4, y, lw, shapecolor)
	canvas.Circle(c3, y, dotsize, dotcolor)
	canvas.Circle(c4, y, dotsize, dotcolor)
	y -= yspace
	// polygon
	xp := []float64{c3, c2, c4}
	yp := []float64{y, y + yspace2, y}
	canvas.Circle(xp[0], yp[0], dotsize, dotcolor)
	canvas.Circle(xp[1], yp[1], dotsize, dotcolor)
	canvas.Circle(xp[2], yp[2], dotsize, dotcolor)
	canvas.Polygon(xp, yp, shapecolor)
	y -= yspace
	// quadcurve
	bx := c3
	by := y
	cx := c2 + 5
	cy := y + yspace2
	ex := c4
	ey := y
	canvas.QuadCurve(bx, by, cx, cy, ex, ey, lw, shapecolor)
	canvas.Circle(bx, by, dotsize, dotcolor)
	canvas.Circle(cx, cy, dotsize, dotcolor)
	canvas.Circle(ex, ey, dotsize, dotcolor)
	y -= yspace
	// rect
	canvas.Rect(c2, y, 4, 5, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// text
	tc0 := c2 - 16
	tc1 := c2 - 8
	tc2 := c2
	tc3 := c2 + 3
	canvas.BText(tc0, y, 2, "hello", shapecolor)
	canvas.Circle(tc0, y, dotsize, dotcolor)
	canvas.CText(tc1, y, 2, "hello", shapecolor)
	canvas.Circle(tc1, y, dotsize, dotcolor)
	canvas.EText(tc2, y, 2, "hello", shapecolor)
	canvas.Circle(tc2, y, dotsize, dotcolor)
	canvas.RText(tc3, y, 2, 45, "hello", shapecolor)
	canvas.Circle(tc3, y, dotsize, dotcolor)
	//canvas.Etext()
	y -= yspace
	// textcode
	canvas.TextCode("hello.txt", c3, y+4, 12, yspace*0.6, 1.2, 1.6, 0.3, shapecolor, "gray")
	canvas.Circle(c3, y+4, dotsize, dotcolor)
	y -= yspace
	// grid
	gy1 := y - 2.5
	gy2 := y + 2.5
	canvas.Grid(c3, c4, gy1, gy2, 0.1, 2.5, shapecolor)
	canvas.Circle(85, gy1, dotsize, dotcolor)
	canvas.Circle(95, gy2, dotsize, dotcolor)

	// usage
	canvas.NewPage(cw, ch)
	canvas.Background("black")

	for range 500 {
		xr := rand.Float64() * 100
		yr := rand.Float64() * 100
		canvas.Circle(xr, yr, 0.25, "white")
	}
	canvas.Circle(50, 0, 50, "blue")
	canvas.CText(50, 20, 10, "hello, world", "white")
	codesize := 1.0
	border := codesize / 8
	canvas.TextCode("excode", 2, 98, 40, 45, codesize, 1.5, border, "black", "white")

	canvas.NewPage(cw, ch)
	err = canvas.LoadFont(fontmap["serif"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	palette := []string{"red", "green", "blue", "orange", "black"}
	a := 0.0
	n := 10
	for i := 0; i < n; i++ {
		color := palette[i%(len(palette))]
		canvas.RText(50, 50, 50, a, "i", color)
		a += 360 / float64(n)
	}

	err = canvas.Creator.WriteToFile(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
