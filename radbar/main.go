// radbar - radial barchart
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
	smallest   = -math.MaxFloat64
	beginAngle = 0.0
	endAngle   = 180.0
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
	ymax   float64
	rtext  float64
	color  string
	title  string
	output string
}

// vmap maps one interval to another
func vmap(value float64, low1 float64, high1 float64, low2 float64, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// readData reads tab separated name, value pairs
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

// radbar makes a radial bar chart
func radbar(canvas *gpdf.Canvas, w io.Writer, r io.Reader, opts options) error {
	// read data
	data, maxval, title, err := readData(r)
	if err != nil {
		return err
	}
	if opts.title == "" {
		opts.title = title
	}
	if opts.ymax == -1 {
		opts.ymax = maxval
	}
	xint := opts.xint
	rad := opts.radius
	lw := opts.lw
	yint := opts.yint
	maxval = opts.ymax
	ts := 2.5

	margin := ts
	cx, cy, color := opts.cx, opts.cy, opts.color

	// title
	if len(opts.title) > 0 {
		canvas.CText(cx, cy+rad+(ts*2), ts, opts.title, "black")
	}
	// data axis
	canvas.Circle(cx, cy, 1, color)
	for i := 0.0; i <= maxval; i += yint {
		a := vmap(i, 0, maxval, endAngle, beginAngle)
		px, py := canvas.PolarDegrees(cx, cy, (rad+(xint/2))-margin, a)
		canvas.Line(cx, cy, px, py, 0.1, "gray:50")
		canvas.CText(px, py, ts/2, strconv.FormatFloat(i, 'g', -1, 64), "black")
	}
	// data arcs
	x := rad * 2
	lx := cx - rad
	li := xint / 2
	hts := ts / 2
	l1y := cy - hts
	l2y := l1y - (hts * .6)
	for i := range data {
		// staggered x labels
		if opts.rtext == 0 {
			if i%2 == 0 {
				canvas.CText(lx, l1y, min(li, hts), data[i].name, "black")
			} else {
				canvas.CText(lx, l2y, min(li, hts), data[i].name, "black")
			}
		} else {
			canvas.RText(lx, l2y-ts, min(li, hts), opts.rtext, data[i].name, "black")
		}
		// map data to an angle
		v := vmap(data[i].value, 0, maxval, endAngle, beginAngle)
		canvas.Arc(cx, cy, x/2, v, 180, lw, color)
		x -= xint
		lx += li
	}
	// write output
	_, err = canvas.Creator.WriteTo(w)
	return err
}

func main() {
	var opts options
	flag.StringVar(&opts.color, "color", "steelblue", "line color")
	flag.StringVar(&opts.title, "title", "", "title")
	flag.StringVar(&opts.output, "o", "", "output file")
	flag.Float64Var(&opts.cx, "cx", 50, "center x")
	flag.Float64Var(&opts.cy, "cy", 40, "center y")
	flag.Float64Var(&opts.lw, "lw", 0.1, "line width")
	flag.Float64Var(&opts.xint, "xint", 2, "x-interval")
	flag.Float64Var(&opts.yint, "yint", 10, "y-interval")
	flag.Float64Var(&opts.ymax, "ymax", -1.0, "y-max")
	flag.Float64Var(&opts.rtext, "rot", 0, "text rotation for labels")
	flag.Float64Var(&opts.radius, "r", 45, "chart radius")
	flag.Parse()

	var (
		r   io.Reader
		w   io.Writer
		err error
	)
	// default reader and writer
	r = os.Stdin
	w = os.Stdout
	// open input file; if no file specified for input, use stdin
	args := flag.Args()
	if len(args) > 0 {
		r, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	// create output file; if no file specfied for output, use stdout
	if len(opts.output) > 0 {
		w, err = os.Create(opts.output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}
	}
	// set up the canvas
	canvas, err := gpdf.SetupCanvas(792, 612)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	// load the font
	ff, err := canvas.LoadFontFile("PublicSans-Regular.ttf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(4)
	}
	canvas.SetFont(ff)
	// make the chart
	err = radbar(canvas, w, r, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}

}
