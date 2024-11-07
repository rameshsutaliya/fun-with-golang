package widget

import (
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"image"
)

type (
	Split struct {
		// Ratio keeps the current layout.
		// 0 is center, -1 completely to the left, 1 completely to the right.
		Ratio float32
		// Bar is the width for resizing the layout
		Bar unit.Dp

		drag   bool
		dragID pointer.ID
		dragX  float32
	}
)

const defaultBarWidth = unit.Dp(10)

func (s *Split) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	// START BAR OMIT
	bar := gtx.Dp(s.Bar)
	if bar <= 1 {
		bar = gtx.Dp(defaultBarWidth)
	}

	proportion := (s.Ratio + 1) / 2
	leftSize := int(proportion*float32(gtx.Constraints.Max.X) - float32(bar))

	rightOffset := leftSize + bar
	rightsize := gtx.Constraints.Max.X - rightOffset
	// END BAR OMIT

	{ // handle input
		// START INPUTCODE OMIT
		barRect := image.Rect(leftSize, 0, rightOffset, gtx.Constraints.Max.X)
		area := clip.Rect(barRect).Push(gtx.Ops)

		// register for input
		event.Op(gtx.Ops, s)
		pointer.CursorColResize.Add(gtx.Ops)

		for {
			ev, ok := gtx.Event(pointer.Filter{
				Target: s,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release | pointer.Cancel,
			})
			if !ok {
				break
			}

			e, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch e.Kind {
			case pointer.Press:
				if s.drag {
					break
				}

				s.dragID = e.PointerID
				s.dragX = e.Position.X
				s.drag = true

			case pointer.Drag:
				if s.dragID != e.PointerID {
					break
				}

				deltaX := e.Position.X - s.dragX
				s.dragX = e.Position.X

				deltaRatio := deltaX * 2 / float32(gtx.Constraints.Max.X)
				s.Ratio += deltaRatio

				if e.Priority < pointer.Grabbed {
					gtx.Execute(pointer.GrabCmd{
						Tag: s,
						ID:  s.dragID,
					})
				}

			case pointer.Release:
				fallthrough
			case pointer.Cancel:
				s.drag = false
			}
		}

		area.Pop()
		// END INPUTCODE OMIT
	}

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftSize, gtx.Constraints.Max.Y))
		left(gtx)
	}

	{
		off := op.Offset(image.Pt(rightOffset, 0)).Push(gtx.Ops)
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		right(gtx)
		off.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
