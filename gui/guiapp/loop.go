package guiapp

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
)

// Render is a utility to start a rendering Gio app.
func Render(fn func(gtx layout.Context)) {
	go func() {
		w := new(app.Window)

		w.Run(func() {
			gtx := layout.Context{}
			fn(gtx)
			w.Invalidate()
		})
	}()
	app.Main()
}

// Input is a utility to start a rendering and input Gio app.
func Input(fn func(gtx layout.Context)) {
	go func() {
		w := new(app.Window)

		w.Run(func() {
			gtx := layout.Context{}
			fn(gtx)
			w.Invalidate()
		})
	}()
	app.Main()
}

// InputSize is a utility to start a rendering and input Gio app.
func InputSize(fn func(gtx layout.Context, windowSize image.Point)) {
	go func() {
		w := new(app.Window)

		w.Run(func() {
			gtx := layout.Context{}
			fn(gtx, gtx.Constraints.Max)
			w.Invalidate()
		})
	}()
	app.Main()
}

// Metric is a utility to start a rendering, input and metrics Gio app.
func Metric(fn func(gtx layout.Context)) {
	go func() {
		w := new(app.Window)

		w.Run(func() {
			gtx := layout.Context{}
			fn(gtx)
			w.Invalidate()
		})
	}()
	app.Main()
}

// Layout is a utility to start a layouting Gio app.
func Layout(draw func(gtx layout.Context, th *material.Theme) layout.Dimensions) {
	go func() {
		th := material.NewTheme()
		th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
		w := new(app.Window)
		w.Option(app.Title("fun-with-golang"))
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				draw(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
