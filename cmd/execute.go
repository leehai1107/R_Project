package cmd

import "main/windows"

type Execute interface{
  Windows()
}

type execute struct {
	windows windows.Windows
}

func Start() Execute {
	return &execute{
		windows: windows.NewWindows(),
	}
}

func (e *execute) Windows() {
	e.windows.Init()
	e.windows.Process()
	defer e.windows.Close()
}
