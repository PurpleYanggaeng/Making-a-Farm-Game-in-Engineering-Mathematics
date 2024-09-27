package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth    = 640
	screenHeight   = 480
	radius         = 10
	precisionCount = 50
	gameEndCount   = 1
	moveSpeed      = 2
)

type Game struct {
	TargetX, TargetY int
	X, Y             int
	hitCount         int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.X -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.X += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Y -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Y += moveSpeed
	}
	if dist(g.X, g.Y, g.TargetX, g.TargetY) <= 2*radius {
		g.hitCount++
		g.TargetX = rand.Intn(screenWidth)
		g.TargetY = rand.Intn(screenHeight)
	}
	return nil
}

func dist(x, y, p, q int) int {
	return int(math.Sqrt(float64((x-p)*(x-p) + (y-q)*(y-q))))
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.hitCount < gameEndCount {
		drawCircle(g.X, g.Y, screen)
		drawFilledCircle(g.TargetX, g.TargetY, screen)
	} else {
		ebitenutil.DebugPrint(screen, "Done")
	}
}

func drawCircle(x, y int, screen *ebiten.Image) {
	for i := 0; i < precisionCount; i++ {
		screen.Set(x+int(float64(radius)*math.Cos(2*math.Pi*float64(i)/precisionCount)),
			y+int(float64(radius)*math.Sin(2*math.Pi*float64(i)/precisionCount)),
			color.White)
	}
}

func drawFilledCircle(x, y int, screen *ebiten.Image) {
	for i := 0; i < precisionCount; i++ {
		for j := 0; j < radius; j++ {
			screen.Set(x+int(float64(j)*math.Cos(2*math.Pi*float64(i)/precisionCount)),
				y+int(float64(j)*math.Sin(2*math.Pi*float64(i)/precisionCount)),
				color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("A white dot")
	if err := ebiten.RunGame(&Game{
		TargetX: rand.Intn(screenWidth),
		TargetY: rand.Intn(screenHeight),
	}); err != nil {
		log.Fatal(err)
	}
}
