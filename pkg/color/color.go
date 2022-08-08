package color

import (
	// import third-party packages
	"github.com/fatih/color"
)

func Blue(text string) string {
	blue := color.New(color.FgBlue).SprintFunc()
	// return
	return blue(text)
}

func Green(text string) string {
	green := color.New(color.FgGreen).SprintFunc()
	// return
	return green(text)
}
