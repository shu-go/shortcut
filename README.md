[![](https://godoc.org/github.com/shu-go/shortcut?status.svg)](https://godoc.org/github.com/shu-go/shortcut)
[![Go Report Card](https://goreportcard.com/badge/github.com/shu-go/shortcut)](https://goreportcard.com/report/github.com/shu-go/shortcut)
![MIT License](https://img.shields.io/badge/License-MIT-blue)

# go get

```
go get github.com/shu-go/shortcut
```

It is available only in Windows.

# Usage

```
s, _ := shortcut.Open("./myshortcut")
s.TargetPath = "new path"
s.Save()
```

```
s, _ := shortcut.New("path to target")
s.WindowStyle = 3 // max
s.Arguments = "a b c"
s.IconLocation = "path to exe, 0"
s.Save()
```

# memo (URL Shortcut)

Not supported.

## memo

```
[InternetShortcut]
URL=https://www.google.com/
HotKey=1601
IconIndex=0
IconFile=C:\Program Files (x86)\Google\Chrome\Application\chrome.exe
```

### HotKey

`ACS KEY_ CODE`
`011 0100 0001` Ctrl+Shift+A
`011 0100 0010` Ctrl+Shift+B
`110 0100 0001` Ctrl+Alt+A
