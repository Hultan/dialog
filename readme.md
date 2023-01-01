## DIALOG

A simple message dialog package. 

### Usage:

This package needs GoTK3 (GTK) to run (see link below), and you will need to run gtk.Init() somewhere in your application, before using the dialogs. 
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
* Source of inspiration: https://github.com/sqweek/dialog
* GoTK3 : https://github.com/gotk3/gotk3
# TODO
* Choose between smaller and larger icons

