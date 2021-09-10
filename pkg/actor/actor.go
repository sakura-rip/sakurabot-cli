package actor

import (
	"github.com/deiwin/interact"
	"os"
)

// DefaultActor interacting with the user over a CLI
var DefaultActor = interact.NewActor(os.Stdin, os.Stdout)

// PromptAndRetry asks the user for input and performs the list of added checks
func PromptAndRetry(message string, checks ...interact.InputCheck) (string, error) {
	return DefaultActor.PromptAndRetry(Input(message), checks...)
}

// PromptOptionalAndRetry works exactly like GetInputAndRetry, but also has
// a default option which will be used instead if the user simply presses enter.
func PromptOptionalAndRetry(message, defaultOption string, checks ...interact.InputCheck) (string, error) {
	return DefaultActor.PromptOptionalAndRetry(Input(message), defaultOption, checks...)
}

// Prompt asks the user for input and performs the list of added checks on the
// provided input. If any of the checks fail, the error will be returned.
func Prompt(message string, checks ...interact.InputCheck) (string, error) {
	return DefaultActor.Prompt(Input(message), checks...)
}

// PromptOptional works exactly like Prompt, but also has a default option
// which will be used instead if the user simply presses enter.
func PromptOptional(message, defaultOption string, checks ...interact.InputCheck) (string, error) {
	return DefaultActor.PromptOptional(Input(message), defaultOption, checks...)
}
