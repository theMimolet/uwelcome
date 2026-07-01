package state

import (
	"fmt"
	"os"

	"github.com/leonelquinteros/gotext"
)

func getDisabledFile() string {
	return os.ExpandEnv("$HOME/.config/uwelcome/disabled")
}

func IsDisabled() bool {
	_, err := os.Stat(getDisabledFile())
	return err == nil
}

func Enable(l *gotext.Locale) {
	err := os.Remove(getDisabledFile())
	if err != nil {
		fmt.Println(l.Get("Failed to enable uwelcome."))
		println(l.Get("Error ~> %s", err.Error()))
		return
	}
	fmt.Println(l.Get("Uwelcome has been enabled."))
}

func Disable(l *gotext.Locale) {
	_, err := os.Create(getDisabledFile())
	if err != nil {
		fmt.Println(l.Get("Failed to disable uwelcome."))
		println(l.Get("Error ~> %s", err.Error()))
		return
	}
	fmt.Println(l.Get("Uwelcome has been disabled."))
}
