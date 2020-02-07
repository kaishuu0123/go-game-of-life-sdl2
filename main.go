package main

import (
	"os"

	"go-game-of-life-sdl2/life"
	"go-game-of-life-sdl2/sdl2canvas"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  int = 640
	windowHeight int = 640
	cellCount    int = 80
	cellSize     int = windowWidth / cellCount
)

var (
	aliveColor  sdl2canvas.Color = sdl2canvas.Color{R: 1, G: 220, B: 200}
	deadColor   sdl2canvas.Color = sdl2canvas.Color{R: 0, G: 0, B: 0}
	borderColor sdl2canvas.Color = sdl2canvas.Color{R: 50, G: 50, B: 50}
)

func main() {
	var canvas sdl2canvas.SDL2Canvas
	var life life.Life

	canvas.Setup("Conway's Game of Life", windowWidth, windowHeight)
	life.Setup(cellCount)

	life.InitCurrentField()

	for canvas.Running {
		canvas.HandleEvents()

		// Handle Mouse Event
		if canvas.MouseClicked {
			x := int(canvas.MouseX) / cellSize
			y := int(canvas.MouseY) / cellSize
			life.SetCurrentField(x, y, true)
		}

		life.Step()

		// mapping life field to sdl2canvas
		for y := 0; y < windowHeight; y++ {
			for x := 0; x < windowWidth; x++ {
				if (y%cellSize) == 0 || (x%cellSize) == 0 {
					canvas.SetPixel(x, y, borderColor)
				} else if life.Current.Alive(life.CellCount, x/cellSize, y/cellSize) {
					canvas.SetPixel(x, y, aliveColor)
				} else {
					canvas.SetPixel(x, y, deadColor)
				}
			}
		}

		canvas.Update()
		canvas.Render()

		sdl.Delay(100)
	}

	canvas.Shutdown()
	os.Exit(0)
}
