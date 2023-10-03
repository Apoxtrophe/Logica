package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

func (g *Game) MouseInteract() {
	x, y := ebiten.CursorPosition()
	ix, iy := x/worldFactor, y/worldFactor

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if ix >= 0 && ix < len(g.array1) && iy >= 0 && iy < len(g.array1[0]) {
			g.array1[ix][iy] = 1
		}
	}
}

func (g *Game) DrawArray(screen *ebiten.Image) {
	pixelSize := ebiten.NewImage(worldFactor, worldFactor)
	pixelSize.Fill(colornames.Orange)
	for x := 0; x < len(g.array1); x++ {
		for y := 0; y < len(g.array1[x]); y++ {
			element, exists := elementMap[g.array1[x][y]]
			if exists {
				g.pixelImage.Fill(element.Color)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*worldFactor), float64(y*worldFactor))
				screen.DrawImage(pixelSize, op)
			}
		}
	}

	//Debug ish
	tps := ebiten.ActualTPS()
	debugText := fmt.Sprintf("TPS: %0.2f", tps)
	ebitenutil.DebugPrint(screen, debugText)
}
