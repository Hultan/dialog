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

	_, _ = dialog.Title("%d custom icons!", 10).
		Text("This is a custom icon, really? This is a really long text that needs %d line breaks.", 5).
		ExtraExpand(getLongText()).
		ExtraHeight(200).
		CustomIcon("./example/assets/armour.png").
		HeaderColor("#6879D0FF").
		OkButton().
		Size(300, 50).
		Show()

	_, _ = dialog.Title("Hello World!").
		TextMarkup("<span foreground=\"black\">How are you on this <i><b>%s</b></i>?</span>", "Tuesday").
		InfoIcon().
		OkButton().
		Size(300, 100).
		Show()

	_, _ = dialog.Title("No image dialog!").
		Text("How are you today?").
		Extra(getLongText()).
		ExtraHeight(50).
		ExtraName("Extra name test").
		OkButton().
		Height(100).
		Show()

	response, _ := dialog.Title("Hello World!").
		Text("How are you today?").
		Extra(getLongText()).
		ExtraHeight(50).
		QuestionIcon().
		YesNoButtons().
		Height(125).
		Show()

	if response == gtk.RESPONSE_YES {
		_, _ = dialog.Title("Your response...").
			Text("...was affirmative!").
			WarningIcon().
			OkButton().
			Show()
	} else {
		_, _ = dialog.Title("Your response...").
			Text("...was very negative!").
			ErrorIcon().
			OkButton().
			Show()
	}
}

func getLongText() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
}
