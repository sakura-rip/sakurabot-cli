package actor

import (
	"github.com/fatih/color"
	"strings"
)

func Input(msg string, opt ...string) string {
	if len(opt) == 0 {
		return color.CyanString("[INPUT]: ") + msg
	}
	return color.CyanString("[INPUT]: ") + msg + " (" + strings.Join(opt, ", ") + ")"
}
