package gui

import (
	"fun-with-golang/gui/guiapp"
	"fun-with-golang/gui/guiapp/widget"
)

func Init() {
	//guiapp.Layout(greet.Layout)
	//guiapp.Layout(widget.DrawSplitVisual)
	guiapp.Layout(widget.ExampleSplit)
}
