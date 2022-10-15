package main

import (
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/dialog/pkg/dialog"
)

const (
	ApplicationId    = "se.softteam.softimdb"
	ApplicationFlags = glib.APPLICATION_FLAGS_NONE
)

func main() {
	// Initialize gtk
	gtk.Init(&os.Args)

	dialog.Title("Hello World!").OkButton()
	dialog.Title("Hello World!").Text("How are you today?").OkButton()
	dialog.Title("Hello World!").
		TextMarkup("How are you <i><b>today</b></i>?").
		InfoIcon().
		OkButton().
		Size(300, 150).
		Show()

	response := dialog.Title("Hello World!").
		Text("How are you today?").
		Extra(getLongText()).
		QuestionIcon().
		YesNoButtons().
		Height(400).
		Show()

	if response == gtk.RESPONSE_YES {
		dialog.Title("Your response...").
			Text("...was affirmative!").
			WarningIcon().
			OkButton().
			Width(400).
			Show()
	} else {
		dialog.Title("Your response...").
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
