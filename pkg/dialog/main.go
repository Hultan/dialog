package dialog

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// iconPaths holds information about what iconType corresponds to what icon path
var iconPaths = map[iconType]string{
	iconInformation: iconInformationFilename,
	iconWarning:     iconWarningFilename,
	iconQuestion:    iconQuestionFilename,
	iconError:       iconErrorFilename,
}

// gtkButtons holds information about what buttonsType corresponds to which gtk buttons
var gtkButtons = map[buttonsType][][]interface{}{
	buttonsOk:          {{"Ok", gtk.RESPONSE_OK}},
	buttonsOkCancel:    {{"Ok", gtk.RESPONSE_OK}, {"Cancel", gtk.RESPONSE_CANCEL}},
	buttonsYesNo:       {{"Yes", gtk.RESPONSE_YES}, {"No", gtk.RESPONSE_NO}},
	buttonsYesNoCancel: {{"Yes", gtk.RESPONSE_YES}, {"No", gtk.RESPONSE_NO}, {"Cancel", gtk.RESPONSE_CANCEL}},
}

//
// Public methods
//

// Title is the starting method, since every dialog needs a title.
func Title(title string) *Dialog {
	return &Dialog{title: title}
}

//
// Dialog information & size
//

// Text sets the main text in the dialog.
func (d *Dialog) Text(text string) *Dialog {
	d.text = text
	return d
}

// TextMarkup sets the main text in the dialog in the GTK markup format.
func (d *Dialog) TextMarkup(textMarkup string) *Dialog {
	d.textMarkup = textMarkup
	return d
}

// Extra sets the extra text that will be displayed in a scrollable text box.
func (d *Dialog) Extra(extra string) *Dialog {
	d.extra = extra
	return d
}

// Size sets the minimum size of the dialog.
func (d *Dialog) Size(width, height int) *Dialog {
	d.width = width
	d.height = height
	return d
}

// Width sets the minimum width of the dialog.
func (d *Dialog) Width(width int) *Dialog {
	d.width = width
	return d
}

// Height sets the minimum height of the dialog.
func (d *Dialog) Height(height int) *Dialog {
	d.height = height
	return d
}

//
// Icons
//

// InfoIcon adds an information icon to the dialog
func (d *Dialog) InfoIcon() *Dialog {
	d.icon = iconInformation
	return d
}

// WarningIcon adds a warning icon to the dialog
func (d *Dialog) WarningIcon() *Dialog {
	d.icon = iconWarning
	return d
}

// QuestionIcon adds a question icon to the dialog
func (d *Dialog) QuestionIcon() *Dialog {
	d.icon = iconQuestion
	return d
}

// ErrorIcon adds an error icon to the dialog
func (d *Dialog) ErrorIcon() *Dialog {
	d.icon = iconError
	return d
}

//
// Buttons
//

// OkButton adds an ok button to the dialog
func (d *Dialog) OkButton() *Dialog {
	d.buttons = buttonsOk
	return d
}

// OkCancelButtons adds an ok button and a cancel button to the dialog
func (d *Dialog) OkCancelButtons() *Dialog {
	d.buttons = buttonsOkCancel
	return d
}

// YesNoButtons adds a yes button and no cancel button to the dialog
func (d *Dialog) YesNoButtons() *Dialog {
	d.buttons = buttonsYesNo
	return d
}

// YesNoCancelButtons adds a yes button, a no button, and a cancel button to the dialog
func (d *Dialog) YesNoCancelButtons() *Dialog {
	d.buttons = buttonsYesNoCancel
	return d
}

//
// Show
//

// Show will display the dialog
func (d *Dialog) Show() gtk.ResponseType {
	return d.createAndShowDialog()
}

//
// Private methods
//

func (d *Dialog) createAndShowDialog() gtk.ResponseType {
	dialog, err := d.createDialog()
	handleError(err)

	return d.showDialog(dialog)
}

func (d *Dialog) createDialog() (*gtk.Dialog, error) {
	dialog, err := gtk.DialogNewWithButtons(d.title, nil, gtk.DIALOG_MODAL, gtkButtons[d.buttons]...)
	handleError(err)

	content, err := dialog.GetContentArea()
	handleError(err)

	imageBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	handleError(err)

	imageBox.SetMarginTop(20)
	imageBox.SetMarginBottom(10)
	imageBox.SetMarginStart(20)
	imageBox.SetMarginEnd(20)
	content.Add(imageBox)

	if d.icon != iconNone {
		image, err := gtk.ImageNewFromFile(iconPaths[d.icon])
		handleError(err)

		imageBox.Add(image)
	}

	if d.textMarkup != "" {
		label, err := gtk.LabelNew("")
		label.SetMarkup(d.textMarkup)
		label.SetUseMarkup(true)
		handleError(err)

		imageBox.Add(label)
	} else if d.text != "" {
		label, err := gtk.LabelNew(d.text)
		handleError(err)

		imageBox.Add(label)
	}

	if d.extra != "" {
		scroll, err := gtk.ScrolledWindowNew(nil, nil)
		handleError(err)
		content.PackEnd(scroll, true, true, 20)

		buffer, err := gtk.TextBufferNew(nil)
		handleError(err)

		buffer.SetText(d.extra)
		extraTextView, err := gtk.TextViewNewWithBuffer(buffer)
		extraTextView.SetAcceptsTab(false)
		extraTextView.SetEditable(false)
		extraTextView.SetWrapMode(gtk.WRAP_WORD)
		extraTextView.SetMarginStart(20)
		extraTextView.SetMarginEnd(20)
		scroll.Add(extraTextView)
	}

	dialog.SetSizeRequest(d.width, d.height)
	dialog.ShowAll()
	return dialog, nil
}

func (d *Dialog) showDialog(dialog *gtk.Dialog) gtk.ResponseType {
	response := dialog.Run()
	dialog.Destroy()
	return response
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
