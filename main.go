package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

var (
	dirPath string = ""
	termGap string = ""
)

func init() {

	if len(os.Args) > 1 {
		dirPath = os.Args[1]
	}

	ws, _ := unix.IoctlGetWinsize(unix.Stdout, unix.TIOCGWINSZ)
	termGap = strings.Repeat(" ", int(ws.Col)%10)
}

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	files, err := os.ReadDir(dir + "/" + dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		info, _ := file.Info()
		if file.IsDir() {
			fmt.Println(Format(info))
		} else {
			defer func() {
				fmt.Println(Format(info))
			}()
		}
	}
}
