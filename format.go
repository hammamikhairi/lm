package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

const (
	MEM_SPACE int     = 9
	BLOC_SIZE float64 = 1024.0
)

var SIZES map[int]string = map[int]string{
	0: "",
	1: " K",
	2: " M",
	3: " G",
}

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

func FormatHead(fileName string, isDir bool, executable bool) string {

	head := FileHead(fileName, isDir)
	asterisk := ""

	if executable && !isDir {
		asterisk += "<fg=fb4934>*</>"
	}

	return color.Sprintf("<fg=%s>%s  %s</>%s",
		head.Color,
		head.Icon,
		fileName,
		asterisk,
	)
}

func Format(info os.FileInfo) string {
	modDate := strings.Split(info.ModTime().Format("02/01/2006 03:04 PM"), ".")[0]
	mode := info.Mode().String()
	return fmt.Sprintf("%s%s%s%s%s  %s",
		mode, termGap,
		modDate, termGap,
		MemFmt(info.Size(), info.IsDir()),
		FormatHead(info.Name(), info.IsDir(), mode[3] == 'x'),
	)
}

func FormatSize(size float64, depth int) string {
	if size >= BLOC_SIZE {
		return FormatSize(float64(size)/BLOC_SIZE, depth+1)
	}

	fmtSize := fmt.Sprintf("%.2f%s",
		size,
		SIZES[depth],
	)

	if strings.Contains(fmtSize, ".00") {
		return strings.Replace(fmtSize, ".00", "", 1)
	}
	return fmtSize
}
