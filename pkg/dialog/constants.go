package dialog

// Dialog contains information about the dialog the user wants
type Dialog struct {
	title, text, extra string
	textMarkup         string
	width, height      int
	icon               iconType
	buttons            buttonsType
}

// iconType describes what type of icon the user wants
type iconType int

const (
	iconNone iconType = iota
	iconInformation
	iconWarning
	iconQuestion
	iconError
)

// buttonsType describes the number of different buttons the user wants
type buttonsType int

const (
	buttonsOk buttonsType = iota
	buttonsOkCancel
	buttonsYesNo
	buttonsYesNoCancel
)

// Icon paths
const (
	iconInformationFilename = "/home/per/code/dialog/assets/info.png"
	iconWarningFilename     = "/home/per/code/dialog/assets/warning.png"
	iconQuestionFilename    = "/home/per/code/dialog/assets/question.png"
	iconErrorFilename       = "/home/per/code/dialog/assets/error.png"
)
