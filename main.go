package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/sys/unix"
)

const (
	MEM_SPACE int = 9
)

type (
	ExtMap  map[string]IconCol
	IconCol struct {
		Icon, Color string
	}
)

var (
	dirPath string = ""
	termGap string = ""
	ICONS   ExtMap = ExtMap{
		"go":    IconCol{"", "00CDCD"},
		"jsx":   IconCol{"", "00CDCD"},
		"ts":    IconCol{"", "00CDCD"},
		"tsx":   IconCol{"", "00CDCD"},
		"py":    IconCol{"", "548588"},
		"mod":   IconCol{"", "83a598"},
		"sum":   IconCol{"", "83a598"},
		"sh":    IconCol{"", "B16286"},
		"json":  IconCol{"", "d79921"},
		"sass":  IconCol{"", "b16286"},
		"js":    IconCol{"", "FABD2F"},
		"md":    IconCol{"", "98971A"},
		"css":   IconCol{"", "98971A"},
		"yaml":  IconCol{"", "FE8019"},
		"yml":   IconCol{"", "FE8019"},
		"txt":   IconCol{"", "EBDBB2"},
		"html":  IconCol{"", "D65D0E"},
		"php":   IconCol{"", "B16286"},
		"log":   IconCol{"", "EBDBB2"},
		"jpg":   IconCol{"", "B8BB26"},
		"tmp":   IconCol{"", "CC241D"},
		"ipynb": IconCol{"", "b16286"},
		"pdf":   IconCol{"", "EBDBB2"},
		"csv":   IconCol{"", "928374"},

		"DIR":       IconCol{"", "928374"},
		"gitignore": IconCol{"", "D65D0E"},
		"git":       IconCol{"", "D65D0E"},
		"ignore":    IconCol{" ", "D65D0E"},

		"404":  IconCol{"", "000000"},
		"TEST": IconCol{"", "D79921"},
	}

	SPECIAL_ICONS ExtMap = ExtMap{
		"md":           IconCol{"", "98971A"},
		"LICENSE":      IconCol{"", "CC241D"},
		"package.json": IconCol{"", "CC241D"},
	}

	SIZES map[int]string = map[int]string{
		0: "",
		1: " K",
		2: " M",
		3: " G",
	}
)

const BLOC_SIZE float64 = 1024.0

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
