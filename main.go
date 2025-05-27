package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init_folder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func write_json(file_path string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func read_json(file_path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InterfaceSliceToStringSlice(interfaceSlice []interface{}) ([]string, error) {
	stringSlice := make([]string, len(interfaceSlice))
	for i, v := range interfaceSlice {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("%d not string", i)
		}
		stringSlice[i] = str
	}
	return stringSlice, nil
}

func get_vm_list() []string {
	data, err := read_json("./data/config/config.json")
	if err != nil {
		return []string{}
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_list
}

func main() {

	// init folder
	init_folder("./data")
	init_folder("./data/config")
	init_folder("./data/vms")

	// init config file
	if _, err := os.Stat("./data/config/config.json"); os.IsNotExist(err) {
		write_json("./data/config/config.json", map[string]interface{}{
			"vm_list": []string{},
			"vm_uuid": map[string]string{},
		})
	}

	// init window
	myApp := app.New()
	myWindow := myApp.NewWindow("QEMU GUI")
	myWindow.Resize(fyne.NewSize(0, 600))

	// vm list
	vmList := container.NewVBox()
	vm_list_refresh := func() {
		vm_list := get_vm_list()
		vmList.RemoveAll()
		if len(vm_list) > 0 {
			for _, vm_name := range vm_list {
				vmList.Add(widget.NewLabel(vm_name))
			}
		} else {
			vmList.Add(widget.NewLabel("Click New to create new VM."))
		}
	}
	vm_list_refresh()

	// top buttons
	topButtons := container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),

			// new vm
			widget.NewButtonWithIcon("New", theme.DocumentCreateIcon(), func() {
				fmt.Println("new vm")
			}),

			// refresh vm list
			widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
				vm_list_refresh()
			}),

			// settings
			widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
				fmt.Println("settings")
			}),

			// help
			widget.NewButtonWithIcon("Help", theme.HelpIcon(), func() {
				fmt.Println("help")
			}),

			// about
			widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
				fmt.Println("about")
			}),

			// exit
			widget.NewButtonWithIcon("Exit", theme.CancelIcon(), func() {
				myApp.Quit()
			}),

			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	// show window
	myWindow.SetContent(container.NewBorder(
		topButtons,                   // top
		nil,                          // bottom
		nil,                          // left
		nil,                          // right
		container.NewVScroll(vmList), // content
	))
	myWindow.ShowAndRun()
}
