package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)


func (g *Game) MouseInteract(){
	x,y := ebiten.CursorPosition()
	ix , iy := x / worldFactor, y / worldFactor

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
		if ix >= 0 && ix < len(g.array1) && iy >= 0 && iy < len(g.array1[0]){
			g.array1[ix][iy] = 1
		}
	}
}

func (g *Game) DrawArray(screen *ebiten.Image) {
	white := color.RGBA{255,255,255,255}

	pixelSize := ebiten.NewImage(worldFactor,worldFactor)
	pixelSize.Fill(white)
	for x := 0; x < len(g.array1); x++ {
		for y := 0; y < len(g.array1[x]); y++ {
			if g.array1[x][y] == 1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x * worldFactor), float64(y * worldFactor))
				screen.DrawImage(pixelSize, op)
			}
		}
	}
}