package theme

import (
	"fun-with-golang/ch1/functions"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

// Theme should be part of the application loop

var isChecked widget.Bool

func ColorUpdateLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var checkboxLabel string
	isChecked.Update(gtx)
	updatedThemeBackground := th
	if isChecked.Value {
		checkboxLabel = "checked"
		updatedThemeBackground.Fg = color.NRGBA{R: 121, G: 0, B: 0, A: 255}
		updatedThemeBackground.Bg = color.NRGBA{B: 0, G: 0, R: 255, A: 255}
	} else {
		checkboxLabel = "not-checked"
		updatedThemeBackground.Fg = color.NRGBA{R: 0, G: 200, B: 0, A: 255}
		updatedThemeBackground.Bg = color.NRGBA{B: 0, G: 255, R: 255, A: 255}
	}
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(material.H3(updatedThemeBackground, functions.GetGreetMsg()).Layout),
		layout.Rigid(material.CheckBox(updatedThemeBackground, &isChecked, checkboxLabel).Layout),
	)
	//return layout.Flex{
	//	Axis: layout.Vertical,
	//}.Layout(gtx,
	//	layout.Rigid(material.H3(updatedThemeBackground, "Hello, World!").Layout),
	//	layout.Rigid(material.CheckBox(updatedThemeBackground, &isChecked, checkboxLabel).Layout),
	//)
}
