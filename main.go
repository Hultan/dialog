package dialog

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

// Dialog contains information about the dialog the user wants
type Dialog struct {
	title            string
	text, textMarkup string
	headerColor      string
	width, height    int
	icon             iconType
	customIconPath   string
	buttons          buttonsType
	extra, extraName string
	extraHeight      int
	extraExpand      bool

	dialog *gtk.Dialog
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

var (
	colors map[iconType][4]float64
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
// To avoid getting the error: non-constant format string in call to (*github.com/hultan/dialog.Dialog).Text
// from go vet, the first argument is named message instead of format.
func Title(message string) *Dialog {
	return Titlef(message)
}

func Titlef(message string, a ...any) *Dialog {
	colors = make(map[iconType][4]float64, 6)
	colors[iconNone] = [4]float64{1, 1, 1, 1}
	colors[iconInformation] = [4]float64{1, 1, 1, 1}
	colors[iconWarning] = [4]float64{0.941, 0.729, 0.192, 1.0}
	colors[iconQuestion] = [4]float64{0.118, 0.69, 0.157, 1.0}
	colors[iconError] = [4]float64{0.941, 0.259, 0.192, 1.0}
	colors[iconCustom] = [4]float64{1, 1, 1, 1}

	return &Dialog{title: fmt.Sprintf(message, a...), width: 300, extraName: "Details"}
}

// Text sets the main text in the dialog.
// To avoid getting the error: non-constant format string in call to (*github.com/hultan/dialog.Dialog).Text
// from go vet, the first argument is named message instead of format.
func (d *Dialog) Text(message string) *Dialog {
	return d.Textf(message)
}

func (d *Dialog) Textf(message string, a ...any) *Dialog {
	d.text = fmt.Sprintf(message, a...)
	return d
}

// TextMarkup sets the main text in the dialog in the GTK markup format.
// To avoid getting the error: non-constant format string in call to (*github.com/hultan/dialog.Dialog).Text
// from go vet, the first argument is named message instead of format.
func (d *Dialog) TextMarkup(message string) *Dialog {
	return d.TextMarkupf(message)
}

func (d *Dialog) TextMarkupf(message string, a ...any) *Dialog {
	d.textMarkup = fmt.Sprintf(message, a...)
	return d
}

// HeaderColor sets the color of the header. The default value (if
// HeaderColor is not called) depends on the chosen icon:
//
// No icon: Gives a white header
// InfoIcon: Gives a white header
// WarningIcon: Gives an orange header
// ErrorIcon: Gives a red header
// QuestionIcon: Gives a green header
// CustomIcon: Gives a White header
func (d *Dialog) HeaderColor(color string) *Dialog {
	d.headerColor = color
	return d
}

// Size sets the minimum size of the dialog.
func (d *Dialog) Size(width, height int) *Dialog {
	d.width = width
	d.height = height
	return d
}

// Width sets the minimum width of the dialog.
// The default width is 300.
func (d *Dialog) Width(width int) *Dialog {
	d.width = width
	return d
}

// Height sets the minimum height of the dialog. The dialog will expand if the user expands the extra field (by ExtraHeight pixels).
func (d *Dialog) Height(height int) *Dialog {
	d.height = height
	return d
}

// Extra sets the extra text that will be displayed in a scrollable text box.
// To avoid getting the error: non-constant format string in call to (*github.com/hultan/dialog.Dialog).Text
// from go vet, the first argument is named message instead of format.
func (d *Dialog) Extra(message string) *Dialog {
	return d.Extraf(message)
}

func (d *Dialog) Extraf(message string, a ...any) *Dialog {
	d.extra = fmt.Sprintf(message, a...)
	return d
}

// ExtraExpand sets the extra text that will be displayed in a scrollable text box and expands it.
// To avoid getting the error: non-constant format string in call to (*github.com/hultan/dialog.Dialog).Text
// from go vet, the first argument is named message instead of format.
func (d *Dialog) ExtraExpand(message string) *Dialog {
	d.extraExpand = true
	return d.ExtraExpandf(message)
}

func (d *Dialog) ExtraExpandf(message string, a ...any) *Dialog {
	d.extra = fmt.Sprintf(message, a...)
	d.extraExpand = true
	return d
}

// ExtraName sets the name of the expander that shows the extra text.
// The default name is "Details".
func (d *Dialog) ExtraName(extraName string) *Dialog {
	d.extraName = extraName
	return d
}

// ExtraHeight sets the height of the extra field when it is expanded.
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
	d.customIconPath = path
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

	response := dialog.Run()
	dialog.Destroy()
	return response, nil
}

//
// Private methods
//

func (d *Dialog) createDialog() (*gtk.Dialog, error) {
	dialog, err := gtk.DialogNewWithButtons(d.title, nil, gtk.DIALOG_MODAL, gtkButtons[d.buttons]...)
	if err != nil {
		return nil, err
	}
	d.dialog = dialog

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

	drawingArea, err := d.getDrawingArea()
	if err != nil {
		return nil, err
	}

	label, err := d.getLabel(d.icon != iconNone)
	if err != nil {
		return nil, err
	}

	// Add widgets to the overlay
	if drawingArea != nil {
		overlay.Add(drawingArea) // Image is the base layer
	}
	overlay.AddOverlay(label) // Label goes on top
	overlay.SetSizeRequest(d.width, 50)

	if d.extra != "" {
		expander, err := d.getExtraExpander()
		if err != nil {
			return nil, err
		}

		if d.extraExpand {
			expander.SetExpanded(true)
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
	expander, err := gtk.ExpanderNew(d.extraName)
	if err != nil {
		return nil, err
	}
	expander.SetVExpand(true)
	expander.SetHExpand(true)

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
	extraTextView.SetHExpand(true)
	extraTextView.SetVExpand(true)
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
		label.SetName("headerLabel") // Set a name for CSS targeting

		// Apply CSS styling
		err := applyCSS(`#headerLabel { color: black; }`)
		if err != nil {
			log.Fatal("Failed to apply CSS:", err)
		}
	}

	// Set margins AFTER SetHExpand and SetVExpand
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_CENTER)
	label.SetHExpand(true)
	label.SetVExpand(false)
	label.SetLineWrap(true)
	label.SetLineWrapMode(pango.WRAP_WORD_CHAR)

	if hasImage {
		label.SetMarginStart(45)
	} else {
		label.SetMarginStart(10)
	}

	return label, nil
}

func (d *Dialog) getDrawingArea() (*gtk.DrawingArea, error) {
	// Create a DrawingArea
	drawingArea, _ := gtk.DrawingAreaNew()
	drawingArea.SetSizeRequest(d.width, 50) // Set control size
	drawingArea.SetHExpand(true)

	// Connect the "draw" signal to render content
	drawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		d.renderIconAndBackground(cr)
	})

	return drawingArea, nil
}

// renderIconAndBackground renders a background + PNG icon
func (d *Dialog) renderIconAndBackground(cr *cairo.Context) {
	var pic *gdk.Pixbuf
	var col = colors[d.icon]

	if d.headerColor != "" {
		col = d.getHeaderColor()
	}

	// Set the background color (light blue)
	width, _ := d.dialog.GetSize()
	cr.SetSourceRGBA(col[0], col[1], col[2], col[3])
	cr.Rectangle(0, 0, float64(width), 50)
	cr.Fill()

	switch d.icon {
	case iconInformation:
		pic, _ = gdk.PixbufNewFromBytesOnly(infoIcon)
	case iconWarning:
		pic, _ = gdk.PixbufNewFromBytesOnly(warningIcon)
	case iconQuestion:
		pic, _ = gdk.PixbufNewFromBytesOnly(questionIcon)
	case iconError:
		pic, _ = gdk.PixbufNewFromBytesOnly(errorIcon)
	case iconCustom:
		// TODO : Cache this image
		pic, _ = gdk.PixbufNewFromFile(d.customIconPath)
	default:
		return
	}

	// Render Pixbuf onto Cairo surface
	surface, _ := gdk.CairoSurfaceCreateFromPixbuf(pic, 0, nil)
	if surface == nil {
		log.Fatal("Failed to convert Pixbuf to Cairo surface")
	}

	// Draw image at position (9, 9)
	cr.SetSourceSurface(surface, 9, 9)
	cr.Paint()
}

func (d *Dialog) getHeaderColor() [4]float64 {
	var color [4]float64
	if len(d.headerColor) != 9 && len(d.headerColor) != 8 {
		log.Fatal(fmt.Errorf("invalid color string length: %s", d.headerColor))
	}
	if d.headerColor[0] == '#' {
		d.headerColor = d.headerColor[1:]
	}
	for i := 0; i < 4; i++ {
		hex, err := strconv.ParseUint(d.headerColor[i*2:i*2+2], 16, 8)
		if err != nil {
			log.Fatal(fmt.Errorf("invalid color string: %s", d.headerColor))
		}
		color[i] = float64(hex) / 255.0
	}
	return color
}

// Apply CSS to GTK widgets
func applyCSS(css string) error {
	// Create a CSS provider
	provider, err := gtk.CssProviderNew()
	if err != nil {
		return err
	}

	err = provider.LoadFromData(css)
	if err != nil {
		return err
	}

	// Apply CSS to the default screen
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		return err
	}

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	return nil
}
