package widget

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"image/color"
)

var split *Split // Declaration

func init() {
	// Initialize globally
	split = &Split{
		Ratio: 0,  // default ratio (centered)
		Bar:   10, // default bar width
	}
}

func ExampleSplit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return split.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", color.NRGBA{B: 0, G: 0, R: 255, A: 255})
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", color.NRGBA{B: 255, G: 0, R: 0, A: 255})
	})
}
