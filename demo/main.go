package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/component"
	"github.com/owenoclee/tfw-go/layout"
)

var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer eget nibh et risus porta semper id id diam. Nam blandit, arcu at eleifend facilisis, augue nunc pulvinar nisl, a finibus sapien orci quis mi. Donec malesuada, tellus vitae finibus ornare, nunc nibh consectetur risus, sed porttitor risus nisl a libero. Nullam odio lectus, pellentesque sed ipsum a, convallis sodales odio. Nam justo sapien, posuere eget felis at, accumsan varius augue. Integer pulvinar tempor consequat. Praesent in maximus massa. Fusce posuere sit amet ligula eu condimentum. Nullam porta augue maximus lectus euismod, porttitor placerat risus facilisis. Nunc interdum neque metus, quis suscipit dui pretium vitae."
var number = 0

func main() {
	exampleStyle := tcell.Style{}.Foreground(tcell.ColorBlue)
	exampleStyle2 := tcell.Style{}.Foreground(tcell.ColorOrange)
	leftBoxes := layout.NewHorizontalSplit(
		component.NewBox(component.PrettyBorder, component.NewWrappedTextWithStyle(loremIpsum, &exampleStyle)),
		component.NewBox(component.PrettyBorder, component.NewWrappedTextWithStyle(loremIpsum, &exampleStyle2)),
	)

	numberDisplay := component.NewWrappedText(fmt.Sprint(number))
	numberDisplayAdd := func(add int) func() {
		return func() {
			number += add
			numberDisplay.SetText(fmt.Sprint(number))
		}
	}
	middleBox := component.NewBox(component.PrettyBorder,
		layout.NewMargin(1, 1, 2, 2,
			layout.NewRows(2, 0,
				numberDisplay,
				component.NewShortcutOption('k', "Increase number", numberDisplayAdd(+1)),
				component.NewShortcutOption('j', "Decrease number", numberDisplayAdd(-1)),
			),
		),
	)

	rightBox := component.NewBox(component.PrettyBorder,
		layout.NewColumns(12, 0,
			component.NewWrappedText("column 1"),
			component.NewWrappedText("column 2"),
			component.NewWrappedText("column 3"),
		),
	)

	quit := make(chan struct{})
	app := tfw.NewApp(quit,
		layout.NewWithToolbar(1,
			component.NewBox(component.PrettyBorder,
				layout.NewVerticalSplit(
					leftBoxes,
					middleBox,
					rightBox,
				),
			),
			component.NewShortcutOption(
				'q',
				"quit",
				func() { close(quit) },
			),
		),
	)
	app.Run()
}
