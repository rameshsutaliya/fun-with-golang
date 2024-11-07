package guiapp

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"image"
	"log"
	"os"
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
func Layout(lay func(gtx layout.Context) layout.Dimensions) {
	go func() {
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
				lay(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
