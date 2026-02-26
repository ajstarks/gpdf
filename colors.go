package gpdf

import (
	"math"
	"strconv"
	"strings"

	"github.com/coregx/gxpdf/creator"
)

var colornames = map[string]creator.ColorRGBA{
	"aliceblue":            {R: 0.941, G: 0.973, B: 1.000, A: 1},
	"antiquewhite":         {R: 0.980, G: 0.922, B: 0.843, A: 1},
	"aqua":                 {R: 0.000, G: 1.000, B: 1.000, A: 1},
	"aquamarine":           {R: 0.498, G: 1.000, B: 0.831, A: 1},
	"azure":                {R: 0.941, G: 1.000, B: 1.000, A: 1},
	"beige":                {R: 0.961, G: 0.961, B: 0.863, A: 1},
	"bisque":               {R: 1.000, G: 0.894, B: 0.769, A: 1},
	"black":                {R: 0.000, G: 0.000, B: 0.000, A: 1},
	"blanchedalmond":       {R: 1.000, G: 0.922, B: 0.804, A: 1},
	"blue":                 {R: 0.000, G: 0.000, B: 1.000, A: 1},
	"blueviolet":           {R: 0.541, G: 0.169, B: 0.886, A: 1},
	"brown":                {R: 0.647, G: 0.165, B: 0.165, A: 1},
	"burlywood":            {R: 0.871, G: 0.722, B: 0.529, A: 1},
	"cadetblue":            {R: 0.373, G: 0.620, B: 0.627, A: 1},
	"chartreuse":           {R: 0.498, G: 1.000, B: 0.000, A: 1},
	"chocolate":            {R: 0.824, G: 0.412, B: 0.118, A: 1},
	"coral":                {R: 1.000, G: 0.498, B: 0.314, A: 1},
	"cornflowerblue":       {R: 0.392, G: 0.584, B: 0.929, A: 1},
	"cornsilk":             {R: 1.000, G: 0.973, B: 0.863, A: 1},
	"crimson":              {R: 0.863, G: 0.078, B: 0.235, A: 1},
	"cyan":                 {R: 0.000, G: 1.000, B: 1.000, A: 1},
	"darkblue":             {R: 0.000, G: 0.000, B: 0.545, A: 1},
	"darkcyan":             {R: 0.000, G: 0.545, B: 0.545, A: 1},
	"darkgoldenrod":        {R: 0.722, G: 0.525, B: 0.043, A: 1},
	"darkgray":             {R: 0.663, G: 0.663, B: 0.663, A: 1},
	"darkgreen":            {R: 0.000, G: 0.392, B: 0.000, A: 1},
	"darkgrey":             {R: 0.663, G: 0.663, B: 0.663, A: 1},
	"darkkhaki":            {R: 0.741, G: 0.718, B: 0.420, A: 1},
	"darkmagenta":          {R: 0.545, G: 0.000, B: 0.545, A: 1},
	"darkolivegreen":       {R: 0.333, G: 0.420, B: 0.184, A: 1},
	"darkorange":           {R: 1.000, G: 0.549, B: 0.000, A: 1},
	"darkorchid":           {R: 0.600, G: 0.196, B: 0.800, A: 1},
	"darkred":              {R: 0.545, G: 0.000, B: 0.000, A: 1},
	"darksalmon":           {R: 0.914, G: 0.588, B: 0.478, A: 1},
	"darkseagreen":         {R: 0.561, G: 0.737, B: 0.561, A: 1},
	"darkslateblue":        {R: 0.282, G: 0.239, B: 0.545, A: 1},
	"darkslategray":        {R: 0.184, G: 0.310, B: 0.310, A: 1},
	"darkslategrey":        {R: 0.184, G: 0.310, B: 0.310, A: 1},
	"darkturquoise":        {R: 0.000, G: 0.808, B: 0.820, A: 1},
	"darkviolet":           {R: 0.580, G: 0.000, B: 0.827, A: 1},
	"deeppink":             {R: 1.000, G: 0.078, B: 0.576, A: 1},
	"deepskyblue":          {R: 0.000, G: 0.749, B: 1.000, A: 1},
	"dimgray":              {R: 0.412, G: 0.412, B: 0.412, A: 1},
	"dimgrey":              {R: 0.412, G: 0.412, B: 0.412, A: 1},
	"dodgerblue":           {R: 0.118, G: 0.565, B: 1.000, A: 1},
	"firebrick":            {R: 0.698, G: 0.133, B: 0.133, A: 1},
	"floralwhite":          {R: 1.000, G: 0.980, B: 0.941, A: 1},
	"forestgreen":          {R: 0.133, G: 0.545, B: 0.133, A: 1},
	"fuchsia":              {R: 1.000, G: 0.000, B: 1.000, A: 1},
	"gainsboro":            {R: 0.863, G: 0.863, B: 0.863, A: 1},
	"ghostwhite":           {R: 0.973, G: 0.973, B: 1.000, A: 1},
	"gold":                 {R: 1.000, G: 0.843, B: 0.000, A: 1},
	"goldenrod":            {R: 0.855, G: 0.647, B: 0.125, A: 1},
	"gray":                 {R: 0.502, G: 0.502, B: 0.502, A: 1},
	"green":                {R: 0.000, G: 0.502, B: 0.000, A: 1},
	"greenyellow":          {R: 0.678, G: 1.000, B: 0.184, A: 1},
	"grey":                 {R: 0.502, G: 0.502, B: 0.502, A: 1},
	"honeydew":             {R: 0.941, G: 1.000, B: 0.941, A: 1},
	"hotpink":              {R: 1.000, G: 0.412, B: 0.706, A: 1},
	"indianred":            {R: 0.804, G: 0.361, B: 0.361, A: 1},
	"indigo":               {R: 0.294, G: 0.000, B: 0.510, A: 1},
	"ivory":                {R: 1.000, G: 1.000, B: 0.941, A: 1},
	"khaki":                {R: 0.941, G: 0.902, B: 0.549, A: 1},
	"lavender":             {R: 0.902, G: 0.902, B: 0.980, A: 1},
	"lavenderblush":        {R: 1.000, G: 0.941, B: 0.961, A: 1},
	"lawngreen":            {R: 0.486, G: 0.988, B: 0.000, A: 1},
	"lemonchiffon":         {R: 1.000, G: 0.980, B: 0.804, A: 1},
	"lightblue":            {R: 0.678, G: 0.847, B: 0.902, A: 1},
	"lightcoral":           {R: 0.941, G: 0.502, B: 0.502, A: 1},
	"lightcyan":            {R: 0.878, G: 1.000, B: 1.000, A: 1},
	"lightgoldenrodyellow": {R: 0.980, G: 0.980, B: 0.824, A: 1},
	"lightgray":            {R: 0.827, G: 0.827, B: 0.827, A: 1},
	"lightgreen":           {R: 0.565, G: 0.933, B: 0.565, A: 1},
	"lightgrey":            {R: 0.827, G: 0.827, B: 0.827, A: 1},
	"lightpink":            {R: 1.000, G: 0.714, B: 0.757, A: 1},
	"lightsalmon":          {R: 1.000, G: 0.627, B: 0.478, A: 1},
	"lightseagreen":        {R: 0.125, G: 0.698, B: 0.667, A: 1},
	"lightskyblue":         {R: 0.529, G: 0.808, B: 0.980, A: 1},
	"lightslategray":       {R: 0.467, G: 0.533, B: 0.600, A: 1},
	"lightslategrey":       {R: 0.467, G: 0.533, B: 0.600, A: 1},
	"lightsteelblue":       {R: 0.690, G: 0.769, B: 0.871, A: 1},
	"lightyellow":          {R: 1.000, G: 1.000, B: 0.878, A: 1},
	"lime":                 {R: 0.000, G: 1.000, B: 0.000, A: 1},
	"limegreen":            {R: 0.196, G: 0.804, B: 0.196, A: 1},
	"linen":                {R: 0.980, G: 0.941, B: 0.902, A: 1},
	"magenta":              {R: 1.000, G: 0.000, B: 1.000, A: 1},
	"maroon":               {R: 0.502, G: 0.000, B: 0.000, A: 1},
	"mediumaquamarine":     {R: 0.400, G: 0.804, B: 0.667, A: 1},
	"mediumblue":           {R: 0.000, G: 0.000, B: 0.804, A: 1},
	"mediumorchid":         {R: 0.729, G: 0.333, B: 0.827, A: 1},
	"mediumpurple":         {R: 0.576, G: 0.439, B: 0.859, A: 1},
	"mediumseagreen":       {R: 0.235, G: 0.702, B: 0.443, A: 1},
	"mediumslateblue":      {R: 0.482, G: 0.408, B: 0.933, A: 1},
	"mediumspringgreen0":   {R: 0.980, G: 0.604, B: 1.000, A: 0},
	"mediumturquoise":      {R: 0.282, G: 0.820, B: 0.800, A: 1},
	"mediumvioletred":      {R: 0.780, G: 0.082, B: 0.522, A: 1},
	"midnightblue":         {R: 0.098, G: 0.098, B: 0.439, A: 1},
	"mintcream":            {R: 0.961, G: 1.000, B: 0.980, A: 1},
	"mistyrose":            {R: 1.000, G: 0.894, B: 0.882, A: 1},
	"moccasin":             {R: 1.000, G: 0.894, B: 0.710, A: 1},
	"navajowhite":          {R: 1.000, G: 0.871, B: 0.678, A: 1},
	"navy":                 {R: 0.000, G: 0.000, B: 0.502, A: 1},
	"oldlace":              {R: 0.992, G: 0.961, B: 0.902, A: 1},
	"olive":                {R: 0.502, G: 0.502, B: 0.000, A: 1},
	"olivedrab":            {R: 0.420, G: 0.557, B: 0.137, A: 1},
	"orange":               {R: 1.000, G: 0.647, B: 0.000, A: 1},
	"orangered":            {R: 1.000, G: 0.271, B: 0.000, A: 1},
	"orchid":               {R: 0.855, G: 0.439, B: 0.839, A: 1},
	"palegoldenrod":        {R: 0.933, G: 0.910, B: 0.667, A: 1},
	"palegreen":            {R: 0.596, G: 0.984, B: 0.596, A: 1},
	"paleturquoise":        {R: 0.686, G: 0.933, B: 0.933, A: 1},
	"palevioletred":        {R: 0.859, G: 0.439, B: 0.576, A: 1},
	"papayawhip":           {R: 1.000, G: 0.937, B: 0.835, A: 1},
	"peachpuff":            {R: 1.000, G: 0.855, B: 0.725, A: 1},
	"peru":                 {R: 0.804, G: 0.522, B: 0.247, A: 1},
	"pink":                 {R: 1.000, G: 0.753, B: 0.796, A: 1},
	"plum":                 {R: 0.867, G: 0.627, B: 0.867, A: 1},
	"powderblue":           {R: 0.690, G: 0.878, B: 0.902, A: 1},
	"purple":               {R: 0.502, G: 0.000, B: 0.502, A: 1},
	"red":                  {R: 1.000, G: 0.000, B: 0.000, A: 1},
	"rosybrown":            {R: 0.737, G: 0.561, B: 0.561, A: 1},
	"royalblue":            {R: 0.255, G: 0.412, B: 0.882, A: 1},
	"saddlebrown":          {R: 0.545, G: 0.271, B: 0.075, A: 1},
	"salmon":               {R: 0.980, G: 0.502, B: 0.447, A: 1},
	"sandybrown":           {R: 0.957, G: 0.643, B: 0.376, A: 1},
	"seagreen":             {R: 0.180, G: 0.545, B: 0.341, A: 1},
	"seashell":             {R: 1.000, G: 0.961, B: 0.933, A: 1},
	"sienna":               {R: 0.627, G: 0.322, B: 0.176, A: 1},
	"silver":               {R: 0.753, G: 0.753, B: 0.753, A: 1},
	"skyblue":              {R: 0.529, G: 0.808, B: 0.922, A: 1},
	"slateblue":            {R: 0.416, G: 0.353, B: 0.804, A: 1},
	"slategray":            {R: 0.439, G: 0.502, B: 0.565, A: 1},
	"slategrey":            {R: 0.439, G: 0.502, B: 0.565, A: 1},
	"snow":                 {R: 1.000, G: 0.980, B: 0.980, A: 1},
	"springgreen":          {R: 0.000, G: 1.000, B: 0.498, A: 1},
	"steelblue":            {R: 0.275, G: 0.510, B: 0.706, A: 1},
	"tan":                  {R: 0.824, G: 0.706, B: 0.549, A: 1},
	"teal":                 {R: 0.000, G: 0.502, B: 0.502, A: 1},
	"thistle":              {R: 0.847, G: 0.749, B: 0.847, A: 1},
	"tomato":               {R: 1.000, G: 0.388, B: 0.278, A: 1},
	"turquoise":            {R: 0.251, G: 0.878, B: 0.816, A: 1},
	"violet":               {R: 0.933, G: 0.510, B: 0.933, A: 1},
	"wheat":                {R: 0.961, G: 0.871, B: 0.702, A: 1},
	"white":                {R: 1.000, G: 1.000, B: 1.000, A: 1},
	"whitesmoke":           {R: 0.961, G: 0.961, B: 0.961, A: 1},
	"yellow":               {R: 1.000, G: 1.000, B: 0.000, A: 1},
	"yellowgreen":          {R: 0.604, G: 0.804, B: 0.196, A: 1},
}

func colorop(s string) (string, float64) {
	co := strings.Split(s, ":")
	if len(co) != 2 {
		return s, 100
	}
	o, err := strconv.ParseFloat(co[1], 64)
	if err != nil {
		o = 100
	}
	return co[0], o
}

// cc converts a color string to number
func cc(s string) float64 {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return float64(v) / 255.0
}

// hc converts a hex color string to number
func hc(s string) float64 {
	v, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		return 0
	}
	return float64(v) / 255.0
}

// ColorLookup returns a color.RGBA corresponding to the named color or
// "rgb(r)", "rgb(r,b)", "rgb(r,g,b), "rgb(r,g,b,a)",
// "#rr",     "#rrgg",   "#rrggbb",   "#rrggbbaa" string.
// "hsv(hue,sat,value)"
// On error, return black.
func ColorLookup(s string) creator.ColorRGBA {
	c, ok := colornames[s]
	if ok {
		return c
	}
	black := creator.ColorRGBA{R: 0, G: 0, B: 0, A: 1}
	ls := len(s)
	// rgb(...)
	if strings.HasPrefix(s, "rgb(") && strings.HasSuffix(s, ")") && ls > 5 {
		c.R, c.G, c.B, c.A = 0, 0, 0, 1.0
		v := strings.Split(s[4:ls-1], ",")
		switch len(v) {
		case 1:
			c.R = cc(v[0])
			return c
		case 2:
			c.R = cc(v[0])
			c.G = cc(v[1])
			return c
		case 3:
			c.R = cc(v[0])
			c.G = cc(v[1])
			c.B = cc(v[2])
			return c
		case 4:
			c.R = cc(v[0])
			c.G = cc(v[1])
			c.B = cc(v[2])
			c.A = cc(v[3])
			return c
		default:
			return black
		}
	}
	// hsv(h,s,v) or hsv(h,s,v,a); h=0-360, s, v=0-100, a=0-100
	if strings.HasPrefix(s, "hsv(") && strings.HasSuffix(s, ")") && ls > 5 {
		v := strings.Split(s[4:ls-1], ",")
		switch len(v) {
		case 3:
			hue, _ := strconv.ParseFloat(v[0], 64)
			sat, _ := strconv.ParseFloat(v[1], 64)
			value, _ := strconv.ParseFloat(v[2], 64)
			c.R, c.G, c.B = hsv2rgb(hue, sat, value)
			c.A = 1
			return c
		case 4:
			hue, _ := strconv.ParseFloat(v[0], 64)
			sat, _ := strconv.ParseFloat(v[1], 64)
			value, _ := strconv.ParseFloat(v[2], 64)
			a := cc(v[3])
			if a > 100 {
				a = 100
			}
			c.R, c.G, c.B = hsv2rgb(hue, sat, value)
			c.A = float64(a) / 100.0
			return c
		default:
			return black
		}
	}
	// #rrggbb
	if strings.HasPrefix(s, "#") && (ls >= 3) {
		c.R, c.G, c.B, c.A = 0, 0, 0, 255
		switch ls {
		case 3:
			c.R = hc(s[1:3])
		case 5:
			c.R = hc(s[1:3])
			c.G = hc(s[3:5])
		case 7:
			c.R = hc(s[1:3])
			c.G = hc(s[3:5])
			c.B = hc(s[5:7])
		case 9:
			c.R = hc(s[1:3])
			c.G = hc(s[3:5])
			c.B = hc(s[5:7])
			c.A = hc(s[7:9])
		default:
			return black
		}
		return c
	}
	return black
}

// hsv2rgb converts hsv(h (0-360), s (0-100), v (0-100)) to rgb
// reference: https://en.wikipedia.org/wiki/HSL_and_HSV#HSV_to_RGB
func hsv2rgb(h, s, v float64) (float64, float64, float64) {
	s /= 100
	v /= 100
	if s > 1 || v > 1 {
		return 0, 0, 0
	}
	h = math.Mod(h, 360)
	c := v * s
	section := h / 60
	x := c * (1 - math.Abs(math.Mod(section, 2)-1))

	var r, g, b float64
	switch {
	case section >= 0 && section <= 1:
		r = c
		g = x
		b = 0
	case section > 1 && section <= 2:
		r = x
		g = c
		b = 0
	case section > 2 && section <= 3:
		r = 0
		g = c
		b = x
	case section > 3 && section <= 4:
		r = 0
		g = x
		b = c
	case section > 4 && section <= 5:
		r = x
		g = 0
		b = c
	case section > 5 && section <= 6:
		r = c
		g = 0
		b = x
	default:
		return 0, 0, 0
	}
	m := v - c
	r += m
	g += m
	b += m
	return r, g, b
}
