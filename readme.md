## DIALOG

A simple message dialog package. 

I don't recommend using this package, use **sqweeks** package instead (see links below). I am confident that my package contains lots of bugs.

### Usage:

This package needs GoTK3 (GTK) to run (see link below). 
```go
// Initialize gtk
gtk.Init(&os.Args)

// Super simple dialog
dialog.Title("Hello World!").OkButton().Show()

// Normal dialog
dialog.Title("Hello World!").Text("How are you today?").OkButton().Show()

// Question dialog
response := dialog.Title("Hello World!").
    Text("How are you today?").
    Extra(getLongText()).
    QuestionIcon().
    YesNoButtons().
    Height(400).
    Show()

if response == gtk.RESPONSE_YES {
    // Warning dialog
    dialog.Title("Your response...").
        Text("...was affirmative!").
        WarningIcon().
        OkButton().
        Width(400).
        Show()
} else {
    // Error dialog
    dialog.Title("Your response...").
        Text("...was very negative!").
        ErrorIcon().
        OkButton().
        Width(400).
        Show()
}
```
# LINKS
* https://github.com/sqweek/dialog
* https://github.com/gotk3/gotk3
# TODO
