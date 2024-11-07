package greet

import (
	"fun-with-golang/ch1/functions"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"image/color"
)

// Layout handles rendering and input.
func Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return Title(th, functions.GetGreetMsg()).Layout(gtx)
}

// Title creates a center aligned H1.
func Title(th *material.Theme, caption string) material.LabelStyle {
	l := material.H1(th, caption)
	l.Color = color.NRGBA{R: 127, G: 22, B: 100, A: 255}
	l.Alignment = text.Middle
	return l
}
