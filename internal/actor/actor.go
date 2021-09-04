package actor

import (
	"github.com/deiwin/interact"
	"os"
)

// Actor interacting with the user over a CLI
var Actor = interact.NewActor(os.Stdin, os.Stdout)
