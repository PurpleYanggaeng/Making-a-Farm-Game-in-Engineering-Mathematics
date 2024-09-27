package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	springRadius    = 5
	springCoilCount = 10
	precisionCount  = 500
)

type Game struct {
	Stat     int
	SpringXY [2][]int
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.SpringXY[g.Stat%2] = []int{x, y}
		g.Stat++
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x0, y0 := float64(g.SpringXY[0][0]), float64(g.SpringXY[0][1])
	x1, y1 := float64(g.SpringXY[1][0]), float64(g.SpringXY[1][1])
	for i := 0; i < precisionCount; i++ {
		x, y := float64(x0)+(x1-x0)*float64(i)/float64(precisionCount), float64(y0)+(y1-y0)*float64(i)/float64(precisionCount)
		x += float64(springRadius) * math.Cos(float64(springCoilCount*i)*2*math.Pi/float64(precisionCount))
		y += float64(springRadius) * math.Sin(float64(springCoilCount*i)*2*math.Pi/float64(precisionCount))
		screen.Set(int(x), int(y), color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("2d spring")
	if err := ebiten.RunGame(&Game{
		SpringXY: [2][]int{
			{0, 0},
			{0, 0},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
