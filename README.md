# gpdf
Generate PDF using high-level objects (text, lines, curves, shapes) using a percentage-based coordinate system.

## Placemat

![gpdf-api](gpdf-api.png)

## Initialize

```canvas, err := SetupCanvas(width, height float64) (*Canvas, error)```

## Set page background

```(c *Canvas) Background(color string)```

## Line between (x0,y0) and (x1, y1)

```(c *Canvas) Line(x0, y0, x1, y1, size float64, color string, opacity ...float64)```

## Filled circle centered at (x,y) with radius r

```(c *Canvas) Circle(x, y, r float64, color string, opacity ...float64)```

## Filled ellipse centered at (x,y), dimensions (w,h)

```(c *Canvas) Ellipse(x, y, w, h float64, color string, opacity ...float64)```

## Filled rectangle upper left corner at (x,y), dimensions (w,h)

```(c *Canvas) CornerRect(x, y, w, h float64, color string, opacity ...float64)```

## Filled rectangle centered  at (x,y), dimensions (w,h)

```(c *Canvas) Rect(x, y, w, h float64, color string, opacity ...float64)```

## Gradient filled rectangle 

```(c *Canvas) GradRect(x, y, w, h float64, color1 string, color2 string, percent float64)```

## Filled square centered at (x,y), both sized are w.

```(c *Canvas) Square(x,y, w, float64 color string, opacity ...float64)```

## Text beginning at (x,y) at the specified size and color

```(c *Canvas) Text(x, y, size float64, s string, color string, opacity ...float64)```

```(c *Canvas) BText(x, y, size float64, s string, color string, opacity ...float64)```

## Text centered at (x,y) at the specified size and color

```(c *Canvas) CText(x, y, size float64, s string, color string, opacity ...float64)```

## Text end-aligned at (x,y) at the specified size and color

```(c *Canvas) EText(x, y, size float64, s string, color string, opacity ...float64)```

## Rotated text at the specified angle, anchored at (x,y), with the specified size and color

```(c *Canvas) RText(x, y, size, angle float64, s string, color string, opacity ...float64)```

## Filled polygon with specified coordinates

```(c *Canvas) Polygon(x, y []float64, color string, opacity ...float64)```

## Stroked polyline with specified coordinates

```(c *Canvas) Polyline(x, y []float64, size, color string, opacity ...float64)```

## Cubic Bezier curve; begin (bx,by), control1 (cx1,cy1), control2 (cx2,cy2), end (ex,ey), with specifed stroke size and color

```(c *Canvas) CubicCurve(bx, by, cx1, cy1, cx2, cy2, ex, ey, size float64, strokecolor string, opacity ...float64)```

## Quadradic Bezier curve; begin (bx,by), control (cx,cy), end (ex,ey), with specifed stroke size and color

```(c *Canvas) QuadCurve(bx, by, cx, cy, ex, ey, size float64, color string, opacity ...float64)```

## Circular arc; center at (x,y), radius r, between angles a1 and s2 (degrees), with specifed stroke size and color

```(c *Canvas) Arc(x, y, r, a1, a2, size float64, fillcolor string, opacity ...float64)```

## Image centered at (x,y), dimensions (w, h)

```(c *Canvas) ImageName(x, y, w, h float64, name string)```

## Plain and labeled grids between (xmin,ymin) and (xmax,ymax) at the specified increment, using specified size and color

```(c *Canvas) Grid(xmin, xmax, ymin, ymax, size, incr float64, color string)```

```(c *Canvas) LGrid(xmin, xmax, ymin, ymax, size, incr float64, color string, opacity ...float64)```


## hello, world
![hello](hello.png)

