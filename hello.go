package main

import (
	"fmt"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Video Recorder")
	w.Resize(fyne.NewSize(720, 720))
	displayMessage := widget.NewLabel("")
	displayMessage.TextStyle.Bold = true
	displayMessage.TextStyle.Italic = true

	startButton := widget.NewButtonWithIcon("Start", theme.MediaRecordIcon(), func() {

		filename := fmt.Sprintf("output_%s.mp4", time.Now().Format("20060102150405"))

		cmd := exec.Command("ffmpeg", "-f", "x11grab", "-video_size", "1920x1080", "-framerate", "30", "-i", ":1", filename)

		err := cmd.Start()
		if err != nil {
			fmt.Println("Error while starting recording:", err)
			return
		}

		fmt.Println("Screen recording started!")
		w.SetTitle("recording started")
		displayMessage.SetText("Recording Started")

	})

	stopButton := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
		err := exec.Command("pkill", "ffmpeg").Run()
		if err != nil {
			fmt.Println("Error while stopping recording:", err)
			return
		}

		fmt.Println("Screen recording stopped!")
		w.SetTitle("recording stopped")
		displayMessage.SetText("Recording Stopped")

	})

	buttonsBox := container.NewVBox(startButton, stopButton)
	content := container.NewVBox(
		buttonsBox,
		displayMessage,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
