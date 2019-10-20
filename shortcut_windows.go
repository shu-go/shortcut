package shortcut

// https://stackoverflow.com/questions/32438204/create-a-windows-shortcut-lnk-in-go
// https://www.atmarkit.co.jp/ait/articles/0712/27/news083_2.html

import (
	"fmt"
	"strings"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func Open(path string) (*Shortcut, error) {
	if !strings.HasSuffix(strings.ToLower(path), ".lnk") {
		path += ".lnk"
	}

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return nil, fmt.Errorf("CreateObject: %v", err)
	}
	defer oleShellObject.Release()

	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, fmt.Errorf("QueryInterface: %v", err)
	}
	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", path)
	if err != nil {
		return nil, fmt.Errorf("CreateShortcut: %v", err)
	}

	s := &Shortcut{}

	idispatch := cs.ToIDispatch()

	if v, err := oleutil.GetProperty(idispatch, "TargetPath"); err == nil {
		s.TargetPath = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty TargetPath: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "Arguments"); err == nil {
		s.Arguments = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty Arguments: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "Description"); err == nil {
		s.Description = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty Description: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "Hotkey"); err == nil {
		s.Hotkey = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty Hotkey: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "WindowStyle"); err == nil {
		s.WindowStyle = int(v.Val)
	} else {
		return nil, fmt.Errorf("GetProperty WindowStyle: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "IconLocation"); err == nil {
		s.IconLocation = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty IconLocation: %v", err)
	}

	if v, err := oleutil.GetProperty(idispatch, "WorkingDirectory"); err == nil {
		s.WorkingDirectory = v.ToString()
	} else {
		return nil, fmt.Errorf("GetProperty WorkingDirectory: %v", err)
	}

	return s, nil
}

func New(targetPath string) *Shortcut {
	return &Shortcut{
		TargetPath:   targetPath,
		WindowStyle:  1,
		IconLocation: ",0",
	}
}

func (s Shortcut) Save(path string) error {
	if !strings.HasSuffix(strings.ToLower(path), ".lnk") {
		path += ".lnk"
	}

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return fmt.Errorf("CreateObject: %v", err)
	}
	defer oleShellObject.Release()

	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return fmt.Errorf("QueryInterface: %v", err)
	}
	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", path)
	if err != nil {
		return fmt.Errorf("CreateShortcut: %v", err)
	}

	idispatch := cs.ToIDispatch()

	if _, err := oleutil.PutProperty(idispatch, "TargetPath", s.TargetPath); err != nil {
		return fmt.Errorf("PutProperty TargetPath: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "Arguments", s.Arguments); err != nil {
		return fmt.Errorf("PutProperty Arguments: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "Description", s.Description); err != nil {
		return fmt.Errorf("PutProperty Description: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "Hotkey", s.Hotkey); err != nil {
		return fmt.Errorf("PutProperty Hotkey: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "WindowStyle", s.WindowStyle); err != nil {
		return fmt.Errorf("PutProperty WindowStyle: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "IconLocation", s.IconLocation); err != nil {
		return fmt.Errorf("PutProperty IconLocation: %v", err)
	}

	if _, err := oleutil.PutProperty(idispatch, "WorkingDirectory", s.WorkingDirectory); err != nil {
		return fmt.Errorf("PutProperty WorkingDirectory: %v", err)
	}

	if _, err := oleutil.CallMethod(idispatch, "Save"); err != nil {
		return fmt.Errorf("Save: %v", err)
	}

	return nil
}
