// +build !windows

package shortcut

import "errors"

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
