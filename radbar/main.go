package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ajstarks/gpdf"
)

const (
	smallest     = -math.MaxFloat64
	canvasWidth  = 792
	canvasHeight = 612
)

type NameValue struct {
	name  string
	value float64
}

type options struct {
	cx     float64
	cy     float64
	lw     float64
	radius float64
	xint   float64
	yint   float64
	color  string
	title  string
}

// vmap maps one interval to another
func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func readData(r io.Reader) ([]NameValue, float64, string, error) {
	var (
		data  []NameValue
		d     NameValue
		title string
	)
	maxval := smallest

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 { // skip blank lines
			continue
		}
		if strings.HasPrefix(t, "#") {
			title = strings.TrimSpace(t[1:])
			continue
		}
		fields := strings.Split(t, "\t")
		if len(fields) != 2 {
			continue
		}
		v, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			continue
		}
		d.name = fields[0]
		d.value = v
		if d.value > maxval {
			maxval = d.value
		}
		data = append(data, d)
	}
	return data, maxval, title, scanner.Err()
}

func colorop(s string) (string, float64) {
	c := strings.Split(s, ":")
	if len(c) == 2 {
		o, _ := strconv.ParseFloat(c[1], 64)
		return c[0], o
	}
	return s, 100
}

func radbar(canvas *gpdf.Canvas, filename string, data []NameValue, maxval float64, opts options) {
	//grid(w, 2.5)
	xint := opts.xint
	rad := opts.radius
	lw := opts.lw
	yint := opts.yint
	ts := 2.5
	cx, cy, color := opts.cx, opts.cy, opts.color
	if len(opts.title) > 0 {
		canvas.CText(cx, cy+rad+(ts*2), ts, opts.title, "black")
	}
	canvas.Circle(cx, cy, 1, color)

	for a := 0.0; a <= 180; a += 180 / yint {
		px, py := canvas.PolarDegrees(cx, cy, rad+(xint/2), a)
		canvas.Line(cx, cy, px, py, 0.1, "gray:50")
		l := vmap(a, 0, 180, 1000, 0)
		canvas.CText(px, py, ts/2, fmt.Sprintf("%v", l), "black")
	}

	x := rad * 2
	lx := cx - rad
	li := xint / 2
	hts := ts / 2
	for i := range data {
		v := vmap(data[i].value, 0, maxval, 180, 0)
		if i%2 == 0 {
			canvas.CText(lx, cy-hts, min(li, hts), data[i].name, "black")
		} else {
			canvas.CText(lx, cy-ts, min(li, hts), data[i].name, "black")
		}
		canvas.Arc(cx, cy, x/2, v, 180, lw, color)
		x -= xint
		lx += li
	}
}

func main() {
	var opts options
	flag.StringVar(&opts.color, "color", "maroon", "line color")
	flag.StringVar(&opts.title, "title", "", "title")
	flag.Float64Var(&opts.cx, "cx", 50, "center x")
	flag.Float64Var(&opts.cy, "cy", 40, "center y")
	flag.Float64Var(&opts.lw, "lw", 0.1, "line width")
	flag.Float64Var(&opts.xint, "xint", 2, "x-interval")
	flag.Float64Var(&opts.yint, "yint", 10, "y-interval")
	flag.Float64Var(&opts.radius, "r", 45, "chart radius")
	flag.Parse()

	data, maxval, title, err := readData(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if opts.title == "" {
		opts.title = title
	}

	output := "f.pdf"
	args := flag.Args()
	if len(args) > 0 {
		output = args[0]
	}
	ps := gpdf.Letter
	canvas, err := gpdf.SetupCanvas(ps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = canvas.LoadFont("PublicSans-Regular.ttf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	radbar(canvas, output, data, maxval, opts)
	err = canvas.Creator.WriteToFile(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
