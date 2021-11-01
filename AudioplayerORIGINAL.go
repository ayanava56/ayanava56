package main

//Necessary imports for the program
import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"         //Need to install via go get
	"github.com/faiface/beep/mp3"     //Need to install via go get
	"github.com/faiface/beep/speaker" //Need to install via go get
)

//Declaring the variables
var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false //Creating So called flag variable and giving it False as input

func main() {
	go func(msg string) {
		fmt.Println(msg)
		if streamer == nil {
		} else {

			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("App Online...")
	time.Sleep(time.Second)
	a := app.New()
	w := a.NewWindow("audio player...")
	w.Resize(fyne.NewSize(200, 400))
	//Taking Path for  the background image for the app
	logo := canvas.NewImageFromFile("C:\\Users\\User\\OneDrive\\Desktop\\pepcoding proj\\Audio.jpg")
	logo.FillMode = canvas.ImageFillOriginal //Letting the original size of the pic to set the background
	//Creating Toolbar or Navbar for Play, Pause and Stop Buttons
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		//Play Btn
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		//Pause Btn
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		//Stop Btn
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()

		}),
		widget.NewToolbarSpacer(),
	)
	label := widget.NewLabel("Audio MP3")
	label.Alignment = fyne.TextAlignCenter //Aligning the texts bars in the centre of the screen
	label2 := widget.NewLabel("Play")
	label2.Alignment = fyne.TextAlignCenter
	browse_files := widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc) //Ignoring the error by using underscore
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"})) //String Slicing for recognizing .mp3 Files
	})

	c := container.NewVBox(label, browse_files, label2, toolbar) // vbox container
	w.SetContent(
		container.NewBorder(logo, nil, nil, nil, c),
	)
	w.ShowAndRun()
}
