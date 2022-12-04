package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"strings"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

// fyne package -appVersion 1.0.0 -name MarkDown -release --appID io.fyne.terminal

var cfg config

func main() {
	// create fyne app
	a := app.New()

	// create a window for app
	win := a.NewWindow("Markdown")

	// get user interaface
	edit, preview := cfg.makeUI()
	cfg.CreateMenuItems(win)

	// set the current content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (c *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	c.EditWidget = edit
	c.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown // update as people type using Parse Markdown function

	return edit, preview
}

func (c *config) CreateMenuItems(win fyne.Window) {
	openMenu := fyne.NewMenuItem("Open...", c.openMenuFunc(win))
	saveMenu := fyne.NewMenuItem("Save...", c.saveFunc(win))
	c.SaveMenuItem = saveMenu
	c.SaveMenuItem.Disabled = true

	saveAsMenu := fyne.NewMenuItem("Save As...", c.saveAsFunc(win))
	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (c *config) saveFunc(win fyne.Window) func() {
	return func() {
		if c.CurrentFile != nil {
			write, err := storage.Writer(c.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(c.EditWidget.Text))

			defer write.Close()
		}
	}
}

func (c *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil {
				//user cancelled
				return
			}
			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please Name as .md", win)
				return
			}
			//save file
			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (c *config) openMenuFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				return
			}
			defer read.Close()

			data, err := ioutil.ReadAll(read)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			c.EditWidget.SetText(string(data))

			c.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			c.SaveMenuItem.Disabled = false

		}, win)
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}
