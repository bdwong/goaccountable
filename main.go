package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"./icon"
	"github.com/getlantern/systray"
	// "github.com/getlantern/systray/example/icon"
	// "github.com/skratchdot/open-golang/open"
)

func main() {
	onExit := func() {
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("GoAccountable1")                // Required to register menu
	systray.SetTooltip("What did you commit to do?1") // init tooltip
	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("GoAccountable")
		systray.SetTooltip("What did you commit to do?")

		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		mQuit.SetIcon(icon.Data)

		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
