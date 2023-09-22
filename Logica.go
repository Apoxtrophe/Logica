package main

import (
	"log"
	_ "net/http/pprof"
    "net/http"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
	worldFactor  = 10
)

var (
	worldWidth  = screenWidth / worldFactor
	worldHeight = screenHeight / worldFactor
)

// Game implements ebiten.Game interface.
type Game struct {
	array1 [][]int
	array2 [][]int
	pixelImage *ebiten.Image
}

func NewGame() *Game {
	g := &Game{}
	g.array1 = make([][]int, worldWidth)
	g.array2 = make([][]int, worldWidth)
	for i := range g.array1 {
		g.array1[i] = make([]int, worldHeight)
		g.array2[i] = make([]int, worldHeight)
	}
	g.pixelImage = ebiten.NewImage(worldFactor, worldFactor)

	return g
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.MouseInteract()
	// Write your game's update logic here
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawArray(screen)
	// Write your game's draw logic here
}

// Layout takes the outside size (e.g., the window size) and returns (logical screen width, logical screen height).
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set the screen size
	return screenWidth, screenHeight
}

func main() {
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Logica")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
