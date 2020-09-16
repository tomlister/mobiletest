package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	x, y, z    float64
	x2, y2, z2 float64
	x3, y3, z3 float64
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	x, y, z := ebiten.GetSensorAccelerometer()
	g.x = x
	g.y = y
	g.z = z
	x, y, z = ebiten.GetSensorGyroscope()
	g.x2 = x
	g.y2 = y
	g.z2 = z
	x, y, z = ebiten.GetSensorMagnetometer()
	g.x3 = x
	g.y3 = y
	g.z3 = z
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Accelerometer\nx %f, y %f, z %f", g.x, g.y, g.z))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nGyroscope\nx %f, y %f, z %f", g.x2, g.y2, g.z2))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\n\nMagnetometer\nx %f, y %f, z %f", g.x3, g.y3, g.z3))
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	fmt.Println("test")
	log.Println("testlog")
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
