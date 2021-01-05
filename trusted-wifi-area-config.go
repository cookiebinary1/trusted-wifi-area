package main

import (
	"fmt"
	_ "fmt"
	ui "github.com/VladimirMarkelov/clui"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func updateProgress(value string, pb *ui.ProgressBar) {
	v, _ := strconv.Atoi(value)
	pb.SetValue(v)
}

func changeTheme(lb *ui.ListBox, btn *ui.Button, tp int) {
	items := ui.ThemeNames()
	dlgType := ui.SelectDialogRadio
	if tp == 1 {
		dlgType = ui.SelectDialogList
	}

	curr := -1
	for i, tName := range items {
		if tName == ui.CurrentTheme() {
			curr = i
			break
		}
	}

	selDlg := ui.CreateSelectDialog("Choose a theme", items, curr, dlgType)
	selDlg.OnClose(func() {
		switch selDlg.Result() {
		case ui.DialogButton1:
			idx := selDlg.Value()
			lb.AddItem(fmt.Sprintf("Selected item: %v", selDlg.Value()))
			lb.SelectItem(lb.ItemCount() - 1)
			if idx != -1 {
				ui.SetCurrentTheme(items[idx])
			}
		}

		btn.SetEnabled(true)
		// ask the composer to repaint all windows
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})

}

const wx = 80
const wy = 20

func createView() {
	//ui.SetCurrentTheme("turbovision")

	sx, sy := ui.ScreenSize()

	view := ui.AddWindow(int((sx-wx)/2), int((sy-wy)/2), wx, wy, "Trusted Wi-Fi Area")
	view.SetPack(ui.Vertical)
	view.SetPaddings(10, 2)
	//frmLeft := ui.CreateFrame(view, 8, 4, ui.BorderNone, 1)
	//frmLeft.SetPack(ui.Vertical)
	view.SetGaps(ui.KeepValue, 1)
	//frmLeft.SetPaddings(1, 1)

	frmList := ui.CreateFrame(view, 8, 1, ui.BorderNone, ui.Fixed)
	frmList.SetGaps(1, ui.KeepValue)

	//frmChk := ui.CreateFrame(view, 8, 5, ui.BorderNone, ui.Fixed)
	frmList.SetPack(ui.Vertical)

	checkBox := ui.CreateCheckBox(frmList, ui.AutoSize, "Use ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)

	frmBtn := ui.CreateFrame(view, 8, 5, ui.BorderNone, ui.Fixed)
	frmBtn.SetPack(ui.Horizontal)

	btnSave := ui.CreateButton(frmBtn, ui.AutoSize, 2, "Save & Exit", ui.Fixed)
	btnQuit := ui.CreateButton(frmBtn, ui.AutoSize, 3, "Discard", ui.Fixed)

	_ = checkBox.State()

	btnSave.Active()
	//btnTheme := ui.CreateButton(frmList, ui.AutoSize, 4, "Select theme", ui.Fixed)

	//logBox := ui.CreateListBox(view, 28, 5, ui.Fixed)

	//btnTheme.OnClick(func(ev ui.Event) {
	//	btnTheme.SetEnabled(false)
	//	tp := checkBox.State()
	//	changeTheme(logBox, btnTheme, tp)
	//})
	//btnSet.OnClick(func(ev ui.Event) {
	//	v := edit.Title()
	//	logBox.AddItem(fmt.Sprintf("New ProgressBar value: %v", v))
	//	logBox.SelectItem(logBox.ItemCount() - 1)
	//	updateProgress(v, pb)
	//})
	//btnStep.OnClick(func(ev ui.Event) {
	//	go pb.Step()
	//	logBox.AddItem("ProgressBar step")
	//	logBox.SelectItem(logBox.ItemCount() - 1)
	//	ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	//})
	btnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	btnSave.OnClick(func(event ui.Event) {
		// save
		go ui.Stop()
	})

	//
	//btnTheme.SetEnabled(false)
	//tp := checkBox.State()
	//changeTheme(logBox, btnTheme, tp)
	//items := ui.ThemeNames()
	//fmt.Println(items)

}

func mainLoop() {
	// Every application must create a single Composer and
	// call its initialize method
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	ui.SetThemePath("themes")

	createView()

	// start event processing loop - the main core of the library
	ui.MainLoop()
}

func main() {

	dir := "/etc/NetworkManager/system-connections/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var names []string

	for _, f := range files {
		name := f.Name()
		name = strings.ReplaceAll(name, "Auto ", "")
		name = strings.ReplaceAll(name, ".nmconnection", "")
		fmt.Println(name)
		names = append(names, name)
	}

	mainLoop()
}
