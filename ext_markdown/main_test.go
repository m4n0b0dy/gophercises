package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_MakeUI(t *testing.T) {
	var testCfg config
	edit, preview := testCfg.makeUI()

	test.Type(edit, "hello")
	if preview.String() != "hello" {
		t.Errorf("Failed -- did not find value in preview")
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test Markdown")

	edit, preview := testCfg.makeUI()

	testCfg.CreateMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))
	test.Type(edit, "something")
	if preview.String() != "something" {
		t.Errorf("Failed -- did not find value in preview")
	}
}
