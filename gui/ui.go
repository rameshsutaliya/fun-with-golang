package gui

import (
	"fun-with-golang/gui/guiapp"
	"fun-with-golang/gui/guiapp/widget"
)

func Init() {
	//guiapp.Layout(greet.Layout)
	guiapp.Layout(widget.DrawSplitVisual)
	//guiapp.Layout(theme.ColorUpdateLayout)
	//guiapp.Layout(widget.ExampleSplit)
	//guiapp.Layout(gio_example.Counter{}.Layout)
	//counter.CounterExample()
	//temperature.TempConvertor()
	//timer.TimerView()
	//tabs.TabsView()
	//galaxy.View()
}
