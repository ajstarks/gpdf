// hello, world
package main

import (
	"math/rand/v2"

	"github.com/ajstarks/gpdf"
)

func main() {
	canvas, err := gpdf.SetupCanvas(792, 612)
	if err != nil {
		return
	}

	canvas.Background("black")
	for range 500 {
		xr := rand.Float64() * 100
		yr := rand.Float64() * 100
		canvas.Circle(xr, yr, 0.25, "white")
	}
	canvas.Circle(50, 0, 50, "blue")
	canvas.Text(25, 20, 10, "hello, world", "white")

	err = canvas.Creator.WriteToFile("hello.pdf")
	if err != nil {
		return
	}
}
