package main

import "main/windows"

func main() {
	windows.Init()
	windows.Process()
	defer windows.Close()
}
