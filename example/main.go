package main

import (
	_ "embed"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/dialog"
)

func main() {
	// Initialize gtk
	gtk.Init(&os.Args)

	dialog.Title("Hello World!").OkButton()
	dialog.Title("Hello World!").Text("How are you today?").OkButton()
	_, _ = dialog.Title("Hello World!").
		TextMarkup("How are you <i><b>today</b></i>?").
		InfoIcon().
		OkButton().
		Size(300, 150).
		Show()

	response, _ := dialog.Title("Hello World!").
		Text("How are you today?").
		Extra(getLongText()).
		QuestionIcon().
		YesNoButtons().
		Height(400).
		Show()

	if response == gtk.RESPONSE_YES {
		_, _ = dialog.Title("Your response...").
			Text("...was affirmative!").
			WarningIcon().
			OkButton().
			Width(400).
			Show()
	} else {
		_, _ = dialog.Title("Your response...").
			Text("...was very negative!").
			ErrorIcon().
			OkButton().
			Width(400).
			Show()
	}
}

func getLongText() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
}
