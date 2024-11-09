package gui

import (
	"fun-with-golang/gui/guiapp"
	"fun-with-golang/gui/guiapp/theme"
)

func Init() {
	//guiapp.Layout(greet.Layout)
	//guiapp.Layout(widget.DrawSplitVisual)
	guiapp.Layout(theme.ColorUpdateLayout)
	//guiapp.Layout(widget.ExampleSplit)
}
