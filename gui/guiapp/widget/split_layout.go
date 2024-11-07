package widget

import (
	"gioui.org/layout"
	"gioui.org/op"
	"image"
)

type (
	SplitWidget struct {
	}
)

func (s SplitWidget) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	lSize := gtx.Constraints.Min.X / 2
	rSize := gtx.Constraints.Min.X - lSize
	// creating a code block with left widget own scope
	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(lSize, gtx.Constraints.Max.Y))
		left(gtx)
	}
	// right side widget code block
	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rSize, gtx.Constraints.Max.Y))
		trans := op.Offset(image.Pt(lSize, 0)).Push(gtx.Ops)
		right(gtx)
		trans.Pop()
	}
	return layout.Dimensions{Size: gtx.Constraints.Max}
}
