package shortcut_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/shu-go/gotwant"
	"github.com/shu-go/shortcut"
)

func TestOpen(t *testing.T) {
	s, err := shortcut.Open("./shortcut")
	gotwant.TestError(t, err, nil)
	gotwant.TestExpr(t, *s, strings.Contains(s.TargetPath, "shortcut.go"))
	gotwant.Test(t, s.Arguments, "a b c")
	gotwant.Test(t, s.Description, "はっはっは！")
	gotwant.Test(t, s.Hotkey, "Ctrl+Shift+A")
	gotwant.Test(t, s.WindowStyle, 3)
	gotwant.Test(t, filepath.Base(s.WorkingDirectory), "shortcut")
}
