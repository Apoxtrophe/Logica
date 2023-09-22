package main 

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type Element struct {
	Color color.RGBA
	Index int
}

var elementMap = map[int]Element{
	1: {Color: colornames.Grey, Index: 1}, //Wire off
	2: {Color: colornames.White, Index: 2}, //Wire On
	15: {Color: colornames.Brown, Index: 15}, //Battery
}