// Package shortcut manupulates Windows shortcut file.
package shortcut

import "errors"

// Shortcut contains shortcut properties.
type Shortcut struct {
	// path to be invoked, except arguments.
	TargetPath string

	// optional arguments in a string
	Arguments string

	// optional description of the shortcut
	Description string

	// hotkey to invoke the shortcut.
	// like "Ctrl+Shift+A"
	Hotkey string

	// see https://docs.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/windows-scripting/w88k7fw2%28v%3dvs.84%29
	// 1: normal window(size restored), activated
	// 3: max, activated
	// 7: min, deactivated
	WindowStyle int

	// filepath(fullpath),index
	// like "c:/path/to/file,0"
	IconLocation string

	// working directory
	WorkingDirectory string
}

// Open reads a shortcut file.
func Open(path string) (*Shortcut, error) {
	return nil, errors.New("not supported")
}

// New returns a simle shortcut with default value.
func New(targetPath string) *Shortcut {
	return nil
}

// Save saves a shortcut file.
func (s Shortcut) Save(path string) error {
	return errors.New("not supported")
}
