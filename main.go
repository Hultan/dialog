package dialog

import (
	_ "embed"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

// Dialog contains information about the dialog the user wants
type Dialog struct {
	title, text, extra string
	textMarkup         string
	width, height      int
	extraHeight        int
	icon               iconType
	buttons            buttonsType
	path               string
}

// iconType describes what type of icon the user wants
type iconType int

const (
	iconNone iconType = iota
	iconInformation
	iconWarning
	iconQuestion
	iconError
	iconCustom
)

// buttonsType describes the number of different buttons the user wants
type buttonsType int

const (
	buttonsOk buttonsType = iota
	buttonsOkCancel
	buttonsYesNo
	buttonsYesNoCancel
)

//go:embed error.png
var errorIcon []byte

//go:embed warning.png
var warningIcon []byte

//go:embed info.png
var infoIcon []byte

//go:embed question.png
var questionIcon []byte

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

// Title is the starting method (constructor), since every dialog needs a title.
func Title(title string) *Dialog {
	return &Dialog{title: title, width: 300}
}

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

// Height sets the minimum height of the dialog. The dialog will expand if the user expands the extra field (by ExtraHeight pixels).
func (d *Dialog) Height(height int) *Dialog {
	d.height = height
	return d
}

// ExtraHeight sets the height of the extra field, when it is expanded.
func (d *Dialog) ExtraHeight(extraHeight int) *Dialog {
	d.extraHeight = extraHeight
	return d
}

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

// CustomIcon adds a custom icon to the dialog
func (d *Dialog) CustomIcon(path string) *Dialog {
	d.icon = iconCustom
	d.path = path
	return d
}

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

// Show will display the dialog
func (d *Dialog) Show() (gtk.ResponseType, error) {
	dialog, err := d.createDialog()
	if err != nil {
		return gtk.RESPONSE_REJECT, err
	}

	return d.showDialog(dialog), err
}

//
// Private methods
//

func (d *Dialog) createDialog() (*gtk.Dialog, error) {
	dialog, err := gtk.DialogNewWithButtons(d.title, nil, gtk.DIALOG_MODAL, gtkButtons[d.buttons]...)
	if err != nil {
		return nil, err
	}

	content, err := dialog.GetContentArea()
	if err != nil {
		return nil, err
	}

	// Create an Overlay (for stacking widgets)
	overlay, err := gtk.OverlayNew()
	if err != nil {
		return nil, err
	}
	content.Add(overlay)

	image, err := d.getImage()
	if err != nil {
		return nil, err
	}

	label, err := d.getLabel(image != nil)
	if err != nil {
		return nil, err
	}

	// Add widgets to the overlay
	if image != nil {
		overlay.Add(image) // Image is the base layer
	}
	overlay.AddOverlay(label) // Label goes on top
	overlay.SetSizeRequest(d.width, 50)

	if d.extra != "" {
		expander, err := d.getExtraExpander()
		if err != nil {
			return nil, err
		}

		// Adjust window height dynamically when expanding/collapsing the expander
		expander.Connect("notify::expanded", func() {
			if expander.GetExpanded() {
				dialog.Resize(d.width, d.height+d.extraHeight) // Expand height
			} else {
				dialog.Resize(d.width, d.height) // Shrink height
			}
		})

		content.PackEnd(expander, true, true, 5)
	}

	dialog.SetSizeRequest(d.width, d.height)
	dialog.ShowAll()

	return dialog, nil
}

func (d *Dialog) getExtraExpander() (*gtk.Expander, error) {
	expander, err := gtk.ExpanderNew("Extra information")
	if err != nil {
		return nil, err
	}

	scroll, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	// Height for the expanded content
	scroll.SetSizeRequest(d.width, d.extraHeight)
	expander.Add(scroll)

	buffer, err := gtk.TextBufferNew(nil)
	if err != nil {
		return nil, err
	}

	buffer.SetText(d.extra)
	extraTextView, err := gtk.TextViewNewWithBuffer(buffer)
	if err != nil {
		return nil, err
	}
	extraTextView.SetAcceptsTab(false)
	extraTextView.SetEditable(false)
	extraTextView.SetWrapMode(gtk.WRAP_WORD)
	extraTextView.SetMarginStart(20)
	extraTextView.SetMarginEnd(20)
	scroll.Add(extraTextView)

	return expander, nil
}

func (d *Dialog) getLabel(hasImage bool) (*gtk.Label, error) {
	label, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	if d.textMarkup != "" {
		label.SetMarkup(d.textMarkup)
		label.SetUseMarkup(true)
	} else if d.text != "" {
		label.SetText(d.text)
	}
	if hasImage {
		label.SetMarginStart(45)
	} else {
		label.SetMarginStart(10)
	}
	label.SetLineWrapMode(pango.WRAP_WORD)
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVExpand(true)
	return label, nil
}

func (d *Dialog) getImage() (*gtk.Image, error) {
	var pic *gdk.Pixbuf
	var img *gtk.Image
	var err error

	switch d.icon {
	case iconError:
		pic, err = gdk.PixbufNewFromBytesOnly(errorIcon)
	case iconInformation:
		pic, err = gdk.PixbufNewFromBytesOnly(infoIcon)
	case iconQuestion:
		pic, err = gdk.PixbufNewFromBytesOnly(questionIcon)
	case iconWarning:
		pic, err = gdk.PixbufNewFromBytesOnly(warningIcon)
	case iconCustom:
		pic, err = gdk.PixbufNewFromFile(d.path)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	img, err = gtk.ImageNewFromPixbuf(pic)
	if err != nil {
		return nil, err
	}
	img.SetHAlign(gtk.ALIGN_START)
	img.SetMarginStart(5)
	img.SetMarginTop(5)
	img.SetMarginBottom(5)

	return img, err
}

func (d *Dialog) showDialog(dialog *gtk.Dialog) gtk.ResponseType {
	response := dialog.Run()
	dialog.Destroy()
	return response
}
