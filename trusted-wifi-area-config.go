package main

import (
	_ "fmt"
	ui "github.com/VladimirMarkelov/clui"
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

func configFilePath() string {
	return homedir() + "/.trusted-wifi-area"
}

func homedir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadConfig() []string {
	content, _ := ioutil.ReadFile(configFilePath())
	data := strings.Split(string(content), "\n")
	return data
}

func saveConfig(data string) {
	err := ioutil.WriteFile(configFilePath(), []byte(data), 0644)
	check(err)
}

func wifiNames() []string {
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
		names = append(names, name)
	}
	return names
}

const wx = 80
const wy = 25

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

	_ = ui.CreateLabel(view, 20, 1, "Please select trusted Wi-fi networks:", 1)

	frmList := ui.CreateFrame(view, 30, 13, ui.BorderNone, ui.Fixed)
	frmList.SetGaps(1, ui.KeepValue)

	//frmChk := ui.CreateFrame(view, 8, 5, ui.BorderNone, ui.Fixed)
	frmList.SetPack(ui.Vertical)

	frmList.SetScrollable(true)

	configData := loadConfig()
	var checkboxes []*ui.CheckBox
	for _, name := range wifiNames() {
		checkBox := ui.CreateCheckBox(frmList, ui.AutoSize, name, ui.Fixed)
		checkboxes = append(checkboxes, checkBox)
		checkBox.SetActive(false)

		if stringInSlice(name, configData) {
			checkBox.SetState(1)
		}
	}
	//_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	//_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	//_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	//ee := ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	//_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)
	//_ = ui.CreateCheckBox(frmList, ui.AutoSize, "Use 222ListBox", ui.Fixed)

	frmBtn := ui.CreateFrame(view, 8, 4, ui.BorderNone, ui.Fixed)
	frmBtn.SetPack(ui.Horizontal)

	btnSave := ui.CreateButton(frmBtn, ui.AutoSize, 2, "Save & Exit", ui.Fixed)
	btnSave.SetActive(false)
	btnQuit := ui.CreateButton(frmBtn, ui.AutoSize, 2, "Discard", ui.Fixed)
	btnQuit.SetActive(false)

	//ee.SetActive(false)
	//view.SetActive(false)
	//ee.OnActive(func(active bool) {
	//fmt.Println(active)
	//})
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

		configData := ""
		for _, c := range checkboxes {
			if c.State() == 1 {
				configData += c.Title() + "\n"
			}
		}
		saveConfig(configData)
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

	mainLoop()
}
