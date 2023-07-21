package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/sys/unix"
)

func MemFmt(mem int64, isDir bool) string {

	if isDir {
		return strings.Repeat(" ", MEM_SPACE)
	}

	var size string = FormatSize(float64(mem), 0)

	return fmt.Sprintf(
		"%s%s",
		strings.Repeat(" ", MEM_SPACE-len(size)),
		size,
	)
}

func FileHead(filename string, isDir bool) IconCol {

	if head, ok := SPECIAL_ICONS[filename]; ok {
		return head
	}

	if strings.Contains(filename, ".test") {
		return ICONS["TEST"]
	}

	if isDir {
		return ICONS["DIR"]
	}

	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex != -1 {
		ext := filename[dotIndex+1:]

		if head, ok := ICONS[ext]; ok {
			return head
		}
	}

	return ICONS["404"]
}

func FormatHead(fileName string, isDir bool) string {

	head := FileHead(fileName, isDir)

	return color.Sprintf("<fg=%s>%s  %s</>",
		head.Color,
		head.Icon,
		fileName,
	)
}

func Format(info os.FileInfo) string {
	modDate := strings.Split(info.ModTime().Format("02/01/2006 03:04 PM"), ".")[0]
	return fmt.Sprintf("%s%s%s%s%s  %s",
		info.Mode().String(), termGap,
		modDate, termGap,
		MemFmt(info.Size(), info.IsDir()),
		FormatHead(info.Name(), info.IsDir()),
	)
}

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
