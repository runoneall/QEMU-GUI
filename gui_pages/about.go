package gui_pages

import (
	"fmt"
	"qemu-gui/helper"
	"qemu-gui/vars"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func About_Page(myApp fyne.App) {

	// about window
	aboutWindow := myApp.NewWindow("About")
	aboutWindow.Resize(fyne.NewSize(400, 300))

	// right area
	aboutRight := container.NewVBox(
		widget.NewLabel("Click Left Button To Use"),
	)

	// left button
	aboutLeft := container.NewVBox(
		widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
			aboutRight.RemoveAll()
			aboutRight.Add(widget.NewLabel("关于"))
		}),
		widget.NewButtonWithIcon("", theme.ComputerIcon(), func() {
			aboutRight.RemoveAll()
			aboutRight.Add(widget.NewLabelWithStyle(
				"QEMU Excutable Check",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			))

			// check qemu executable
			go func() {
				for _, qemu_system := range vars.QEMU_SYSTEMS {
					status, output := helper.Excutable_Command(qemu_system + " --version")
					if status {

						// find version
						re := regexp.MustCompile(`QEMU emulator version (\d+\.\d+\.\d+)`)
						matches := re.FindStringSubmatch(output)

						// if version not found
						if len(matches) < 2 {
							aboutRight.Add(widget.NewLabel(
								fmt.Sprintf("%s is not installed or version not found", qemu_system),
							))
							return
						}

						// show version
						aboutRight.Add(widget.NewLabel(
							fmt.Sprintf("%s version: %s", qemu_system, matches[1]),
						))

					} else {
						aboutRight.Add(widget.NewLabel(
							fmt.Sprintf("%s is not installed or not found", qemu_system),
						))
					}
				}
			}()

		}),
	)

	// show window
	aboutWindow.SetContent(container.NewHBox(
		aboutLeft,
		container.NewVScroll(aboutRight),
	))
	aboutWindow.Show()

}
